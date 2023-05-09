package http

import (
	"context"
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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"net/http"
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
	// basic endpoints
	router.POST("/api/login", h.Login)
	router.PATCH("/api/initial_login", AuthMiddleware(h.InitialLogin))
	router.GET("/api/logout", h.Logout)
	router.GET("/api/force_enter_details", AuthMiddleware(h.ForceEnterDetails))
	router.GET("/api/available_votings", AuthMiddleware(h.AvailableVotings))
	router.GET("/api/voting_information", AuthMiddleware(h.VotingInformation))

	//to interact with the client
	router.GET("/api/public_key_voting", AuthMiddleware(h.PublicKeyVoting))
	router.GET("/api/sign_blinded_address", AuthMiddleware(h.SignBlindedAddress))
	router.POST("/api/register_address", h.RegisterAddress)

	router.GET("/api/registration_statuses", h.RegistrationStatuses)

	//since it was not planned to write the client part, a demo part was
	//developed where the client logic occurs on the server
	router.POST("/demo/api/register_to_voting", AuthMiddleware(h.DemoRegisterToVoting))

	router.GET("/demo/api/registration_statuses", AuthMiddleware(h.DemoRegistrationStatuses))
	router.POST("/demo/api/cast_a_vote", AuthMiddleware(h.CastAVote))

}

func (h *Handler) ForceEnterDetails(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

func (h *Handler) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
		Path:    "/",
	})
}

func (h *Handler) InitialLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

