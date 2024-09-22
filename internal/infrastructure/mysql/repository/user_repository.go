package repository

import (
	"context"
	"database/sql"
	"errors"
	errDomain "github/tkuramot/echo-practice/internal/domain/error"
	"github/tkuramot/echo-practice/internal/domain/user"
	"github/tkuramot/echo-practice/internal/infrastructure/mysql/db"
	"github/tkuramot/echo-practice/internal/infrastructure/mysql/db/dbgen"
)

type userRepository struct{}

func NewUserRepository() user.UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindByID(ctx context.Context, id string) (*user.User, error) {
	query := db.GetQuery(ctx)
	u, err := query.UserFindById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errDomain.NotFountErr
		}
		return nil, err
	}
	ud, err := user.Reconstruct(
		u.ID,
		u.Email,
		u.Nickname,
	)
	if err != nil {
		return nil, err
	}
	return ud, nil
}

func (r *userRepository) Save(ctx context.Context, u *user.User) error {
	query := db.GetQuery(ctx)
	if err := query.UserUpsert(ctx, dbgen.UserUpsertParams{
		ID:       u.ID(),
		Email:    u.Email(),
		Nickname: u.Nickname(),
	}); err != nil {
		return err
	}
	return nil
}

func (r *userRepository) FindAll(ctx context.Context) ([]*user.User, error) {
	query := db.GetQuery(ctx)
	us, err := query.UserFindAll(ctx)
	if err != nil {
		return nil, err
	}
	var users []*user.User
	for _, u := range us {
		ud, err := user.Reconstruct(
			u.ID,
			u.Email,
			u.Nickname,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, ud)
	}
	return users, nil
}
