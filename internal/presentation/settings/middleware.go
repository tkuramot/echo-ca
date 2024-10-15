package settings

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	errDomain "github/tkuramot/echo-practice/internal/domain/error"
	sessionDomain "github/tkuramot/echo-practice/internal/domain/session"
	echoRepo "github/tkuramot/echo-practice/internal/infrastructure/echo/repository"
)

func errorMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err == nil {
			return nil
		}

		log.Error(err)

		var domainErr *errDomain.Error
		if errors.As(err, &domainErr) {
			switch {
			case errors.Is(domainErr, sessionDomain.ErrInvalidSession):
				return ReturnStatusUnauthorized(c, err)
			case errors.Is(domainErr, errDomain.ErrNotFound):
				return ReturnStatusNotFound(c, err)
			default:
				return ReturnStatusBadRequest(c, err)
			}
		}
		var echoErr *echo.HTTPError
		if errors.As(err, &echoErr) {
			switch {
			case echoErr.Code == 404:
				return ReturnStatusNotFound(c, errDomain.ErrNotFound)
			}
		}
		return ReturnStatusInternalServerError(c, errDomain.NewError(errDomain.Internal, "サーバーでエラーが発生しました"))
	}
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionRepo := echoRepo.NewSessionRepository(c)
		if err := sessionRepo.Verify(); err != nil {
			return ReturnStatusUnauthorized(c, err)
		}
		return next(c)
	}
}
