package storage

import (
	"context"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-manager/internal/model"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-manager/pkg/client/postgresql"
	"github.com/georgysavva/scany/pgxscan"
)

type Storage struct {
	client postgresql.Client
}

func NewStorage(client postgresql.Client) *Storage {
	return &Storage{
		client: client,
	}
}

func (s *Storage) FindVotings(ctx context.Context) ([]*model.Voting, error) {
	query := `
		SELECT id, title, end_time, address, created_on
		FROM voting;
	`

	var votings []*model.Voting

	err := pgxscan.Select(ctx, s.client, &votings, query)
	if err != nil {
		return nil, err
	}

	return votings, nil
}

func (s *Storage) FindVotingByID(ctx context.Context, votingID int32) (*model.Voting, error) {
	query := `
		SELECT id, title, end_time, address, created_on
		FROM voting
		WHERE id = $1;
	`

	var votings []*model.Voting

	err := pgxscan.Select(ctx, s.client, &votings, query, votingID)
	if err != nil {
		return nil, err
	}

	return votings[0], nil
}

func (s *Storage) AddRightToVote(ctx context.Context, userID, votingID int32) error {
	query := `
		INSERT INTO right_to_vote(user_id, voting_id)
		VALUES ($1, $2);
	`

	_, err := s.client.Exec(ctx, query, userID, votingID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) AddVoting(ctx context.Context, title string, endTime int64, address string) (int32, error) {
	query := `
		INSERT INTO voting(title, end_time, address)
		VALUES ($1, $2, $3) RETURNING id;
	`

	var votingID int32

	err := s.client.QueryRow(ctx, query, title, endTime, address).Scan(&votingID)
	if err != nil {
		return -1, err
	}

	return votingID, nil
}

func (s *Storage) FindVotingsAvailableToUserID(ctx context.Context, userID int32) ([]*model.VotingAvailableToUserID, error) {
	query := `
		SELECT rtv.user_id, rtv.voting_id, rtv.created_on, v.title, v.address
		FROM right_to_vote rtv
		INNER JOIN voting v on v.id = rtv.voting_id
		WHERE rtv.user_id = $1
	`

	var votingsAvailableToUserID []*model.VotingAvailableToUserID

	err := pgxscan.Select(ctx, s.client, &votingsAvailableToUserID, query, userID)
	if err != nil {
		return nil, err
	}

	return votingsAvailableToUserID, nil
}
