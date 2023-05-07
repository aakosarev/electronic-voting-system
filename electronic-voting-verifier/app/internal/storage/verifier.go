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

func (s *Storage) CheckExistsSigningBlindedAddressRequest(ctx context.Context, userID, votingID int32) (int, error) {
	query := `
		SELECT COUNT(*) FROM signing_blinded_address_request WHERE user_id = $1 AND voting_id = $2;
	`

	var count int

	row := s.client.QueryRow(ctx, query, userID, votingID)
	if err := row.Scan(&count); err != nil {
		return -1, err
	}

	return count, nil
}

func (s *Storage) AddSigningBlindedAddressRequest(ctx context.Context, blindedAddress string, userID, votingID int32) error {
	query := `
		INSERT INTO signing_blinded_address_request(blinded_address, user_id, voting_id)
		VALUES ($1, $2, $3);
	`

	_, err := s.client.Exec(ctx, query, blindedAddress, userID, votingID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) CheckExistsRegisterAddressRequest(ctx context.Context, address string, votingID int32) (int, error) {
	query := `
		SELECT COUNT(*) FROM register_address_request WHERE address = $1 AND voting_id = $2;
	`

	var count int

	row := s.client.QueryRow(ctx, query, address, votingID)
	if err := row.Scan(&count); err != nil {
		return -1, err
	}

	return count, nil
}

func (s *Storage) AddRegisterAddressRequest(ctx context.Context, address string, votingID int32) error {
	query := `
		INSERT INTO register_address_request(address, voting_id)
		VALUES ($1, $2);
	`

	_, err := s.client.Exec(ctx, query, address, votingID)
	if err != nil {
		return err
	}

	return nil
}
