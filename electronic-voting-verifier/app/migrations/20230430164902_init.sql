-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS blinded_token_signing_request (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    voting_id INTEGER NOT NULL,
    blinded_token_hash TEXT UNIQUE NOT NULL,
    created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id, voting_id)
);

CREATE TABLE IF NOT EXISTS register_address_to_voting_by_signed_token_request (
    id SERIAL PRIMARY KEY,
    address VARCHAR(42) NOT NULL,
    voting_id INTEGER NOT NULL,
    signed_token_hash TEXT UNIQUE NOT NULL,
    UNIQUE (voting_id, address)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS blinded_token_signing_request;
DROP TABLE IF EXISTS register_address_to_voting_by_signed_token_request;
-- +goose StatementEnd
