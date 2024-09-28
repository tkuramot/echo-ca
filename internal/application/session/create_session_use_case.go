package session

import (
	sessionDomain "github/tkuramot/echo-practice/internal/domain/session"
)

type CreateSessionUseCase struct {
}

func NewCreateSessionUseCase() *CreateSessionUseCase {
	return &CreateSessionUseCase{}
}

type CreateSessionUseCaseInputDto struct {
	IsAuthenticated bool
	UserID          string
}

func (uc CreateSessionUseCase) Run(
	sessionRepo sessionDomain.Repository,
	dto CreateSessionUseCaseInputDto,
) error {
	sess := sessionDomain.NewSession(dto.UserID, dto.IsAuthenticated)
	return sessionRepo.Save(sess)
}
