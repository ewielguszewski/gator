-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, feed_id, user_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
) 
SELECT inserted_feed_follow.*, f.name as feed_name, u.name as user_name
FROM inserted_feed_follow
INNER JOIN feeds f ON inserted_feed_follow.feed_id = f.id
INNER JOIN users u ON inserted_feed_follow.user_id = u.id;

-- name: GetFeedFollowsForUser :many
SELECT ff.*, f.name as feed_name, u.name as user_name
FROM feed_follows ff
INNER JOIN feeds f ON ff.feed_id = f.id
INNER JOIN users u ON ff.user_id = u.id
WHERE ff.user_id = $1;

-- name: DeleteFeedFollow :one
WITH feed_to_unfollow AS (
    SELECT id FROM feeds WHERE url = $1
)
DELETE FROM feed_follows ff
WHERE feed_id = (SELECT id FROM feed_to_unfollow)
AND ff.user_id = $2
RETURNING *;

-- name: MarkFeedFetched :one
UPDATE feeds
SET last_fetched_at = NOW(), updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;