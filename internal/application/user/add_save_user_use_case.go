package user

import (
	"context"

	userDomain "github/tkuramot/echo-practice/internal/domain/user"
)

type SaveUserUseCase struct {
	userRepo userDomain.UserRepository
}

func NewSaveUserUseCase(userRepo userDomain.UserRepository) *SaveUserUseCase {
	return &SaveUserUseCase{userRepo: userRepo}
}

type SaveUserUseCaseDto struct {
	Email    string
	Nickname string
}

func (uc *SaveUserUseCase) Run(ctx context.Context, dto SaveUserUseCaseDto) error {
	user, err := userDomain.NewUser(dto.Email, dto.Nickname)
	if err != nil {
		return err
	}
	return uc.userRepo.Save(ctx, user)
}
