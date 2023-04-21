-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "user" (
    username INTEGER PRIMARY KEY,
    password VARCHAR(50) NOT NULL,
    email VARCHAR(50) NULL,
    first_name VARCHAR(50) NULL,
    second_name VARCHAR(50) NULL,
    force_enter_details BOOL NOT NULL DEFAULT TRUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "user";
-- +goose StatementEnd
