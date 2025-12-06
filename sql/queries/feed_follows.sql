-- name: CreateFeedFollow :one
WITH inserted_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT inf.*, u.name AS user_name, f.name AS feed_name
FROM inserted_follow inf INNER JOIN users u 
                         ON inf.user_id = u.id
                         INNER JOIN feeds f
                         ON inf.feed_id = f.id;

-- name: GetFeedFollowsForUser :many
SELECT ff.*, u.name AS user_name, f.name AS feed_name
FROM feed_follows ff INNER JOIN users u ON ff.user_id = u.id
                     INNER JOIN feeds f ON ff.feed_id = f.id
WHERE u.id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
WHERE feed_id = $1 AND user_id = $2;