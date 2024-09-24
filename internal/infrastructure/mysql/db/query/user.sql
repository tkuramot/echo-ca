-- name: UserFindByEmail :one
SELECT
    *
FROM
    users
WHERE
    email = ?;

-- name: UserFindById :one
SELECT
    *
FROM
    users
WHERE
    id = ?;

-- name: UserInsert :exec
INSERT INTO
    users (
    id,
    email,
    nickname,
    password_digest,
    created_at,
    updated_at
)
VALUES (
    sqlc.arg(id),
    sqlc.arg(email),
    sqlc.arg(nickname),
    sqlc.arg(password_digest),
    NOW(),
    NOW()
);

-- name: UserUpsert :exec
INSERT INTO
    users (
    id,
    email,
    nickname,
    password_digest,
    created_at,
    updated_at
)
VALUES (
    sqlc.arg(id),
    sqlc.arg(email),
    sqlc.arg(nickname),
    sqlc.arg(password_digest),
    NOW(),
    NOW()
) ON DUPLICATE KEY
UPDATE
    email = sqlc.arg(email),
    nickname = sqlc.arg(nickname),
    password_digest = sqlc.arg(password_digest),
    updated_at = NOW();

-- name: UserFindAll :many
SELECT
    *
FROM
    users;