func (h *Handler) AvailableVotings(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	c, _ := r.Context().Value("cookie_session_token").(*http.Cookie)
	sessionToken := c.Value
	userSession := sessions[sessionToken]

	getVotingsAvailableToUserIDReqToVM := &pbvm.GetVotingsAvailableToUserIDRequest{
		UserID: userSession.userID,
	}

	getVotingsAvailableToUserIDRespFromVM, err := h.votingManagerClient.GetVotingsAvailableToUserID(r.Context(),
		getVotingsAvailableToUserIDReqToVM)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	pbVotingsAvailableToUserID := getVotingsAvailableToUserIDRespFromVM.GetVotingsAvailableToUserID()

	var availableVotings []*model.AvailableVoting

	for _, av := range pbVotingsAvailableToUserID {
		availableVotings = append(availableVotings, &model.AvailableVoting{
			UserID:    av.GetUserID(),
			VotingID:  av.GetVotingID(),
			CreatedOn: av.GetCreatedOn().AsTime(),
			Title:     av.GetVotingTitle(),
			Address:   av.GetVotingAddress(),
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

	queryValues := r.URL.Query()
	votingIDStr := queryValues.Get("voting_id")

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

func (h *Handler) SignBlindedAddress(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	c, _ := r.Context().Value("cookie_session_token").(*http.Cookie)
	sessionToken := c.Value
	userSession := sessions[sessionToken]

	var signBlindedAddressReq model.SignBlindedAddressReq

	err := json.NewDecoder(r.Body).Decode(&signBlindedAddressReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	signBlindedAddressReqToVV := &pbvv.SignBlindedAddressRequest{
		UserID:         userSession.userID,
		VotingID:       signBlindedAddressReq.VotingID,
		BlindedAddress: signBlindedAddressReq.BlindedAddress,
	}

	signBlindedAddressRespFromVV, err := h.votingVerifierClient.SignBlindedAddress(r.Context(), signBlindedAddressReqToVV)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	signedBlindedAddress := signBlindedAddressRespFromVV.GetSignedBlindedAddress()

	signedBlindedAddressResp := model.SignBlindedAddressResp{
		SignedBlindedAddress: signedBlindedAddress,
	}

	signedBlindedAddressRespJson, err := json.Marshal(signedBlindedAddressResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(signedBlindedAddressRespJson)
}

func (h *Handler) RegisterAddress(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var registerAddressReq model.RegisterAddressReq

	err := json.NewDecoder(r.Body).Decode(&registerAddressReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	registerAddressReqToVV := &pbvv.RegisterAddressRequest{
		Address:       registerAddressReq.Address,
		SignedAddress: registerAddressReq.SignedAddress,
		VotingID:      registerAddressReq.VotingID,
	}

	_, err = h.votingVerifierClient.RegisterAddress(r.Context(), registerAddressReqToVV)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) RegistrationStatuses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	var registrationStatusesReq model.RegistrationStatusesReq

	err := json.NewDecoder(r.Body).Decode(&registrationStatusesReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	getRegistrationStatusesReqToVV := &pbvv.GetRegistrationStatusesRequest{
		Addresses: registrationStatusesReq.Addresses,
	}

	getRegistrationStatusesRespFromVV, err := h.votingVerifierClient.GetRegistrationStatuses(r.Context(), getRegistrationStatusesReqToVV)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	statuses := getRegistrationStatusesRespFromVV.GetStatuses()

	registrationStatusesResp := model.RegistrationStatusesResp{
		Statuses: statuses,
	}

	registrationStatusesRespJson, err := json.Marshal(registrationStatusesResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(registrationStatusesRespJson)
}

func (h *Handler) VotingInformation(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	queryValues := r.URL.Query()
	votingIDStr := queryValues.Get("voting_id")

	votingID, err := strconv.ParseInt(votingIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	getVotingInformationReqToVM := &pbvm.GetVotingInformationRequest{
		VotingID: int32(votingID),
	}

	getVotingInformationRespFromVM, err := h.votingManagerClient.GetVotingInformation(r.Context(), getVotingInformationReqToVM)

	options := make(map[int64]*model.Option, len(getVotingInformationRespFromVM.GetOptions()))

	for i, option := range getVotingInformationRespFromVM.GetOptions() {
		options[i] = &model.Option{
			Name:        option.GetName(),
			NumberVotes: option.GetNumberVotes(),
		}
	}

	votingInformationResp := model.VotingInformationResp{
		Title:                  getVotingInformationRespFromVM.GetTitle(),
		NumberRegisteredVoters: getVotingInformationRespFromVM.GetNumberRegisteredVoters(),
		EndTime:                getVotingInformationRespFromVM.GetEndTime().AsTime(),
		Options:                options,
	}

	votingInformationRespJson, err := json.Marshal(votingInformationResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(votingInformationRespJson)
}

// ------------- demo (the logic that should be on the client is implemented on the server)-------------

func (h *Handler) DemoRegisterToVoting(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c, _ := r.Context().Value("cookie_session_token").(*http.Cookie)
	sessionToken := c.Value
	userSession := sessions[sessionToken]

	queryValues := r.URL.Query()

	votingIDStr := queryValues.Get("voting_id")

	votingID, err := strconv.ParseInt(votingIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	address, privateKey, err := eth.GenerateNewKeyPair()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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

	publicKeyPEM, _ := pem.Decode(publicKeyBytes)
	publicKey, err := x509.ParsePKCS1PublicKey(publicKeyPEM.Bytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	blindedAddress, unblinder, err := rsablind.Blind(publicKey, []byte(address))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	signBlindedAddressReqToVV := &pbvv.SignBlindedAddressRequest{
		UserID:         userSession.userID,
		VotingID:       int32(votingID),
		BlindedAddress: blindedAddress,
	}

	signBlindedAddressRespFromVV, err := h.votingVerifierClient.SignBlindedAddress(r.Context(), signBlindedAddressReqToVV)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	signedBlindedAddress := signBlindedAddressRespFromVV.GetSignedBlindedAddress()

	signedAddress := rsablind.Unblind(publicKey, signedBlindedAddress, unblinder)

	registerAddressReqToVV := &pbvv.RegisterAddressRequest{
		Address:       address,
		SignedAddress: signedAddress,
		VotingID:      int32(votingID),
	}

	_, err = h.votingVerifierClient.RegisterAddress(r.Context(), registerAddressReqToVV)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.userStorage.DemoAddRegistration(r.Context(), privateKey, userSession.userID, int32(votingID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DemoRegistrationStatuses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	c, _ := r.Context().Value("cookie_session_token").(*http.Cookie)
	sessionToken := c.Value
	userSession := sessions[sessionToken]

	var demoRegistrationStatusesReq model.DemoRegistrationStatusesReq

	err := json.NewDecoder(r.Body).Decode(&demoRegistrationStatusesReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	existsPrivateKeys, err := h.userStorage.DemoGetPrivateKeys(r.Context(), userSession.userID, demoRegistrationStatusesReq.VotingIDs)

	existsAddresses, err := eth.GetAddressesByPrivateKeys(existsPrivateKeys)

	addresses := make(map[int32]string)

	for i := 0; i < len(demoRegistrationStatusesReq.VotingIDs); i++ {
		if address, ok := existsAddresses[demoRegistrationStatusesReq.VotingIDs[i]]; ok {
			addresses[demoRegistrationStatusesReq.VotingIDs[i]] = address
		} else {
			addresses[demoRegistrationStatusesReq.VotingIDs[i]] = ""
		}
	}

	getRegistrationStatusesReqToVV := &pbvv.GetRegistrationStatusesRequest{
		Addresses: addresses,
	}

	getRegistrationStatusesRespFromVV, err := h.votingVerifierClient.GetRegistrationStatuses(r.Context(), getRegistrationStatusesReqToVV)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	statuses := getRegistrationStatusesRespFromVV.GetStatuses()

	registrationStatusesResp := model.RegistrationStatusesResp{
		Statuses: statuses,
	}

	registrationStatusesRespJson, err := json.Marshal(registrationStatusesResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(registrationStatusesRespJson)
}

func (h *Handler) CastAVote(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	c, _ := r.Context().Value("cookie_session_token").(*http.Cookie)
	sessionToken := c.Value
	userSession := sessions[sessionToken]

	var demoCastAVoteReq model.DemoCastAVoteReq

	err := json.NewDecoder(r.Body).Decode(&demoCastAVoteReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	privateKey, err := h.userStorage.DemoGetPrivateKey(r.Context(), userSession.userID, demoCastAVoteReq.VotingID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	eclient, err := ethclient.Dial("https://sepolia.infura.io/v3/4acf1231e76946a59e2f3c2cfc8ce3db")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ethSession, err := eth.NewSession(context.Background(), eclient, privateKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	contractAddress := common.HexToAddress(demoCastAVoteReq.Address)

	transactionHash, err := eth.CastAVote(ethSession, eclient, contractAddress, demoCastAVoteReq.Index)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	demoCastAVoteResp := model.DemoCastAVoteResp{
		LinkToTransaction: fmt.Sprintf("https://sepolia.etherscan.io/tx/%s", transactionHash),
	}

	demoCastAVoteRespJson, err := json.Marshal(demoCastAVoteResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(demoCastAVoteRespJson)
}
