-- name: CreateFeed :one
INSERT INTO feed(id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feed;

-- name: GetFeedFromURL :one
SELECT *
FROM feed
WHERE feed.url = $1;

-- name: MarkFeedFetched :exec
UPDATE feed
SET created_at = $2 AND last_fetched_at = $2
WHERE feed.id = $1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feed
ORDER BY feed.last_fetched_at DESC NULLS FIRST 
LIMIT 1;