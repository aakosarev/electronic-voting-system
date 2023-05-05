package handler

import (
	"context"
	pb "github.com/aakosarev/electronic-voting-system/contracts/gen/go/electronic-voting-manager/v1"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-manager/internal/service"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Handler struct {
	service *service.Service
	pb.UnimplementedVotingManagerServer
}

func NewHandler(service *service.Service, srv pb.UnimplementedVotingManagerServer) *Handler {
	return &Handler{
		service:                          service,
		UnimplementedVotingManagerServer: srv,
	}
}

func (h *Handler) CreateVoting(ctx context.Context, req *pb.CreateVotingRequest) (*emptypb.Empty, error) {
	err := h.service.CreateVoting(ctx, req.GetVotingTitle(), req.GetVotingOptions(), req.GetEndTime().AsTime())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (h *Handler) GetAllVotings(ctx context.Context, _ *emptypb.Empty) (*pb.GetAllVotingsResponse, error) {
	votings, err := h.service.GetAllVotings(ctx)
	if err != nil {
		return nil, err
	}

	pbVotings := make([]*pb.Voting, len(votings))
	for i, v := range votings {
		pbVotings[i] = &pb.Voting{
			VotingID:    v.ID,
			VotingTitle: v.Title,
			EndTime:     v.EndTime,
			Address:     v.Address,
			CreatedOn:   timestamppb.New(v.CreatedOn),
		}
	}

	return &pb.GetAllVotingsResponse{Votings: pbVotings}, nil
}

func (h *Handler) AddRightToVote(ctx context.Context, req *pb.AddRightToVoteRequest) (*emptypb.Empty, error) {
	err := h.service.AddRightToVote(ctx, req.GetUserID(), req.GetVotingID())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (h *Handler) GetVotingsAvailableToUser(ctx context.Context, req *pb.GetVotingsAvailableToUserRequest) (*pb.GetVotingsAvailableToUserResponse, error) {
	votingsAvailableToUser, err := h.service.GetVotingsAvailableToUser(ctx, req.GetUserID())
	if err != nil {
		return nil, err
	}

	pbVotingsAvailableToUser := make([]*pb.VotingAvailableToUser, len(votingsAvailableToUser))
	for i, v := range votingsAvailableToUser {
		pbVotingsAvailableToUser[i] = &pb.VotingAvailableToUser{
			UserID:        v.UserID,
			VotingID:      v.VotingID,
			CreatedOn:     timestamppb.New(v.CreatedOn),
			VotingTitle:   v.Title,
			VotingAddress: v.Address,
		}
	}

	return &pb.GetVotingsAvailableToUserResponse{VotingsAvailableToUser: pbVotingsAvailableToUser}, nil
}

func (h *Handler) RegisterAddressToVoting(ctx context.Context, req *pb.RegisterAddressToVotingRequest) (*emptypb.Empty, error) {

	err := h.service.RegisterAddressToVoting(ctx, req.GetVotingID(), req.GetAddress())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
