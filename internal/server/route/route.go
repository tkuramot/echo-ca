package route

import (
	"github.com/labstack/echo/v4"

	userApp "github/tkuramot/echo-practice/internal/application/user"
	"github/tkuramot/echo-practice/internal/infrastructure/mysql/repository"
	"github/tkuramot/echo-practice/internal/presentation/health_check"
	"github/tkuramot/echo-practice/internal/presentation/settings"
	userPre "github/tkuramot/echo-practice/internal/presentation/user"
)

func InitRoute(e *echo.Echo) {
	e.Use(settings.ErrorHandler)
	v1 := e.Group("/v1")
	v1.GET("/health_check", health_check.HealthCheck)

	{
		userRoute(v1)
	}
}

func userRoute(r *echo.Group) {
	userRepo := repository.NewUserRepository()
	h := userPre.NewHandler(
		userApp.NewFindUserUseCase(userRepo),
		userApp.NewSaveUserUseCase(userRepo),
	)
	group := r.Group("/users")
	group.GET("/:id", h.GetUserByID)
}
