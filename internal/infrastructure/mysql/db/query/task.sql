-- name: TaskFindAll :many
SELECT
    *
FROM tasks
WHERE (sqlc.narg('user_id') IS NULL OR user_id = sqlc.narg('user_id')) AND
      (sqlc.narg('task_status') IS NULL OR status = sqlc.narg('task_status'));

-- name: TaskFindById :one
SELECT
    *
FROM tasks
WHERE user_id = ? AND id = ?;

-- name: TaskFindByStatus :many
SELECT
    *
FROM tasks
WHERE user_id = ? AND status = ?;

-- name: TaskInsert :exec
INSERT INTO
    tasks (
    id,
    title,
    description,
    status,
    created_at,
    updated_at
)
VALUES (
    sqlc.arg(id),
    sqlc.arg(title),
    sqlc.arg(description),
    sqlc.arg(status),
    NOW(),
    NOW()
);

-- name: TaskUpdateStatus :exec
UPDATE
    tasks
SET
    status = ?
WHERE user_id = ? AND id = ?;

-- name: TaskUpdate :exec
UPDATE
    tasks
SET
    title = ?,
    description = ?,
    status = ?
WHERE user_id = ? AND id = ?;