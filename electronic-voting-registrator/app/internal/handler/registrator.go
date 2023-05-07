package handler

import (
	"encoding/json"
	"fmt"
	pbva "github.com/aakosarev/electronic-voting-system/contracts/gen/go/electronic-voting-app/v1"
	pbvm "github.com/aakosarev/electronic-voting-system/contracts/gen/go/electronic-voting-manager/v1"
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
	votingManagerClient pbvm.VotingManagerClient
	votingAppClient     pbva.VotingAppClient
}

func NewHandler(connVotingManager, connVotingApp *grpc.ClientConn) *Handler {
	return &Handler{
		votingManagerClient: pbvm.NewVotingManagerClient(connVotingManager),
		votingAppClient:     pbva.NewVotingAppClient(connVotingApp),
	}
}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, "/create_voting", h.CreateVoting)
	router.HandlerFunc(http.MethodPost, "/add_right_to_vote", h.AddRightToVote)
	router.HandlerFunc(http.MethodGet, "/votings", h.GetAllVotings)
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

	_, err = h.votingManagerClient.CreateVoting(r.Context(), createVotingReqToVM)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//request to verifier for creating RSA key pair

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
	password := string(b)

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	reqToVotingApp := &pbva.RegisterUserRequest{
		Username:     userID,
		PasswordHash: string(passwordHash),
	}

	_, err = h.votingAppClient.RegisterUser(r.Context(), reqToVotingApp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, votingID := range addRightToVoteReq.VotingIDs {
		reqToVotingManager := &pbvm.AddRightToVoteRequest{
			UserID:   userID,
			VotingID: votingID,
		}
		_, err = h.votingManagerClient.AddRightToVote(r.Context(), reqToVotingManager)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	credentials := model.AddRightToVoteResp{
		UserID:   userID,
		Password: password,
	}

	credentialsJson, err := json.Marshal(credentials)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(credentialsJson)
}

func (h *Handler) GetAllVotings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp, err := h.votingManagerClient.GetAllVotings(r.Context(), &emptypb.Empty{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, err.Error())))
		return
	}

	pbVotings := resp.GetVotings()
	var votings []*model.GetAllVotingsResp

	for _, v := range pbVotings {
		votings = append(votings, &model.GetAllVotingsResp{
			VotingID:    v.GetVotingID(),
			VotingTitle: v.GetVotingTitle(),
			EndTime:     v.GetEndTime(),
			Address:     v.GetAddress(),
			CreatedOn:   v.GetCreatedOn().AsTime(),
		})
	}

	votingsJson, err := json.Marshal(votings)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, err.Error())))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(votingsJson)
}
