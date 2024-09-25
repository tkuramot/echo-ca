package user

import (
	"context"

	userDomain "github/tkuramot/echo-practice/internal/domain/user"
)

type FindUserUseCase struct {
	userRepo userDomain.Repository
}

func NewFindUserUseCase(userRepo userDomain.Repository) *FindUserUseCase {
	return &FindUserUseCase{
		userRepo: userRepo,
	}
}

type FindUserUseCaseOutputDto struct {
	ID       string
	Email    string
	Nickname string
}

func (uc FindUserUseCase) Run(ctx context.Context, id string) (*FindUserUseCaseOutputDto, error) {
	user, err := uc.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &FindUserUseCaseOutputDto{
		ID:       user.ID(),
		Email:    user.Email(),
		Nickname: user.Nickname(),
	}, nil
}
