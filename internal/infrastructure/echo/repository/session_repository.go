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
	rememberMe, ok := sess.Values[sessionDomain.KeyRememberMe].(bool)
	if !ok {
		return nil, sessionDomain.ErrInvalidSession
	}

	return sessionDomain.NewSession(
		userID,
		isAuthenticated,
		rememberMe,
	), nil
}

func (r *SessionRepository) UserID() (string, error) {
	sess, err := session.Get(sessionDomain.ID, r.ctx)
	if err != nil {
		return "", err
	}

	userID, ok := sess.Values[sessionDomain.KeyUserID].(string)
	if !ok {
		return "", sessionDomain.ErrInvalidSession
	}
	return userID, nil
}

func (r *SessionRepository) Delete() error {
	sess, err := session.Get(sessionDomain.ID, r.ctx)
	if err != nil {
		return err
	}
	sess.Values = make(map[interface{}]interface{})
	sess.Options.MaxAge = -1
	return sess.Save(r.ctx.Request(), r.ctx.Response())
}

func (r *SessionRepository) Save(s *sessionDomain.Session) error {
	sess, err := session.Get(sessionDomain.ID, r.ctx)
	if err != nil {
		return err
	}

	if s.RememberMe() {
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
	} else {
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   0,
			HttpOnly: true,
		}
	}
	sess.Values[sessionDomain.KeyUserID] = s.UserID()
	sess.Values[sessionDomain.KeyIsAuthenticated] = s.IsAuthenticated()
	sess.Values[sessionDomain.KeyRememberMe] = s.RememberMe()
	return sess.Save(r.ctx.Request(), r.ctx.Response())
}

func (r *SessionRepository) Verify() error {
	sess, err := session.Get(sessionDomain.ID, r.ctx)
	if err != nil {
		return err
	}

	if _, ok := sess.Values[sessionDomain.KeyUserID]; !ok {
		return sessionDomain.ErrInvalidSession
	}
	if _, ok := sess.Values[sessionDomain.KeyIsAuthenticated]; !ok {
		return sessionDomain.ErrInvalidSession
	}
	return nil
}
