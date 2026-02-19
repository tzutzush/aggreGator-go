-- name: CreateFeed :one
INSERT INTO feeds (id, url, created_at, updated_at, name, last_fetched_at, user_id)
VALUES (
           $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
    )
    RETURNING *;

-- name: GetFeedsByUserId :many
SELECT * FROM feeds WHERE user_id = $1;

-- name: GetAllFeeds :many
SELECT * FROM feeds;

-- name: GetFeedByURL :one
SELECT * FROM feeds WHERE url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = NOW(), updated_at = NOW()
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;