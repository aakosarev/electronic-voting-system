package handler

import (
	"encoding/json"
	pbva "github.com/aakosarev/electronic-voting-system/contracts/gen/go/electronic-voting-app/v1"
	pbvm "github.com/aakosarev/electronic-voting-system/contracts/gen/go/electronic-voting-manager/v1"
	pbvv "github.com/aakosarev/electronic-voting-system/contracts/gen/go/electronic-voting-verifier/v1"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-registrator/internal/model"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Handler struct {
	votingManagerClient  pbvm.VotingManagerClient
	votingAppClient      pbva.VotingAppClient
	votingVerifierClient pbvv.VotingVerifierClient
}

func NewHandler(connVotingManager, connVotingApp, connVotingVerifier *grpc.ClientConn) *Handler {
	return &Handler{
		votingManagerClient:  pbvm.NewVotingManagerClient(connVotingManager),
		votingAppClient:      pbva.NewVotingAppClient(connVotingApp),
		votingVerifierClient: pbvv.NewVotingVerifierClient(connVotingVerifier),
	}
}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, "/api/create_voting", h.CreateVoting)
	router.HandlerFunc(http.MethodPost, "/api/add_right_to_vote", h.AddRightToVote)
	router.HandlerFunc(http.MethodGet, "/api/votings", h.GetAllVotings)
}

func (h *Handler) CreateVoting(w http.ResponseWriter, r *http.Request) {
	createVotingReq := model.CreateVotingReq{}
	err := json.NewDecoder(r.Body).Decode(&createVotingReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	options := strings.Split(createVotingReq.Options, ";")

	createVotingReqToVM := &pbvm.CreateVotingRequest{
		VotingTitle:   createVotingReq.Title,
		VotingOptions: options,
		EndTime:       timestamppb.New(createVotingReq.EndTime),
	}

	createVotingRespFromVM, err := h.votingManagerClient.CreateVoting(r.Context(), createVotingReqToVM)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	generateRSAKeyPairForVotingIDReqToVV := &pbvv.GenerateRSAKeyPairForVotingIDRequest{
		VotingID: createVotingRespFromVM.GetVotingID(),
	}

	_, err = h.votingVerifierClient.GenerateRSAKeyPairForVotingID(r.Context(), generateRSAKeyPairForVotingIDReqToVV)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) AddRightToVote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	addRightToVoteReq := model.AddRightToVoteReq{}
	err := json.NewDecoder(r.Body).Decode(&addRightToVoteReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rand.Seed(time.Now().UnixNano())
	userID := rand.Int31n(99001) + 1000

	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+"
	b := make([]byte, 20)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}

	passwordHash, err := bcrypt.GenerateFromPassword(b, bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	registerUserReqToVA := &pbva.RegisterUserRequest{
		Username:     userID,
		PasswordHash: string(passwordHash),
	}

	_, err = h.votingAppClient.RegisterUser(r.Context(), registerUserReqToVA)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, votingID := range addRightToVoteReq.VotingIDs {
		addRightToVoteReqToVM := &pbvm.AddRightToVoteRequest{
			UserID:   userID,
			VotingID: votingID,
		}
		_, err = h.votingManagerClient.AddRightToVote(r.Context(), addRightToVoteReqToVM)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	addRightToVoteResp := model.AddRightToVoteResp{
		UserID:   userID,
		Password: string(b),
	}

	addRightToVoteRespJson, err := json.Marshal(addRightToVoteResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(addRightToVoteRespJson)
}

func (h *Handler) GetAllVotings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	getAllVotingsRespFromVM, err := h.votingManagerClient.GetAllVotings(r.Context(), &emptypb.Empty{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	pbVotings := getAllVotingsRespFromVM.GetVotings()
	var votings []*model.Voting

	for _, v := range pbVotings {
		votings = append(votings, &model.Voting{
			ID:        v.GetVotingID(),
			Title:     v.GetVotingTitle(),
			EndTime:   v.GetEndTime(),
			Address:   v.GetAddress(),
			CreatedOn: v.GetCreatedOn().AsTime(),
		})
	}

	getAllVotingsResp := model.GetAllVotingsResp{
		Votings: votings,
	}

	getAllVotingsRespJson, err := json.Marshal(getAllVotingsResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(getAllVotingsRespJson)
}
