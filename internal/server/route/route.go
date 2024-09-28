package route

import (
	"github.com/labstack/echo/v4"

	authApp "github/tkuramot/echo-practice/internal/application/auth"
	userApp "github/tkuramot/echo-practice/internal/application/user"
	"github/tkuramot/echo-practice/internal/infrastructure/mysql/repository"
	authPre "github/tkuramot/echo-practice/internal/presentation/auth"
	"github/tkuramot/echo-practice/internal/presentation/health_check"
	"github/tkuramot/echo-practice/internal/presentation/settings"
	userPre "github/tkuramot/echo-practice/internal/presentation/user"
)

func InitRoute(e *echo.Echo) {
	v1 := e.Group("/v1")
	{
		v1.GET("/health_check", health_check.HealthCheck)
		authRoute(v1)
	}

	protectedV1 := e.Group("/v1")
	protectedV1.Use(settings.AuthMiddleware)
	{
		protectedUserRoute(protectedV1)
	}
}

func authRoute(g *echo.Group) {
	userRepo := repository.NewUserRepository()
	h := authPre.NewHandler(
		authApp.NewRegisterUserUseCase(userRepo),
		authApp.NewLoginUserUseCase(userRepo),
	)
	group := g.Group("/auth")
	group.POST("/register", h.RegisterUser)
	group.POST("/login", h.LoginUser)
}

func protectedUserRoute(g *echo.Group) {
	userRepo := repository.NewUserRepository()
	h := userPre.NewHandler(
		userApp.NewFindUserUseCase(userRepo),
	)
	group := g.Group("/users")
	group.GET("/:id", h.GetUserByID)
}
