package repository

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	sessionDomain "github/tkuramot/echo-practice/internal/domain/session"
)

type SessionRepository struct {
	ctx echo.Context
}

func NewSessionRepository(ctx echo.Context) *SessionRepository {
	return &SessionRepository{ctx: ctx}
}

func (r *SessionRepository) Get() (*sessionDomain.Session, error) {
	sess, err := session.Get(sessionDomain.ID, r.ctx)
	if err != nil {
		return nil, err
	}

	userID, ok := sess.Values[sessionDomain.KeyUserID].(string)
	if !ok {
		return nil, sessionDomain.ErrInvalidSession
	}
	isAuthenticated, ok := sess.Values[sessionDomain.KeyIsAuthenticated].(bool)
	if !ok {
		return nil, sessionDomain.ErrInvalidSession
	}

	return sessionDomain.NewSession(
		userID,
		isAuthenticated,
	), nil
}

func (r *SessionRepository) Save(s *sessionDomain.Session) error {
	sess, err := session.Get(sessionDomain.ID, r.ctx)
	if err != nil {
		return err
	}

	// TODO persist session when remember me is checked
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
	}
	sess.Values[sessionDomain.KeyUserID] = s.UserID()
	sess.Values[sessionDomain.KeyIsAuthenticated] = s.IsAuthenticated()
	return sess.Save(r.ctx.Request(), r.ctx.Response())
}
