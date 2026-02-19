-- +goose Up
CREATE TABLE feeds (
                       id UUID PRIMARY KEY,
                       url TEXT NOT NULL UNIQUE,
                       created_at TIMESTAMP NOT NULL,
                       updated_at TIMESTAMP NOT NULL,
                       name TEXT UNIQUE NOT NULL,
                       last_fetched_at TIMESTAMP NOT NULL,
                       user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;