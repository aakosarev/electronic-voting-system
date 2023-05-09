-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "user" (
    id INTEGER PRIMARY KEY,
    password_hash VARCHAR(60) NOT NULL,
    email VARCHAR(256) NULL,
    name VARCHAR(50) NULL,
    surname VARCHAR(50) NULL,
    force_enter_details BOOL NOT NULL DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS demo_registrations (
    private_key TEXT PRIMARY KEY,
    user_id INTEGER NOT NULL,
    voting_id INTEGER NOT NULL,
    UNIQUE (user_id, voting_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "user";
DROP TABLE IF EXISTS demo_registrations;
-- +goose StatementEnd
