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
	"net/http"
	"os"
	"strconv"
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
	userID    int32
	expiresAt time.Time
}

func (s session) isExpired() bool {
	return s.expiresAt.Before(time.Now())
}

var sessions = map[string]session{}

func (h *Handler) Register(router *httprouter.Router) {
	router.POST("/api/login", h.Login)
	router.PATCH("/api/initial_login", AuthMiddleware(h.InitialLogin))
	router.GET("/api/logout", h.Logout)
	router.GET("/api/force_enter_details", AuthMiddleware(h.ForceEnterDetails))
	router.GET("/api/available_votings", AuthMiddleware(h.AvailableVotings))
	router.GET("/api/public_key_voting/:voting_id", AuthMiddleware(h.PublicKeyVoting))
	router.GET("/api/sign_blinded_public_key", AuthMiddleware(h.SignBlindedPublicKey))

	//router.HandlerFunc(http.MethodPost, "/api/register_to_voting", AuthMiddleware(h.RegisterToVoting))
}

func (h *Handler) ForceEnterDetails(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	c, _ := r.Context().Value("cookie_session_token").(*http.Cookie)
	sessionToken := c.Value
	userSession := sessions[sessionToken]

	forceEnterDetails, err := h.userStorage.GetForceEnterDetailsByUserID(r.Context(), userSession.userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"force_enter_details": %t}`, forceEnterDetails))) //TODO remake to response-struct
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	loginReq := model.LoginReq{}
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	passwordHash, err := h.userStorage.GetPasswordHashByUserID(r.Context(), loginReq.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(loginReq.Password)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(1 * time.Hour)

	sessions[sessionToken] = session{
		userID:    loginReq.UserID,
		expiresAt: expiresAt,
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
	})
}

func (h *Handler) InitialLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c, _ := r.Context().Value("cookie_session_token").(*http.Cookie)
	sessionToken := c.Value
	userSession := sessions[sessionToken]

	var initialLoginReq model.InitialLoginReq
	err := json.NewDecoder(r.Body).Decode(&initialLoginReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if initialLoginReq.NewPassword != initialLoginReq.ReenteredPassword {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	passwordHash, err := h.userStorage.GetPasswordHashByUserID(r.Context(), userSession.userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(initialLoginReq.OldPassword)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(initialLoginReq.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.userStorage.AddDetails(r.Context(), userSession.userID, string(newPasswordHash),
		initialLoginReq.Email, initialLoginReq.Name, initialLoginReq.Surname)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

func (h *Handler) AvailableVotings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	c, _ := r.Context().Value("cookie_session_token").(*http.Cookie)
	sessionToken := c.Value
	userSession := sessions[sessionToken]

	getVotingsAvailableToUserIDReqToVM := &pbvm.GetVotingsAvailableToUserIDRequest{
		UserID: userSession.userID,
	}

	getVotingsAvailableToUserIDRespFromVM, err := h.votingManagerClient.GetVotingsAvailableToUser(r.Context(),
		getVotingsAvailableToUserIDReqToVM)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	pbVotingsAvailableToUserID := getVotingsAvailableToUserIDRespFromVM.GetVotingsAvailableToUserID()

	var availableVotings []*model.AvailableVoting

	for _, av := range pbVotingsAvailableToUserID {
		availableVotings = append(availableVotings, &model.AvailableVoting{
			UserID:     av.GetUserID(),
			VotingID:   av.GetVotingID(),
			CreatedOn:  av.GetCreatedOn().AsTime(),
			Title:      av.GetVotingTitle(),
			Address:    av.GetVotingAddress(),
			Registered: false, //TODO to change!
		})
	}

	votingsAvailableToUserIDResp := model.VotingsAvailableToUserIDResp{
		AvailableVotings: availableVotings,
	}

	votingsAvailableToUserIDRespJson, err := json.Marshal(votingsAvailableToUserIDResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(votingsAvailableToUserIDRespJson)
}

func (h *Handler) PublicKeyVoting(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	votingIDStr := ps.ByName("voting_id")
	votingID, err := strconv.ParseInt(votingIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	getPublicKeyForVotingIDReqToVV := &pbvv.GetPublicKeyForVotingIDRequest{
		VotingID: int32(votingID),
	}

	getPublicKeyForVotingIDRespFromVV, err := h.votingVerifierClient.GetPublicKeyForVotingID(r.Context(), getPublicKeyForVotingIDReqToVV)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	publicKeyBytes := getPublicKeyForVotingIDRespFromVV.GetPublicKeyBytes()

	publicKeyResp := model.PublicKeyResp{
		PublicKeyBytes: publicKeyBytes,
	}

	publicKeyRespJson, err := json.Marshal(publicKeyResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(publicKeyRespJson)
}

func (h *Handler) SignBlindedPublicKey(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

}

func (h *Handler) RegisterToVoting(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c, _ := r.Context().Value("cookie_session_token").(*http.Cookie)
	sessionToken := c.Value
	userSession := sessions[sessionToken]

	var registerToVotingReq model.RegisterToVotingReq
	err := json.NewDecoder(r.Body).Decode(&registerToVotingReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqGetPublicKeyBytes := &pbvv.GetPublicKeyForVotingIDRequest{VotingID: registerToVotingReq.VotingID}

	respGetPublicKeyBytes, err := h.votingVerifierClient.GetPublicKeyForVotingID(r.Context(), reqGetPublicKeyBytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	publicKeyBytes := respGetPublicKeyBytes.GetPublicKeyBytes()

	publicKeyPEM, _ := pem.Decode(publicKeyBytes)
	publicKey, err := x509.ParsePKCS1PublicKey(publicKeyPEM.Bytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token := make([]byte, 16)
	_, err = rand.Read(token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	blindedToken, unblinder, err := rsablind.Blind(publicKey, token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	reqSignBlindedToken := &pbvv.SignBlindedTokenRequest{
		UserID:       userSession.userID,
		VotingID:     registerToVotingReq.VotingID,
		BlindedToken: blindedToken,
	}

	respSignBlindedToken, err := h.votingVerifierClient.SignBlindedToken(r.Context(), reqSignBlindedToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	signedBlindedToken := respSignBlindedToken.GetSignedBlindedToken()

	signedToken := rsablind.Unblind(publicKey, signedBlindedToken, unblinder)

	passwordHash, err := h.userStorage.GetPasswordHashByUserID(r.Context(), userSession.userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	wd, _ := os.Getwd()
	keyStorePath := fmt.Sprintf("%s/internal/keystorage", wd)
	address, err := eth.GenerateNewAccount(keyStorePath, passwordHash)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	reqToVotingVerifier := &pbvv.RegisterAddressToVotingBySignedTokenRequest{
		VotingID:    registerToVotingReq.VotingID,
		Token:       token,
		SignedToken: signedToken,
		Address:     address,
	}

	_, err = h.votingVerifierClient.RegisterAddressToVotingBySignedToken(r.Context(), reqToVotingVerifier)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
