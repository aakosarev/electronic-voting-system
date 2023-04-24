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

func (s *UserStorage) Create(ctx context.Context, username int32, password string) error {
	query := `
		INSERT INTO "user"(username, password)
		VALUES ($1, $2);
	`

	_, err := s.client.Exec(ctx, query, username, password)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserStorage) GetPassByUsername(ctx context.Context, username int32) (string, error) {
	query := `
		SELECT password
		FROM "user"
		WHERE username = $1;
	`

	var password string

	err := pgxscan.Get(ctx, s.client, &password, query, username)
	if err != nil {
		return "", err
	}

	return password, nil
}

func (s *UserStorage) Update(ctx context.Context, username int32, password, email, firstName, secondName string) error {
	query := `
		UPDATE "user"
		SET password = $2, email = $3, first_name = $4, second_name = $5, force_enter_details = $6
		WHERE username = $1
	`

	_, err := s.client.Exec(ctx, query, username, password, email, firstName, secondName, false)
	if err != nil {
		return err
	}

	return nil
}
