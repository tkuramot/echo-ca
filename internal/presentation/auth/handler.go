package auth

import (
	"github.com/labstack/echo/v4"
	authApp "github/tkuramot/echo-practice/internal/application/auth"
	"github/tkuramot/echo-practice/internal/presentation/settings"
)

type handler struct {
	registerUserUseCase *authApp.RegisterUserUseCase
}

func NewHandler(registerUserUseCase *authApp.RegisterUserUseCase) *handler {
	return &handler{
		registerUserUseCase: registerUserUseCase,
	}
}

func (h *handler) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	var params registerUserParams
	if err := c.Bind(&params); err != nil {
		return settings.ReturnStatusBadRequest(c, err)
	}

	dto := authApp.RegisterUserUseCaseDto{
		Email:    params.Email,
		Nickname: params.Nickname,
		Password: params.Password,
	}
	user, err := h.registerUserUseCase.Run(ctx, dto)
	if err != nil {
		return err
	}

	return settings.ReturnStatusCreated(c, registerUserResponse{
		User: userResponseModel{
			ID:       user.ID(),
			Email:    user.Email(),
			Nickname: user.Nickname(),
		},
	})
}
