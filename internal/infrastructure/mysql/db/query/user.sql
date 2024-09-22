-- name: UserFindById :one
SELECT
    *
FROM
    users
WHERE
    id = ?;

-- name: UserUpsert :exec
INSERT INTO
    users (
    id,
    email,
    nickname,
    created_at,
    updated_at
)
VALUES
    (
        sqlc.arg(id),
        sqlc.arg(email),
        sqlc.arg(nickname),
        NOW(),
        NOW()
    ) ON DUPLICATE KEY
UPDATE
    email = sqlc.arg(email),
    nickname = sqlc.arg(nickname),
    updated_at = NOW();

-- name: UserFindAll :many
SELECT
    *
FROM
    users;