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
		v1.GET("/health", health_check.HealthCheck)
		authRoute(v1)
		userRoute(v1)
	}

	protectedV1 := e.Group("/v1")
	protectedV1.Use(settings.AuthMiddleware)
	{
		protectedAuthRoute(protectedV1)
		protectedUserRoute(protectedV1)
	}
}

func authRoute(g *echo.Group) {
	userRepo := repository.NewUserRepository()
	h := authPre.NewHandler(
		authApp.NewLoginUserUseCase(userRepo),
		authApp.NewLogoutUserUseCase(),
	)
	group := g.Group("/auth")
	group.POST("/login", h.LoginUser)
}

func protectedAuthRoute(g *echo.Group) {
	userRepo := repository.NewUserRepository()
	h := authPre.NewHandler(
		authApp.NewLoginUserUseCase(userRepo),
		authApp.NewLogoutUserUseCase(),
	)
	group := g.Group("/auth")
	group.POST("/logout", h.LogoutUser)
}

func userRoute(g *echo.Group) {
	userRepo := repository.NewUserRepository()
	h := userPre.NewHandler(
		userApp.NewFindUserUseCase(userRepo),
		userApp.NewRegisterUserUseCase(userRepo),
	)
	group := g.Group("/users")
	group.POST("/", h.RegisterUser)
}

func protectedUserRoute(g *echo.Group) {
	userRepo := repository.NewUserRepository()
	h := userPre.NewHandler(
		userApp.NewFindUserUseCase(userRepo),
		userApp.NewRegisterUserUseCase(userRepo),
	)
	group := g.Group("/users")
	group.GET("/:id", h.GetUserByID)
}
