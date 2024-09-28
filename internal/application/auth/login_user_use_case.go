package auth

import (
	"context"

	userDomain "github/tkuramot/echo-practice/internal/domain/user"
)

type LoginUserUseCase struct {
	userRepo userDomain.Repository
}

func NewLoginUserUseCase(userRepo userDomain.Repository) *LoginUserUseCase {
	return &LoginUserUseCase{
		userRepo: userRepo,
	}
}

type LoginUserUseCaseInputDto struct {
	Email    string
	Password string
}

type LoginUserUseCaseOutputDto struct {
	ID       string
	Email    string
	Nickname string
}

func (uc LoginUserUseCase) Run(ctx context.Context, dto LoginUserUseCaseInputDto) (*LoginUserUseCaseOutputDto, error) {
	user, err := uc.userRepo.FindByEmail(ctx, dto.Email)
	if err != nil {
		return nil, err
	}

	if err := user.Authenticate(dto.Password); err != nil {
		return nil, err
	}
	return &LoginUserUseCaseOutputDto{
		ID:       user.ID(),
		Email:    user.Email(),
		Nickname: user.Nickname(),
	}, nil
}
