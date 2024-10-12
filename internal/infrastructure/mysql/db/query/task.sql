-- name: TaskFindAll :many
SELECT
    *
FROM tasks
JOIN user_tasks ON user_tasks.task_id = tasks.id
WHERE user_tasks.user_id = ?;

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
SET
    status = ?,
    updated_at = NOW()
WHERE id = ?;