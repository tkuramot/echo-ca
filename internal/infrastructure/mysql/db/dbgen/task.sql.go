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
    id, title, description, status, created_at, updated_at, user_id, task_id
FROM tasks
JOIN user_tasks ON user_tasks.task_id = tasks.id
WHERE user_tasks.user_id = ?
`

type TaskFindAllRow struct {
	ID          string       `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Status      TasksStatus  `json:"status"`
	CreatedAt   sql.NullTime `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
	UserID      string       `json:"user_id"`
	TaskID      string       `json:"task_id"`
}

func (q *Queries) TaskFindAll(ctx context.Context, userID string) ([]TaskFindAllRow, error) {
	rows, err := q.db.QueryContext(ctx, taskFindAll, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TaskFindAllRow{}
	for rows.Next() {
		var i TaskFindAllRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.TaskID,
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
    id, title, description, status, created_at, updated_at
FROM tasks
WHERE id = ?
`

func (q *Queries) TaskFindById(ctx context.Context, id string) (Task, error) {
	row := q.db.QueryRowContext(ctx, taskFindById, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const taskFindByStatus = `-- name: TaskFindByStatus :many
SELECT
    id, title, description, status, created_at, updated_at, user_id, task_id
FROM tasks
JOIN user_tasks ON user_tasks.task_id = tasks.id
WHERE user_tasks.user_id = ? AND tasks.status = ?
`

type TaskFindByStatusParams struct {
	UserID string      `json:"user_id"`
	Status TasksStatus `json:"status"`
}

type TaskFindByStatusRow struct {
	ID          string       `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Status      TasksStatus  `json:"status"`
	CreatedAt   sql.NullTime `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
	UserID      string       `json:"user_id"`
	TaskID      string       `json:"task_id"`
}

func (q *Queries) TaskFindByStatus(ctx context.Context, arg TaskFindByStatusParams) ([]TaskFindByStatusRow, error) {
	rows, err := q.db.QueryContext(ctx, taskFindByStatus, arg.UserID, arg.Status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TaskFindByStatusRow{}
	for rows.Next() {
		var i TaskFindByStatusRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.TaskID,
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

const taskUpdateStatus = `-- name: TaskUpdateStatus :exec
UPDATE
    tasks
SET
    status = ?,
    updated_at = NOW()
WHERE id = ?
`

type TaskUpdateStatusParams struct {
	Status TasksStatus `json:"status"`
	ID     string      `json:"id"`
}

func (q *Queries) TaskUpdateStatus(ctx context.Context, arg TaskUpdateStatusParams) error {
	_, err := q.db.ExecContext(ctx, taskUpdateStatus, arg.Status, arg.ID)
	return err
}
