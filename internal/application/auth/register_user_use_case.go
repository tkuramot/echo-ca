package auth

import (
	"context"

	userDomain "github/tkuramot/echo-practice/internal/domain/user"
)

type RegisterUserUseCase struct {
	userRepo userDomain.Repository
}

func NewRegisterUserUseCase(
	userRepo userDomain.Repository,
) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		userRepo: userRepo,
	}
}

type RegisterUserUseCaseDto struct {
	Email    string
	Nickname string
	Password string
}

func (uc *RegisterUserUseCase) Run(ctx context.Context, dto RegisterUserUseCaseDto) (*userDomain.User, error) {
	u, err := userDomain.NewUser(dto.Email, dto.Nickname, dto.Password)
	if err != nil {
		return nil, err
	}
	if err := uc.userRepo.Save(ctx, u); err != nil {
		return nil, err
	}

	user, err := uc.userRepo.FindByEmail(ctx, dto.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
