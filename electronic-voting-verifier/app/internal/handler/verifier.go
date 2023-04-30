package handler

import (
	"context"
	pbvm "github.com/aakosarev/electronic-voting-system/contracts/gen/go/electronic-voting-manager/v1"
	pbvv "github.com/aakosarev/electronic-voting-system/contracts/gen/go/electronic-voting-verifier/v1"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-verifier/internal/keystorage"
	"google.golang.org/grpc"
)

type Handler struct {
	keystorage *keystorage.KeyStorage
	pbvv.UnimplementedVotingVerifierServer
	votingManagerClient pbvm.VotingManagerClient
}

func NewHandler(keystorage *keystorage.KeyStorage, srv pbvv.UnimplementedVotingVerifierServer, connVotingManager *grpc.ClientConn) *Handler {
	return &Handler{
		keystorage:                        keystorage,
		UnimplementedVotingVerifierServer: srv,
		votingManagerClient:               pbvm.NewVotingManagerClient(connVotingManager),
	}
}

func (h *Handler) GetPublicKeyForVotingID(ctx context.Context, req *pbvv.GetPublicKeyForVotingIDRequest) (*pbvv.GetPublicKeyForVotingIDResponse, error) {

	publicKeyBytes, err := h.keystorage.GetPublicKeyBytesForVotingID(req.GetVotingID())
	if err != nil {
		return nil, err
	}

	return &pbvv.GetPublicKeyForVotingIDResponse{PublicKeyBytes: publicKeyBytes}, nil
}
