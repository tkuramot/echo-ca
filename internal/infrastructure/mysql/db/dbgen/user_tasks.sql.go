// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user_tasks.sql

package dbgen

import (
	"context"
)

const userTaskInsert = `-- name: UserTaskInsert :exec
INSERT INTO
    user_tasks (
    user_id,
    task_id
)
VALUES (
    ?,
    ?
)
`

type UserTaskInsertParams struct {
	UserID string `json:"user_id"`
	TaskID string `json:"task_id"`
}

func (q *Queries) UserTaskInsert(ctx context.Context, arg UserTaskInsertParams) error {
	_, err := q.db.ExecContext(ctx, userTaskInsert, arg.UserID, arg.TaskID)
	return err
}
