package storage

import (
	"context"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-verifier/pkg/client/postgresql"
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

func (s *Storage) CheckExistsRegistrationRequest(ctx context.Context, userID, votingID int32) (bool, error) {
	query := `
		SELECT EXISTS (SELECT * FROM registration_request WHERE user_id = $1 AND voting_id = $2);
	`
	var exists bool

	err := pgxscan.Get(ctx, s.client, &exists, query, userID, votingID)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (s *Storage) AddRegistrationRequest(ctx context.Context, userID, votingID int32, blindedTokenHash string) error {
	query := `
		INSERT INTO registration_request(user_id, voting_id, blinded_token_hash)
		VALUES ($1, $2, $3);
	`

	_, err := s.client.Exec(ctx, query, userID, votingID, blindedTokenHash)
	if err != nil {
		return err
	}

	return nil
}
