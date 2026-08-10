-- name: Createpost :one
INSERT INTO 
    posts(
        id,
        created_at,
        updated_at,
        title,
        url,
        description,
        published_at,
        feed_id
    )
VALUES(
    $1, NOW(), NOW(), $2, $3, $4, $5, $6
)
RETURNING *;
