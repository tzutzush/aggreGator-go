-- +goose Up
CREATE TABLE feeds (
                       id UUID PRIMARY KEY,
                       url TEXT UNIQUE,
                       created_at TIMESTAMP,
                       updated_at TIMESTAMP,
                       name TEXT UNIQUE NOT NULL,
                       user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;