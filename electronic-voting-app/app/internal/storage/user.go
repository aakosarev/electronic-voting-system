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

func (s *UserStorage) DemoAddRegistration(ctx context.Context, privateKey string, userID, votingID int32) error {
	query := `
		INSERT INTO demo_registrations(private_key, user_id, voting_id)
		VALUES ($1, $2, $3);
	`

	_, err := s.client.Exec(ctx, query, privateKey, userID, votingID)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserStorage) DemoGetPrivateKeys(ctx context.Context, userID int32, votingIDs []int32) (map[int32]string, error) {
	query := `
		SELECT * FROM demo_registrations 
		WHERE user_id = $1 AND voting_id = ANY($2);
	`

	rows, err := s.client.Query(ctx, query, userID, votingIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	privateKeys := make(map[int32]string)

	for rows.Next() {
		var privateKey string
		var usrID int32
		var votingID int32
		err = rows.Scan(&privateKey, &usrID, &votingID)
		if err != nil {
			return nil, err
		}

		privateKeys[votingID] = privateKey
	}

	return privateKeys, nil
}

func (s *UserStorage) DemoGetPrivateKey(ctx context.Context, userID, votingID int32) (string, error) {
	query := `
		SELECT private_key FROM demo_registrations 
		WHERE user_id = $1 AND voting_id = $2;
	`

	var privateKey string

	err := pgxscan.Get(ctx, s.client, &privateKey, query, userID, votingID)
	if err != nil {
		return "", err
	}

	return privateKey, nil
}
