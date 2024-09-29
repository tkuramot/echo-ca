package auth

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"

	authApp "github/tkuramot/echo-practice/internal/application/auth"
	echoRepo "github/tkuramot/echo-practice/internal/infrastructure/echo/repository"
	"github/tkuramot/echo-practice/internal/presentation/settings"
)

type Handler struct {
	loginUserUseCase *authApp.LoginUserUseCase
}

func NewHandler(
	loginUserUseCase *authApp.LoginUserUseCase,
) *Handler {
	return &Handler{
		loginUserUseCase: loginUserUseCase,
	}
}

func (h *Handler) LoginUser(c echo.Context) error {
	ctx := c.Request().Context()
	var params loginUserParams
	if err := c.Bind(&params); err != nil {
		return settings.ReturnStatusBadRequest(c, err)
	}

	dto := authApp.LoginUserUseCaseInputDto{
		Email:    params.Email,
		Password: params.Password,
	}
	sessionRepo := echoRepo.NewSessionRepository(c)
	user, err := h.loginUserUseCase.Run(ctx, sessionRepo, dto)
	if err != nil {
		return settings.ReturnStatusUnauthorized(c, err)
	}

	return settings.ReturnStatusOK(c, loginUserResponse{
		User: userResponseModel{
			ID:       user.ID,
			Email:    user.Email,
			Nickname: user.Nickname,
		},
	})
}

func (h *Handler) LogoutUser(c echo.Context) error {
	sess, err := session.Get(settings.SessionKey, c)
	if err != nil {
		return settings.ReturnStatusInternalServerError(c, err)
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return settings.ReturnStatusInternalServerError(c, err)
	}
	return c.NoContent(http.StatusOK)
}
