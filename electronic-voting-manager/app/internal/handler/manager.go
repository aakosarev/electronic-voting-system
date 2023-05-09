package handler

import (
	"context"
	pb "github.com/aakosarev/electronic-voting-system/contracts/gen/go/electronic-voting-manager/v1"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-manager/internal/config"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-manager/internal/eth/voting"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-manager/internal/storage"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Handler struct {
	storage         *storage.Storage
	contractSession *voting.ContractSession
	ethclient       *ethclient.Client
	pb.UnimplementedVotingManagerServer
}

func NewHandler(storage *storage.Storage, contractSession *voting.ContractSession,
	ethclient *ethclient.Client, srv pb.UnimplementedVotingManagerServer) *Handler {
	return &Handler{
		storage:                          storage,
		contractSession:                  contractSession,
		ethclient:                        ethclient,
		UnimplementedVotingManagerServer: srv,
	}
}

func (h *Handler) CreateVoting(ctx context.Context, req *pb.CreateVotingRequest) (*pb.CreateVotingResponse, error) {
	address, err := voting.CreateVoting(h.contractSession, h.ethclient, req.GetVotingTitle(), req.GetEndTime().AsTime(), req.GetVotingOptions())
	if err != nil {
		return nil, err
	}

	votingID, err := h.storage.AddVoting(ctx, req.GetVotingTitle(), req.GetEndTime().AsTime().Unix(), address)
	if err != nil {
		return nil, err
	}

	return &pb.CreateVotingResponse{VotingID: votingID}, nil
}

func (h *Handler) GetAllVotings(ctx context.Context, _ *emptypb.Empty) (*pb.GetAllVotingsResponse, error) {
	votings, err := h.storage.FindVotings(ctx)
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
	err := h.storage.AddRightToVote(ctx, req.GetUserID(), req.GetVotingID())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (h *Handler) GetVotingsAvailableToUserID(ctx context.Context, req *pb.GetVotingsAvailableToUserIDRequest) (*pb.GetVotingsAvailableToUserIDResponse, error) {
	votingsAvailableToUserID, err := h.storage.FindVotingsAvailableToUserID(ctx, req.GetUserID())
	if err != nil {
		return nil, err
	}

	pbVotingsAvailableToUserID := make([]*pb.VotingAvailableToUserID, len(votingsAvailableToUserID))
	for i, v := range votingsAvailableToUserID {
		pbVotingsAvailableToUserID[i] = &pb.VotingAvailableToUserID{
			UserID:        v.UserID,
			VotingID:      v.VotingID,
			CreatedOn:     timestamppb.New(v.CreatedOn),
			VotingTitle:   v.Title,
			VotingAddress: v.Address,
		}
	}

	return &pb.GetVotingsAvailableToUserIDResponse{VotingsAvailableToUserID: pbVotingsAvailableToUserID}, nil
}

func (h *Handler) RegisterAddressToVoting(ctx context.Context, req *pb.RegisterAddressToVotingRequest) (*emptypb.Empty, error) {
	v, err := h.storage.FindVotingByID(ctx, req.GetVotingID())
	if err != nil {
		return nil, err
	}

	cfg := config.GetConfig()
	err = voting.RegisterAddressToVoting(h.contractSession, h.ethclient, cfg, common.HexToAddress(v.Address), common.HexToAddress(req.GetAddress()))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (h *Handler) GetVotingInformation(ctx context.Context, req *pb.GetVotingInformationRequest) (*pb.GetVotingInformationResponse, error) {
	v, err := h.storage.FindVotingByID(ctx, req.GetVotingID())
	if err != nil {
		return nil, err
	}

	votingInformation, err := voting.GetVotingInformation(h.contractSession, h.ethclient, common.HexToAddress(v.Address))
	if err != nil {
		return nil, err
	}

	pbOptions := make(map[int64]*pb.Option)

	for index, option := range votingInformation.Options {
		pbOptions[index] = &pb.Option{
			Name:        option.Name,
			NumberVotes: option.NumberVotes,
		}
	}

	return &pb.GetVotingInformationResponse{
		Title:                  votingInformation.Title,
		NumberRegisteredVoters: votingInformation.NumberRegisteredVoters,
		EndTime:                timestamppb.New(votingInformation.EndTime),
		Options:                pbOptions,
	}, nil
}
