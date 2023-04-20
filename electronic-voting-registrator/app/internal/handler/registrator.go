package handler

import (
	"encoding/json"
	"fmt"
	pbvm "github.com/aakosarev/electronic-voting-system/contracts/gen/go/electronic-voting-manager/v1"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-registrator/internal/model"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"strings"
)

type Handler struct {
	votingManagerClient pbvm.VotingManagerClient
}

func NewHandler(connVotingManager *grpc.ClientConn) *Handler {
	return &Handler{votingManagerClient: pbvm.NewVotingManagerClient(connVotingManager)}
}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, "/create_voting", h.CreateVoting)
	router.HandlerFunc(http.MethodPost, "/add_right_to_vote", h.AddRightToVote)
	router.HandlerFunc(http.MethodGet, "/votings", h.GetAllVotings)
}

func (h *Handler) CreateVoting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	createVotingReq := model.CreateVotingReq{}
	err := json.NewDecoder(r.Body).Decode(&createVotingReq)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	options := strings.Split(createVotingReq.Options, ";")

	req := &pbvm.CreateVotingRequest{
		VotingTitle:   createVotingReq.Title,
		VotingOptions: options,
		EndTime:       timestamppb.New(createVotingReq.EndTime),
	}

	_, err = h.votingManagerClient.CreateVoting(r.Context(), req)
	if err != nil {
		fmt.Println(err)
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

	req := &pbvm.AddRightToVoteRequest{
		UserID:   addRightToVoteReq.UserID,
		VotingID: addRightToVoteReq.VotingID,
	}

	_, err = h.votingManagerClient.AddRightToVote(r.Context(), req)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
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
		w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, err.Error())))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(votingsJson)
}
