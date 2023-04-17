-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS voting (
    id SERIAL PRIMARY KEY,
    title VARCHAR(500) UNIQUE NOT NULL,
    end_time BIGINT NOT NULL,
    address VARCHAR(500) UNIQUE NOT NULL,
    created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS right_to_vote (
    user_id INTEGER,
    voting_id INTEGER REFERENCES voting(id) ON DELETE CASCADE,
    created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, voting_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS voting;
DROP TABLE IF EXISTS right_to_vote;
-- +goose StatementEnd