-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id) VALUES (
        $1,
        $2,
        $3, 
        $4,
        $5
    )
    RETURNING *
) SELECT 
    inserted_feed_follow.*,
    users.name as user_name,
    feed.name as feed_name
FROM inserted_feed_follow
INNER JOIN users
ON users.id = inserted_feed_follow.user_id
INNER JOIN feed
ON feed.id = inserted_feed_follow.feed_id;

-- name: GetFeedFollowsForUser :many
SELECT 
    feed_follows.user_id,
    feed_follows.feed_id,
    users.name as user_name,
    feed.name as feed_name
FROM feed_follows
INNER JOIN users
ON users.id = feed_follows.user_id
INNER JOIN feed
ON feed.id = feed_follows.feed_id
WHERE feed_follows.user_id = $1;