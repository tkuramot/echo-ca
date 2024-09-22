package user

import (
	"context"

	userDomain "github/tkuramot/echo-practice/internal/domain/user"
)

type FindUserUseCase struct {
	userRepo userDomain.UserRepository
}

func NewFindUserUseCase(userRepo userDomain.UserRepository) *FindUserUseCase {
	return &FindUserUseCase{
		userRepo: userRepo,
	}
}

type FindUseCaseDto struct {
	ID       string
	Email    string
	Nickname string
}

func (uc FindUserUseCase) Run(ctx context.Context, id string) (*FindUseCaseDto, error) {
	user, err := uc.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &FindUseCaseDto{
		ID:       user.ID(),
		Email:    user.Email(),
		Nickname: user.Nickname(),
	}, nil
}
