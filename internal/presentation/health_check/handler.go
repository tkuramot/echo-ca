package health_check

import (
	"github.com/labstack/echo/v4"

	"github/tkuramot/echo-practice/internal/presentation/settings"
)

func HealthCheck(c echo.Context) error {
	res := HealthCheckResponse{
		Status: "ok",
	}
	return settings.ReturnStatusOK(c, res)
}
