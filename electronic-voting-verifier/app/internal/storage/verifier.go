package storage

import (
	"context"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-verifier/pkg/client/postgresql"
)

type Storage struct {
	client postgresql.Client
}

func NewStorage(client postgresql.Client) *Storage {
	return &Storage{
		client: client,
	}
}

func (s *Storage) CheckExistsBlindedTokenSigningRequest(ctx context.Context, userID, votingID int32) (int, error) {
	query := `
		SELECT COUNT(*) FROM blinded_token_signing_request WHERE user_id = $1 AND voting_id = $2;
	`

	var count int

	row := s.client.QueryRow(ctx, query, userID, votingID)
	if err := row.Scan(&count); err != nil {
		return -1, err
	}

	return count, nil
}

func (s *Storage) AddBlindedTokenSigningRequest(ctx context.Context, userID, votingID int32, blindedTokenHash string) error {
	query := `
		INSERT INTO blinded_token_signing_request(user_id, voting_id, blinded_token_hash)
		VALUES ($1, $2, $3);
	`

	_, err := s.client.Exec(ctx, query, userID, votingID, blindedTokenHash)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) CheckExistsRegisterAddressToVotingRequest(ctx context.Context, votingID int32, address string) (int, error) {
	query := `
		SELECT COUNT(*) FROM register_address_to_voting_by_signed_token_request WHERE voting_id = $1 AND address = $2;
	`

	var count int

	row := s.client.QueryRow(ctx, query, votingID, address)
	if err := row.Scan(&count); err != nil {
		return -1, err
	}

	return count, nil
}

func (s *Storage) AddRegisterAddressToVotingRequest(ctx context.Context, address string, votingID int32, signedTokenHash string) error {
	query := `
		INSERT INTO register_address_to_voting_by_signed_token_request(address, voting_id, signed_token_hash)
		VALUES ($1, $2, $3);
	`

	_, err := s.client.Exec(ctx, query, address, votingID, signedTokenHash)
	if err != nil {
		return err
	}

	return nil
}
