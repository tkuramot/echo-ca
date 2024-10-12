-- name: TaskFindAll :many
SELECT
    *
FROM tasks
JOIN user_tasks ON user_tasks.task_id = tasks.id
WHERE (sqlc.narg('user_id') IS NULL OR user_tasks.user_id = sqlc.narg('user_id')) AND
      (sqlc.narg('task_status') IS NULL OR tasks.status = sqlc.narg('task_status'));

-- name: TaskFindById :one
SELECT
    *
FROM tasks
JOIN user_tasks ON user_tasks.task_id = tasks.id
WHERE user_tasks.user_id = ? AND tasks.id = ?;

-- name: TaskFindByStatus :many
SELECT
    *
FROM tasks
JOIN user_tasks ON user_tasks.task_id = tasks.id
WHERE user_tasks.user_id = ? AND tasks.status = ?;

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
JOIN user_tasks ON user_tasks.task_id = tasks.id
SET
    tasks.status = ?
WHERE user_tasks.user_id = ? AND tasks.id = ?;

-- name: TaskUpdate :exec
UPDATE
    tasks
JOIN user_tasks ON user_tasks.task_id = tasks.id
SET
    tasks.title = ?,
    tasks.description = ?,
    tasks.status = ?
WHERE user_tasks.user_id = ? AND tasks.id = ?;