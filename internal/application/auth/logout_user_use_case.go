package auth

import "github/tkuramot/echo-practice/internal/domain/session"

type LogoutUserUseCase struct{}

func NewLogoutUserUseCase() *LogoutUserUseCase {
	return &LogoutUserUseCase{}
}

func (uc LogoutUserUseCase) Run(sessionRepo session.Repository) error {
	return sessionRepo.Delete()
}
