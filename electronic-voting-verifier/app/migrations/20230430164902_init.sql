-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS registration_request (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    voting_id INTEGER NOT NULL,
    blinded_token_hash TEXT UNIQUE NOT NULL,
    created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id, voting_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS registration_request;
-- +goose StatementEnd
