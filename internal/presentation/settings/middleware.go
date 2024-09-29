package settings

import (
	"errors"
	"github.com/labstack/echo/v4"
	errDomain "github/tkuramot/echo-practice/internal/domain/error"
	echoRepo "github/tkuramot/echo-practice/internal/infrastructure/echo/repository"
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
		sessionRepo := echoRepo.NewSessionRepository(c)
		if err := sessionRepo.Verify(); err != nil {
			return ReturnStatusUnauthorized(c, err)
		}
		return next(c)
	}
}
