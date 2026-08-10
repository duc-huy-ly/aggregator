-- name: CreateUser :one
INSERT INTO
    users (id, created_at, updated_at, name)
VALUES
    ($1, NOW(), NOW(), $2) RETURNING *;

-- name: GetUser :one
SELECT
    *
FROM
    users
WHERE
    name = $1;

-- name: Reset :exec
DELETE FROM
    users;

-- name: GetUsers :many
SELECT
    *
FROM
    users;

-- name: GetFeedPlusUser :many
SELECT
    feeds.name,
    feeds.url,
    users.name AS creator_name
FROM
    feeds
    JOIN users ON feeds.user_id = users.id;