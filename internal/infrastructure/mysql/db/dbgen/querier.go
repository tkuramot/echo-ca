// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package dbgen

import (
	"context"
)

type Querier interface {
	UserFindAll(ctx context.Context) ([]User, error)
	UserFindByEmail(ctx context.Context, email string) (User, error)
	UserFindById(ctx context.Context, id string) (User, error)
	UserUpsert(ctx context.Context, arg UserUpsertParams) error
}

var _ Querier = (*Queries)(nil)
