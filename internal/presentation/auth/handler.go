package auth

import (
	"github.com/labstack/echo/v4"
	"net/http"

	authApp "github/tkuramot/echo-practice/internal/application/auth"
	echoRepo "github/tkuramot/echo-practice/internal/infrastructure/echo/repository"
	"github/tkuramot/echo-practice/internal/presentation/settings"
)

type Handler struct {
	loginUserUseCase  *authApp.LoginUserUseCase
	logoutUserUseCase *authApp.LogoutUserUseCase
}

func NewHandler(
	loginUserUseCase *authApp.LoginUserUseCase,
	logoutUserUseCase *authApp.LogoutUserUseCase,
) *Handler {
	return &Handler{
		loginUserUseCase:  loginUserUseCase,
		logoutUserUseCase: logoutUserUseCase,
	}
}

func (h *Handler) LoginUser(c echo.Context) error {
	ctx := c.Request().Context()
	var params loginUserParams
	if err := c.Bind(&params); err != nil {
		return settings.ReturnStatusBadRequest(c, err)
	}

	dto := authApp.LoginUserUseCaseInputDto{
		Email:      params.Email,
		Password:   params.Password,
		RememberMe: params.RememberMe,
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
	sessionRepo := echoRepo.NewSessionRepository(c)
	if err := h.logoutUserUseCase.Run(sessionRepo); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
