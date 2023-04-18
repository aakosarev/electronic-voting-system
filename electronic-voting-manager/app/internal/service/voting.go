package service

import (
	"context"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-manager/internal/eth/voting"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-manager/internal/model"
	"github.com/ethereum/go-ethereum/ethclient"
	"time"
)

type storage interface {
	FindVotings(ctx context.Context) ([]*model.Voting, error)
	FindVotingByID(ctx context.Context, votingID int32) (*model.Voting, error)
	AddRightToVote(ctx context.Context, userID, votingID int32) error
	AddVoting(ctx context.Context, title string, endTime int64, address string) error
	FindVotingsAvailableToUser(ctx context.Context, userID int32) ([]*model.VotingAvailableToUser, error)
}

type Service struct {
	storage         storage
	contractSession *voting.ContractSession
	ethclient       *ethclient.Client
}

func NewService(storage storage, contractSession *voting.ContractSession, ethclient *ethclient.Client) *Service {
	return &Service{
		storage:         storage,
		contractSession: contractSession,
		ethclient:       ethclient,
	}
}

func (s *Service) CreateVoting(ctx context.Context, title string, options []string, endTime time.Time) error {
	address, err := voting.CreateVoting(s.contractSession, s.ethclient, title, endTime, options)
	if err != nil {
		return err
	}

	err = s.storage.AddVoting(ctx, title, endTime.Unix(), address)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetAllVotings(ctx context.Context) ([]*model.Voting, error) {
	votings, err := s.storage.FindVotings(ctx)
	if err != nil {
		return nil, err
	}

	return votings, nil
}

func (s *Service) AddRightToVote(ctx context.Context, userID, votingID int32) error {
	err := s.storage.AddRightToVote(ctx, userID, votingID)
	if err != nil {
		return err
	}
	return nil
}
