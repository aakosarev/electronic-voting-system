package handler

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
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
	if err := h.keystorage.GenerateRSAKeyPairForVotingID(req.GetVotingID()); err != nil {
		return nil, err
	}

	publicKeyBytes, err := h.keystorage.GetPublicKeyBytesForVotingID(req.GetVotingID())
	if err != nil {
		return nil, err
	}

	return &pbvv.GetPublicKeyForVotingIDResponse{PublicKeyBytes: publicKeyBytes}, nil
}

func (h *Handler) SignBlindedPublicKey(ctx context.Context, req *pbvv.SignBlindedPublicKeyRequest) (*pbvv.SignBlindedPublicKeyResponse, error) {
	getVotingsAvailableToUserIDReqToVM := &pbvm.GetVotingsAvailableToUserIDRequest{UserID: req.GetUserID()}

	getVotingsAvailableToUserIDRespFromVM, err := h.votingManagerClient.(ctx, reqToVotingManager)
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

	countRequests, err := h.storage.CheckExistsBlindedTokenSigningRequest(ctx, req.GetUserID(), req.GetVotingID())
	if err != nil {
		return nil, err
	}

	if countRequests > 0 {
		return nil, errors.New("such a token signing request already exists")
	}

	signedBlindedToken, err := h.keystorage.SignBlindedToken(req.GetBlindedToken(), req.GetVotingID())
	if err != nil {
		return nil, err
	}

	blindedTokenHash := sha256.Sum256(req.GetBlindedToken())

	err = h.storage.AddBlindedTokenSigningRequest(ctx, req.GetUserID(), req.GetVotingID(), hex.EncodeToString(blindedTokenHash[:]))
	if err != nil {
		return nil, err
	}

	return &pbvv.SignBlindedTokenResponse{SignedBlindedToken: signedBlindedToken}, nil
}

func (h *Handler) RegisterAddressToVotingBySignedToken(ctx context.Context, req *pbvv.RegisterAddressToVotingBySignedTokenRequest) (*emptypb.Empty, error) {

	ok, err := h.keystorage.VerifySignature(req.GetSignedToken(), req.GetToken(), req.GetVotingID())
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("signature not verified")
	}

	countRequests, err := h.storage.CheckExistsRegisterAddressToVotingRequest(ctx, req.GetVotingID(), req.GetAddress())
	if err != nil {
		return nil, err
	}

	if countRequests > 0 {
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

	err = h.storage.AddRegisterAddressToVotingRequest(ctx, req.GetAddress(), req.GetVotingID(), hex.EncodeToString(signedTokenHash[:]))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (h *Handler) GenerateRSAKeyPairForVotingID(ctx context.Context, req *pbvv.GenerateRSAKeyPairForVotingIDRequest) (*emptypb.Empty, error) {
	err := h.keystorage.GenerateRSAKeyPairForVotingID(req.GetVotingID())
	if err != nil {
		return nil, fmt.Errorf("generate rsa key pair for voting id [%d]: %v", req.GetVotingID(), err)
	}

	return &emptypb.Empty{}, nil
}
