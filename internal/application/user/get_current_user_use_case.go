package user

import (
	"context"
	sessionDomain "github/tkuramot/echo-practice/internal/domain/session"
	userDomain "github/tkuramot/echo-practice/internal/domain/user"
)

type GetCurrentUserUseCase struct {
	userRepo userDomain.Repository
}

type GetCurrentUseCaseOutputDto struct {
	ID       string
	Email    string
	Nickname string
}

func NewGetCurrentUseCase(userRepo userDomain.Repository) *GetCurrentUserUseCase {
	return &GetCurrentUserUseCase{
		userRepo: userRepo,
	}
}

func (uc GetCurrentUserUseCase) Run(
	ctx context.Context,
	sessionRepo sessionDomain.Repository,
) (*GetCurrentUseCaseOutputDto, error) {
	userID, err := sessionRepo.UserID()
	if err != nil {
		return nil, err
	}
	user, err := uc.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &GetCurrentUseCaseOutputDto{
		ID:       user.ID(),
		Email:    user.Email(),
		Nickname: user.Nickname(),
	}, nil
}
