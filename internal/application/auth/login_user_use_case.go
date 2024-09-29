package auth

import (
	"context"
	sessionDomain "github/tkuramot/echo-practice/internal/domain/session"
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

func (uc LoginUserUseCase) Run(
	ctx context.Context,
	sessionRepo sessionDomain.Repository,
	dto LoginUserUseCaseInputDto,
) (*LoginUserUseCaseOutputDto, error) {
	user, err := uc.userRepo.FindByEmail(ctx, dto.Email)
	if err != nil {
		return nil, err
	}

	if err := user.Authenticate(dto.Password); err != nil {
		return nil, err
	}

	sess := sessionDomain.NewSession(user.ID(), true)
	if err := sessionRepo.Save(sess); err != nil {
		return nil, err
	}
	return &LoginUserUseCaseOutputDto{
		ID:       user.ID(),
		Email:    user.Email(),
		Nickname: user.Nickname(),
	}, nil
}
