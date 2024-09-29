package session

import (
	errDomain "github/tkuramot/echo-practice/internal/domain/error"
)

type Session struct {
	isAuthenticated bool
	rememberMe      bool
	userID          string
}

const (
	ID                 = "session_id"
	KeyIsAuthenticated = "is_authenticated"
	KeyUserID          = "user_id"
	KeyRememberMe      = "remember_me"
)

var ErrInvalidSession = errDomain.NewError("無効なセッションです。")

func NewSession(userID string, isAuthenticated, rememberMe bool) *Session {
	return &Session{
		isAuthenticated: isAuthenticated,
		rememberMe:      rememberMe,
		userID:          userID,
	}
}

func (s *Session) IsAuthenticated() bool {
	return s.isAuthenticated
}

func (s *Session) UserID() string {
	return s.userID
}

func (s *Session) RememberMe() bool {
	return s.rememberMe
}
