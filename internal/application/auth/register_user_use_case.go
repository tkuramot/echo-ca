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

type RegisterUserUseCaseInputDto struct {
	Email    string
	Nickname string
	Password string
}

type RegisterUserUseCaseOutputDto struct {
	ID       string
	Email    string
	Nickname string
}

func (uc *RegisterUserUseCase) Run(ctx context.Context, dto RegisterUserUseCaseInputDto) (*RegisterUserUseCaseOutputDto, error) {
	u, err := userDomain.NewUser(dto.Email, dto.Nickname, dto.Password)
	if err != nil {
		return nil, err
	}
	if err := uc.userRepo.Save(ctx, u); err != nil {
		return nil, err
	}
	return &RegisterUserUseCaseOutputDto{
		ID:       u.ID(),
		Email:    u.Email(),
		Nickname: u.Nickname(),
	}, nil
}
