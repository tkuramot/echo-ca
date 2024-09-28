package settings

import (
	"errors"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	errDomain "github/tkuramot/echo-practice/internal/domain/error"
)

func errorMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err == nil {
			return nil
		}

		var domainErr *errDomain.Error
		if errors.As(err, &domainErr) {
			switch {
			case errors.Is(domainErr, errDomain.ErrNotFound):
				return ReturnStatusNotFound(c, domainErr)
			default:
				return ReturnStatusBadRequest(c, domainErr)
			}
		}
		return ReturnStatusInternalServerError(c, err)
	}
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get(SessionKey, c)
		if err != nil {
			return ReturnStatusUnauthorized(c, err)
		}

		userID := sess.Values[SessionUserIDKey]
		if userID == nil {
			return ReturnStatusUnauthorized(c, errors.New("unauthorized"))
		}

		return next(c)
	}
}
