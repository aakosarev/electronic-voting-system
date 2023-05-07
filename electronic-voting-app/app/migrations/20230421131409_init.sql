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
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "user";
-- +goose StatementEnd
