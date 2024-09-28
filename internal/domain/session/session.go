package session

import (
	errDomain "github/tkuramot/echo-practice/internal/domain/error"
	"net/http"
)

type Session struct {
	isAuthenticated bool
	userID          string
}

type Config struct {
	domain   string
	httpOnly bool
	maxAge   int
	path     string
	sameSite http.SameSite
	secure   bool
}

const (
	ID                 = "session_id"
	KeyIsAuthenticated = "is_authenticated"
	KeyUserID          = "user_id"
)

var ErrInvalidSession = errDomain.NewError("無効なセッションです。")

func NewSession(userID string, isAuthenticated bool) *Session {
	return &Session{
		isAuthenticated: isAuthenticated,
		userID:          userID,
	}
}

func (s *Session) IsAuthenticated() bool {
	return s.isAuthenticated
}

func (s *Session) UserID() string {
	return s.userID
}
