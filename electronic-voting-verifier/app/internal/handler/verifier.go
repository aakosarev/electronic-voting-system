package handler

import (
	"context"
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

func (h *Handler) GenerateRSAKeyPairForVotingID(ctx context.Context, req *pbvv.GenerateRSAKeyPairForVotingIDRequest) (*emptypb.Empty, error) {
	err := h.keystorage.GenerateRSAKeyPairForVotingID(req.GetVotingID())
	if err != nil {
		return nil, fmt.Errorf("generate rsa key pair for voting id [%d]: %v", req.GetVotingID(), err)
	}

	return &emptypb.Empty{}, nil
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

func (h *Handler) SignBlindedAddress(ctx context.Context, req *pbvv.SignBlindedAddressRequest) (*pbvv.SignBlindedAddressResponse, error) {
	getVotingsAvailableToUserIDReqToVM := &pbvm.GetVotingsAvailableToUserIDRequest{UserID: req.GetUserID()}

	getVotingsAvailableToUserIDRespFromVM, err := h.votingManagerClient.GetVotingsAvailableToUserID(ctx, getVotingsAvailableToUserIDReqToVM)
	if err != nil {
		return nil, err
	}

	pbAvailableVotings := getVotingsAvailableToUserIDRespFromVM.GetVotingsAvailableToUserID()

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

	countRequests, err := h.storage.CheckExistsSigningBlindedAddressRequest(ctx, req.GetUserID(), req.GetVotingID())
	if err != nil {
		return nil, err
	}

	if countRequests > 0 {
		return nil, errors.New("such a token signing request already exists")
	}

	signedBlindedAddress, err := h.keystorage.SignBlindedAddress(req.GetBlindedAddress(), req.GetVotingID())
	if err != nil {
		return nil, err
	}

	err = h.storage.AddSigningBlindedAddressRequest(ctx, hex.EncodeToString(req.GetBlindedAddress()), req.GetUserID(), req.GetVotingID())
	if err != nil {
		return nil, err
	}

	return &pbvv.SignBlindedAddressResponse{SignedBlindedAddress: signedBlindedAddress}, nil
}

func (h *Handler) RegisterAddress(ctx context.Context, req *pbvv.RegisterAddressRequest) (*emptypb.Empty, error) {
	ok, err := h.keystorage.VerifySignature(req.GetSignedAddress(), req.GetAddress(), req.GetVotingID())
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("signature not verified")
	}

	countRequests, err := h.storage.CheckExistsRegisterAddressRequest(ctx, req.GetAddress(), req.GetVotingID())
	if err != nil {
		return nil, err
	}

	if countRequests > 0 {
		return nil, errors.New("such an address registration request already exists")
	}

	registerAddressToVotingReqToVM := &pbvm.RegisterAddressToVotingRequest{
		VotingID: req.GetVotingID(),
		Address:  req.GetAddress(),
	}

	_, err = h.votingManagerClient.RegisterAddressToVoting(ctx, registerAddressToVotingReqToVM)
	if err != nil {
		return nil, err
	}

	err = h.storage.AddRegisterAddressRequest(ctx, req.GetAddress(), req.GetVotingID())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (h *Handler) GetRegistrationStatuses(ctx context.Context, req *pbvv.GetRegistrationStatusesRequest) (*pbvv.GetRegistrationStatusesResponse, error) {
	addresses := req.GetAddresses()
	statuses := make(map[int32]bool, len(addresses))

	for votingID, address := range addresses {
		count, err := h.storage.CheckExistsRegisterAddressRequest(ctx, address, votingID)
		if err != nil {
			return nil, err
		}
		if count == 1 {
			statuses[votingID] = true
		} else {
			statuses[votingID] = false
		}
	}

	return &pbvv.GetRegistrationStatusesResponse{Statuses: statuses}, nil
}
