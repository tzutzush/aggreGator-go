-- name: CreateFeed :one
INSERT INTO feeds (id, url, created_at, updated_at, name)
VALUES (
           $1,
    $2,
    $3,
    $4,
    $5
    )
    RETURNING *;

-- name: GetFeedsByUserId :many
SELECT * FROM feeds WHERE user_id = $1;

-- name: GetAllFeeds :many
SELECT * FROM feeds;

-- name: GetFeedByURL :one
SELECT * FROM feeds WHERE url = $1;