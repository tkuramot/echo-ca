// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: task.sql

package dbgen

import (
	"context"
	"database/sql"
)

const taskFindAll = `-- name: TaskFindAll :many
SELECT
    id, title, description, status, created_at, updated_at, user_id
FROM tasks
WHERE (? IS NULL OR user_id = ?) AND
      (? IS NULL OR status = ?)
`

type TaskFindAllParams struct {
	UserID     sql.NullString  `json:"user_id"`
	TaskStatus NullTasksStatus `json:"task_status"`
}

func (q *Queries) TaskFindAll(ctx context.Context, arg TaskFindAllParams) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, taskFindAll,
		arg.UserID,
		arg.UserID,
		arg.TaskStatus,
		arg.TaskStatus,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Task{}
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const taskFindById = `-- name: TaskFindById :one
SELECT
    id, title, description, status, created_at, updated_at, user_id
FROM tasks
WHERE user_id = ? AND id = ?
`

type TaskFindByIdParams struct {
	UserID string `json:"user_id"`
	ID     string `json:"id"`
}

func (q *Queries) TaskFindById(ctx context.Context, arg TaskFindByIdParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, taskFindById, arg.UserID, arg.ID)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
	)
	return i, err
}

const taskFindByStatus = `-- name: TaskFindByStatus :many
SELECT
    id, title, description, status, created_at, updated_at, user_id
FROM tasks
WHERE user_id = ? AND status = ?
`

type TaskFindByStatusParams struct {
	UserID string      `json:"user_id"`
	Status TasksStatus `json:"status"`
}

func (q *Queries) TaskFindByStatus(ctx context.Context, arg TaskFindByStatusParams) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, taskFindByStatus, arg.UserID, arg.Status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Task{}
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const taskInsert = `-- name: TaskInsert :exec
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
    ?,
    ?,
    ?,
    ?,
    NOW(),
    NOW()
)
`

type TaskInsertParams struct {
	ID          string      `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Status      TasksStatus `json:"status"`
}

func (q *Queries) TaskInsert(ctx context.Context, arg TaskInsertParams) error {
	_, err := q.db.ExecContext(ctx, taskInsert,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.Status,
	)
	return err
}

const taskUpdate = `-- name: TaskUpdate :exec
UPDATE
    tasks
SET
    title = ?,
    description = ?,
    status = ?
WHERE user_id = ? AND id = ?
`

type TaskUpdateParams struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Status      TasksStatus `json:"status"`
	UserID      string      `json:"user_id"`
	ID          string      `json:"id"`
}

func (q *Queries) TaskUpdate(ctx context.Context, arg TaskUpdateParams) error {
	_, err := q.db.ExecContext(ctx, taskUpdate,
		arg.Title,
		arg.Description,
		arg.Status,
		arg.UserID,
		arg.ID,
	)
	return err
}

const taskUpdateStatus = `-- name: TaskUpdateStatus :exec
UPDATE
    tasks
SET
    status = ?
WHERE user_id = ? AND id = ?
`

type TaskUpdateStatusParams struct {
	Status TasksStatus `json:"status"`
	UserID string      `json:"user_id"`
	ID     string      `json:"id"`
}

func (q *Queries) TaskUpdateStatus(ctx context.Context, arg TaskUpdateStatusParams) error {
	_, err := q.db.ExecContext(ctx, taskUpdateStatus, arg.Status, arg.UserID, arg.ID)
	return err
}
