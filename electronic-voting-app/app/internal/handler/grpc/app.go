package grpc

import (
	"context"
	pb "github.com/aakosarev/electronic-voting-system/contracts/gen/go/electronic-voting-app/v1"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-app/internal/storage"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	userStorage *storage.UserStorage
	pb.UnimplementedVotingAppServer
}

func NewHandler(userStorage *storage.UserStorage, srv pb.UnimplementedVotingAppServer) *Handler {
	return &Handler{
		userStorage:                  userStorage,
		UnimplementedVotingAppServer: srv,
	}
}

func (h *Handler) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*emptypb.Empty, error) {
	err := h.userStorage.Create(ctx, req.GetUsername(), req.GetPasswordHash())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
