package storage

import (
	"context"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-app/pkg/client/postgresql"
	"github.com/georgysavva/scany/pgxscan"
)

type UserStorage struct {
	client postgresql.Client
}

func NewUserStorage(client postgresql.Client) *UserStorage {
	return &UserStorage{
		client: client,
	}
}

func (s *UserStorage) Create(ctx context.Context, username int32, passwordHash string) error {
	query := `
		INSERT INTO "user"(id, password_hash)
		VALUES ($1, $2);
	`

	_, err := s.client.Exec(ctx, query, username, passwordHash)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserStorage) GetPasswordHashByUserID(ctx context.Context, userID int32) (string, error) {
	query := `
		SELECT password_hash
		FROM "user"
		WHERE id = $1;
	`

	var passwordHash string

	err := pgxscan.Get(ctx, s.client, &passwordHash, query, userID)
	if err != nil {
		return "", err
	}

	return passwordHash, nil
}

func (s *UserStorage) AddDetails(ctx context.Context, userID int32, passwordHash, email, name, surname string) error {
	query := `
		UPDATE "user"
		SET password_hash = $2, email = $3, name = $4, surname = $5, force_enter_details = $6
		WHERE id = $1
	`

	_, err := s.client.Exec(ctx, query, userID, passwordHash, email, name, surname, false)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserStorage) GetForceEnterDetailsByUserID(ctx context.Context, userID int32) (bool, error) {
	query := `
		SELECT force_enter_details
		FROM "user"
		WHERE id = $1;
	`

	var forceEnterDetails bool

	err := pgxscan.Get(ctx, s.client, &forceEnterDetails, query, userID)
	if err != nil {
		return false, err
	}

	return forceEnterDetails, nil
}
