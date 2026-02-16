-- name: CreateFeed :one
INSERT INTO feeds (id, url, created_at, updated_at, name, user_id)
VALUES (
           $1,
    $2,
    $3,
    $4,
    $5,
    $6
    )
    RETURNING *;

-- name: ListFeedsByUserId :many
SELECT * FROM feeds WHERE user_id = $1;

-- name: ListAllFeeds :many
SELECT * FROM feeds;