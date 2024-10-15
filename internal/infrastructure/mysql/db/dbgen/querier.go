// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package dbgen

import (
	"context"
)

type Querier interface {
	TaskFindAll(ctx context.Context, arg TaskFindAllParams) ([]Task, error)
	TaskFindById(ctx context.Context, arg TaskFindByIdParams) (Task, error)
	TaskFindByStatus(ctx context.Context, arg TaskFindByStatusParams) ([]Task, error)
	TaskInsert(ctx context.Context, arg TaskInsertParams) error
	TaskUpdate(ctx context.Context, arg TaskUpdateParams) error
	TaskUpdateStatus(ctx context.Context, arg TaskUpdateStatusParams) error
	UserFindAll(ctx context.Context) ([]User, error)
	UserFindByEmail(ctx context.Context, email string) (User, error)
	UserFindById(ctx context.Context, id string) (User, error)
	UserInsert(ctx context.Context, arg UserInsertParams) error
	UserUpsert(ctx context.Context, arg UserUpsertParams) error
}

var _ Querier = (*Queries)(nil)
