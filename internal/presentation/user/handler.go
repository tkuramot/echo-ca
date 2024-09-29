package user

import (
	"github.com/labstack/echo/v4"

	userApp "github/tkuramot/echo-practice/internal/application/user"
	"github/tkuramot/echo-practice/internal/presentation/settings"
)

type Handler struct {
	findUserUseCase     *userApp.FindUserUseCase
	registerUserUseCase *userApp.RegisterUserUseCase
}

func NewHandler(
	findUserUseCase *userApp.FindUserUseCase,
	registerUserUseCase *userApp.RegisterUserUseCase,
) *Handler {
	return &Handler{
		findUserUseCase:     findUserUseCase,
		registerUserUseCase: registerUserUseCase,
	}
}

func (h *Handler) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	var params registerUserParams
	if err := c.Bind(&params); err != nil {
		return settings.ReturnStatusBadRequest(c, err)
	}

	dto := userApp.RegisterUserUseCaseInputDto{
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
			ID:       user.ID,
			Email:    user.Email,
			Nickname: user.Nickname,
		},
	})
}

func (h *Handler) GetUserByID(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	dto, err := h.findUserUseCase.Run(ctx, id)
	if err != nil {
		return err
	}
	res := getUserResponse{
		User: userResponseModel{
			ID:       dto.ID,
			Email:    dto.Email,
			Nickname: dto.Nickname,
		},
	}
	return settings.ReturnStatusOK(c, res)
}
