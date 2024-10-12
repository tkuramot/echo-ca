package settings

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewEcho() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// TODO replace with actual frontend URL using env variable
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// TODO secret key
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Use(errorMiddleware)

	return e
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func ReturnStatusOK[T any](c echo.Context, body T) error {
	return c.JSON(http.StatusOK, &body)
}

func ReturnStatusCreated[T any](c echo.Context, body T) error {
	return c.JSON(http.StatusCreated, &body)
}

func ReturnStatusNoContent(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

func ReturnStatusBadRequest(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, ErrorResponse{
		Error: err.Error(),
	})
}

func ReturnStatusUnauthorized(c echo.Context, err error) error {
	return c.JSON(http.StatusUnauthorized, ErrorResponse{
		Error: err.Error(),
	})
}

func ReturnStatusForbidden(c echo.Context, err error) error {
	return c.JSON(http.StatusForbidden, ErrorResponse{
		Error: err.Error(),
	})
}

func ReturnStatusNotFound(c echo.Context, err error) error {
	return c.JSON(http.StatusNotFound, ErrorResponse{
		Error: err.Error(),
	})
}

func ReturnStatusInternalServerError(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, ErrorResponse{
		Error: err.Error(),
	})
}
