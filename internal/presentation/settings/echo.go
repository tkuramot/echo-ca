package settings

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewEcho() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())
	return e
}

func ReturnStatusOK[T any](c echo.Context, body T) error {
	return c.JSON(http.StatusOK, &body)
}

func ReturnStatusCreated[T any](c echo.Context, body T) error {
	return c.JSON(http.StatusCreated, &body)
}

func ReturnStatusNoContent(c echo.Context) error {
	return c.JSON(http.StatusNoContent, nil)
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
