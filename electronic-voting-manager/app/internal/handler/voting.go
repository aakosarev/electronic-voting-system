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
	err := h.service.AddRightToVote(ctx, req.UserID, req.VotingID)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
