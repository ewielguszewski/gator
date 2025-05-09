-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT *
FROM feeds;

-- name: GetFeedByURL :one
SELECT f.id, f.name, f.url, f.user_id, f.created_at, u.name as user_name
FROM feeds f
INNER JOIN users u ON f.user_id = u.id
WHERE f.url = $1;
