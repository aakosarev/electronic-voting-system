-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS blinded_address_signing_request (
    blinded_address TEXT PRIMARY KEY,
    user_id INTEGER NOT NULL,
    voting_id INTEGER NOT NULL,
    created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id, voting_id)
);

CREATE TABLE IF NOT EXISTS register_address_request (
    address VARCHAR(42) PRIMARY KEY ,
    voting_id INTEGER NOT NULL,
    created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (voting_id, address)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS blinded_address_signing_request;
DROP TABLE IF EXISTS register_address_request;
-- +goose StatementEnd
