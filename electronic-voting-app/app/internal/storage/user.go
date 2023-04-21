package storage

import (
	"context"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-app/pkg/client/postgresql"
)

type UserStorage struct {
	client postgresql.Client
}

func NewStorage(client postgresql.Client) *UserStorage {
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
