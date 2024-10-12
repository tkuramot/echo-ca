-- name: UserTaskInsert :exec
INSERT INTO
    user_tasks (
    user_id,
    task_id,
    created_at,
    updated_at
)
VALUES (
    sqlc.arg(user_id),
    sqlc.arg(task_id),
    NOW(),
    NOW()
);