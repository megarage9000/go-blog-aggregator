
-- name: CreatePost :one
INSERT INTO post (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
) ON CONFLICT(url) DO NOTHING

RETURNING *;

-- name: GetPostsForUser :many
WITH user_feed_follows AS (
    SELECT feed_follows.feed_id
    FROM feed_follows
    WHERE user_id = $1
)
SELECT post.* FROM post
INNER JOIN user_feed_follows
ON user_feed_follows.feed_id = post.feed_id
LIMIT $2;