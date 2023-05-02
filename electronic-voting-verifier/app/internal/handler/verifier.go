package handler

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	pbvm "github.com/aakosarev/electronic-voting-system/contracts/gen/go/electronic-voting-manager/v1"
	pbvv "github.com/aakosarev/electronic-voting-system/contracts/gen/go/electronic-voting-verifier/v1"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-verifier/internal/keystorage"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-verifier/internal/storage"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	keystorage *keystorage.KeyStorage
	storage    *storage.Storage
	pbvv.UnimplementedVotingVerifierServer
	votingManagerClient pbvm.VotingManagerClient
}

func NewHandler(keystorage *keystorage.KeyStorage, storage *storage.Storage, srv pbvv.UnimplementedVotingVerifierServer, connVotingManager *grpc.ClientConn) *Handler {
	return &Handler{
		keystorage:                        keystorage,
		storage:                           storage,
		UnimplementedVotingVerifierServer: srv,
		votingManagerClient:               pbvm.NewVotingManagerClient(connVotingManager),
	}
}

func (h *Handler) GetPublicKeyForVotingID(ctx context.Context, req *pbvv.GetPublicKeyForVotingIDRequest) (*pbvv.GetPublicKeyForVotingIDResponse, error) {
	if err := h.keystorage.GenerateKeyPairForVotingID(req.GetVotingID()); err != nil {
		return nil, err
	}

	publicKeyBytes, err := h.keystorage.GetPublicKeyBytesForVotingID(req.GetVotingID())
	if err != nil {
		return nil, err
	}

	return &pbvv.GetPublicKeyForVotingIDResponse{PublicKeyBytes: publicKeyBytes}, nil
}

func (h *Handler) SignBlindedToken(ctx context.Context, req *pbvv.SignBlindedTokenRequest) (*pbvv.SignBlindedTokenResponse, error) {
	reqToVotingManager := &pbvm.GetVotingsAvailableToUserRequest{UserID: req.GetUserID()}

	resp, err := h.votingManagerClient.GetVotingsAvailableToUser(ctx, reqToVotingManager)
	if err != nil {
		return nil, err
	}

	pbAvailableVotings := resp.GetVotingsAvailableToUser()

	isAllowedVote := false
	for _, av := range pbAvailableVotings {
		if av.GetVotingID() == req.GetVotingID() {
			isAllowedVote = true
			break
		}
	}

	if !isAllowedVote {
		return nil, fmt.Errorf("the user %d does not have the right to vote in the voting %d", req.GetUserID(), req.GetVotingID())
	}

	requestAlreadyExists, err := h.storage.CheckExistsBlindedTokenSigningRequest(ctx, req.GetUserID(), req.GetVotingID())
	if err != nil {
		return nil, err
	}

	if requestAlreadyExists {
		return nil, errors.New("such a token signing request already exists")
	}

	signedBlindedToken, err := h.keystorage.SignBlindedToken(req.GetBlindedToken(), req.GetVotingID())
	if err != nil {
		//return nil, err
		return nil, errors.New("1123123")
	}

	blindedTokenHash := sha256.Sum256([]byte(req.GetBlindedToken()))

	err = h.storage.AddBlindedTokenSigningRequest(ctx, req.GetUserID(), req.GetVotingID(), string(blindedTokenHash[:]))
	if err != nil {
		return nil, err
	}

	return &pbvv.SignBlindedTokenResponse{SignedBlindedToken: signedBlindedToken}, nil
}

func (h *Handler) RegisterAddressToVoting(ctx context.Context, req *pbvv.RegisterAddressToVotingBySignedTokenRequest) (*emptypb.Empty, error) {

	ok, err := h.keystorage.VerifySignature(req.GetSignedToken(), req.GetToken(), req.GetVotingID())
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("signature not verified")
	}

	requestAlreadyExists, err := h.storage.CheckExistsRegisterAddressToVotingRequest(ctx, req.GetVotingID(), req.GetAddress())
	if err != nil {
		return nil, err
	}

	if requestAlreadyExists {
		return nil, errors.New("such an address registration request already exists")
	}

	reqToVotingManager := &pbvm.RegisterAddressToVotingRequest{
		VotingID: req.GetVotingID(),
		Address:  req.GetAddress(),
	}

	_, err = h.votingManagerClient.RegisterAddressToVoting(ctx, reqToVotingManager)
	if err != nil {
		return nil, err
	}

	signedTokenHash := sha256.Sum256([]byte(req.GetSignedToken()))

	err = h.storage.AddRegisterAddressToVotingRequest(ctx, req.GetAddress(), req.GetVotingID(), string(signedTokenHash[:]))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
