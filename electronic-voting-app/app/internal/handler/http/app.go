package http

import (
	"crypto/rand"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	pbvm "github.com/aakosarev/electronic-voting-system/contracts/gen/go/electronic-voting-manager/v1"
	pbvv "github.com/aakosarev/electronic-voting-system/contracts/gen/go/electronic-voting-verifier/v1"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-app/eth"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-app/internal/model"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-app/internal/storage"
	"github.com/cryptoballot/rsablind"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"time"
)

type Handler struct {
	votingManagerClient  pbvm.VotingManagerClient
	votingVerifierClient pbvv.VotingVerifierClient
	userStorage          *storage.UserStorage
}

func NewHandler(userStorage *storage.UserStorage, connVotingManager, connVotingVerifier *grpc.ClientConn) *Handler {
	return &Handler{
		userStorage:          userStorage,
		votingManagerClient:  pbvm.NewVotingManagerClient(connVotingManager),
		votingVerifierClient: pbvv.NewVotingVerifierClient(connVotingVerifier),
	}
}

type session struct {
	username  int32
	expiresAt time.Time
}

func (s session) isExpired() bool {
	return s.expiresAt.Before(time.Now())
}

var sessions = map[string]session{}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, "/login", h.Login)
	router.HandlerFunc(http.MethodPost, "/initial_login", AuthMiddleware(h.InitialLogin))
	router.HandlerFunc(http.MethodGet, "/logout", h.Logout)
	router.HandlerFunc(http.MethodPost, "/refresh", AuthMiddleware(h.Refresh))
	router.HandlerFunc(http.MethodGet, "/force_enter_details", AuthMiddleware(h.ForceEnterDetails))
	router.HandlerFunc(http.MethodGet, "/available_votings", AuthMiddleware(h.AvailableVotings))
	router.HandlerFunc(http.MethodPost, "/register_to_voting", AuthMiddleware(h.RegisterToVoting))
}

func (h *Handler) ForceEnterDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	c, _ := r.Context().Value("cookie_session_token").(*http.Cookie)
	sessionToken := c.Value
	userSession := sessions[sessionToken]

	forceEnterDetails, err := h.userStorage.GetForceEnterDetailsByUsername(r.Context(), userSession.username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"force_enter_details": %t}`, forceEnterDetails)))
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var loginUser model.LoginUser
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	passwordHash, err := h.userStorage.GetPasswordHashByUsername(r.Context(), loginUser.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(loginUser.Password)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(1 * time.Hour)

	sessions[sessionToken] = session{
		username:  loginUser.Username,
		expiresAt: expiresAt,
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
	})
}

func (h *Handler) InitialLogin(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Context().Value("cookie_session_token").(*http.Cookie)
	sessionToken := c.Value
	userSession := sessions[sessionToken]

	var ilu model.InitialLoginUser
	err := json.NewDecoder(r.Body).Decode(&ilu)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if ilu.NewPassword != ilu.ReenteredPassword {
		w.WriteHeader(http.StatusBadRequest)
		// TODO w.Write error
		return
	}

	passwordHash, err := h.userStorage.GetPasswordHashByUsername(r.Context(), userSession.username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(ilu.OldPassword)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		// TODO w.Write error
		return
	}

	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(ilu.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.userStorage.AddDetails(r.Context(), userSession.username, string(newPasswordHash), ilu.Email, ilu.FirstName, ilu.SecondName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sessionToken := c.Value
	delete(sessions, sessionToken)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now(),
	})
}

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Context().Value("cookie_session_token").(*http.Cookie)
	sessionToken := c.Value
	userSession := sessions[sessionToken]

	newSessionToken := uuid.NewString()
	expiresAt := time.Now().Add(1 * time.Hour)

	sessions[newSessionToken] = session{
		username:  userSession.username,
		expiresAt: expiresAt,
	}

	delete(sessions, sessionToken)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   newSessionToken,
		Expires: expiresAt,
	})
}

func (h *Handler) AvailableVotings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	c, _ := r.Context().Value("cookie_session_token").(*http.Cookie)
	sessionToken := c.Value
	userSession := sessions[sessionToken]

	req := &pbvm.GetVotingsAvailableToUserRequest{
		UserID: userSession.username,
	}

	resp, err := h.votingManagerClient.GetVotingsAvailableToUser(r.Context(), req)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	pbAvailableVotings := resp.GetVotingsAvailableToUser()

	var availableVotings []*model.VotingAvailableToUser

	for _, av := range pbAvailableVotings {
		availableVotings = append(availableVotings, &model.VotingAvailableToUser{
			UserID:     av.GetUserID(),
			VotingID:   av.GetVotingID(),
			CreatedOn:  av.GetCreatedOn().AsTime(),
			Title:      av.GetVotingTitle(),
			Address:    av.GetVotingAddress(),
			Registered: false, //TODO to change!
			Requested:  false, //TODO to change!
		})
	}

	availableVotingsJson, err := json.Marshal(availableVotings)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(availableVotingsJson)
}

func (h *Handler) RegisterToVoting(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Context().Value("cookie_session_token").(*http.Cookie)
	sessionToken := c.Value
	userSession := sessions[sessionToken]

	var votingID int32
	err := json.NewDecoder(r.Body).Decode(&votingID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqGetPublicKeyBytes := &pbvv.GetPublicKeyForVotingIDRequest{VotingID: votingID}

	respGetPublicKeyBytes, err := h.votingVerifierClient.GetPublicKeyForVotingID(r.Context(), reqGetPublicKeyBytes)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	publicKeyBytes := respGetPublicKeyBytes.GetPublicKeyBytes()

	publicKeyPEM, _ := pem.Decode(publicKeyBytes)
	publicKey, err := x509.ParsePKCS1PublicKey(publicKeyPEM.Bytes)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token := make([]byte, 16)
	_, err = rand.Read(token)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	blindedToken, unblinder, err := rsablind.Blind(publicKey, token)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	reqSignBlindedToken := &pbvv.SignBlindedTokenRequest{
		UserID:       userSession.username,
		VotingID:     votingID,
		BlindedToken: string(blindedToken),
	}

	respSignBlindedToken, err := h.votingVerifierClient.SignBlindedToken(r.Context(), reqSignBlindedToken)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	signedBlindedToken := respSignBlindedToken.GetSignedBlindedToken()

	signedToken := rsablind.Unblind(publicKey, []byte(signedBlindedToken), unblinder)

	passwordHash, err := h.userStorage.GetPasswordHashByUsername(r.Context(), userSession.username)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	wd, _ := os.Getwd()
	keyStorePath := fmt.Sprintf("%s/internal/keystorage", wd)
	address, err := eth.GenerateNewAccount(keyStorePath, passwordHash) //TODO md use pass (not passHash)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_ = signedToken // TODO delete
	_ = address     // TODO delete

}
