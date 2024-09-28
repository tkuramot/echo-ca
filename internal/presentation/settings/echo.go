package settings

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	SessionKey       = "session_id"
	SessionUserIDKey = "user_id"
)

func NewEcho() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// TODO secret key
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Use(errorHandler)

	return e
}

func ReturnStatusOK[T any](c echo.Context, body T) error {
	return c.JSON(http.StatusOK, &body)
}

func ReturnStatusCreated[T any](c echo.Context, body T) error {
	return c.JSON(http.StatusCreated, &body)
}

func ReturnStatusBadRequest(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, err)
}

func ReturnStatusUnauthorized(c echo.Context, err error) error {
	return c.JSON(http.StatusUnauthorized, err)
}

func ReturnStatusForbidden(c echo.Context, err error) error {
	return c.JSON(http.StatusForbidden, err)
}

func ReturnStatusNotFound(c echo.Context, err error) error {
	return c.JSON(http.StatusNotFound, err)
}

func ReturnStatusInternalServerError(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, err)
}
