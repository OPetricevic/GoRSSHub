-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;


-- name: GetFeeds :many
Select * From feeds;

-- name: GetNextFeedToFetch :one
Select * From feeds
Order by last_fetched_at ASC nulls first
limit 1;

-- name: MarkFeedAsFetched :one
UPDATE feeds
SET last_fetched_at = NOW(),
updated_at = NOW()
WHERE id = $1
RETURNING *;