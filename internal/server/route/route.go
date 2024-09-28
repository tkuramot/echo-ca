package route

import (
	"fmt"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github/tkuramot/echo-practice/internal/presentation/settings"

	authApp "github/tkuramot/echo-practice/internal/application/auth"
	userApp "github/tkuramot/echo-practice/internal/application/user"
	"github/tkuramot/echo-practice/internal/infrastructure/mysql/repository"
	authPre "github/tkuramot/echo-practice/internal/presentation/auth"
	"github/tkuramot/echo-practice/internal/presentation/health_check"
	userPre "github/tkuramot/echo-practice/internal/presentation/user"
)

func InitRoute(e *echo.Echo) {
	v1 := e.Group("/v1")
	v1.GET("/health_check", health_check.HealthCheck)

	{
		authRoute(v1)
		userRoute(v1)
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
	group.GET("/session", func(c echo.Context) error {
		sess, err := session.Get(settings.SessionKey, c)
		if err != nil {
			return err
		}
		return c.String(200, fmt.Sprintf("session: %v", sess.Values["user_id"]))
	})
}

func userRoute(g *echo.Group) {
	userRepo := repository.NewUserRepository()
	h := userPre.NewHandler(
		userApp.NewFindUserUseCase(userRepo),
	)
	group := g.Group("/users")
	group.GET("/:id", h.GetUserByID)
}
