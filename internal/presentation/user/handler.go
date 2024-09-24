package user

import (
	"github.com/labstack/echo/v4"

	userApp "github/tkuramot/echo-practice/internal/application/user"
	"github/tkuramot/echo-practice/internal/presentation/settings"
)

type handler struct {
	findUserUseCase *userApp.FindUserUseCase
}

func NewHandler(
	findUserUseCase *userApp.FindUserUseCase,
) *handler {
	return &handler{
		findUserUseCase: findUserUseCase,
	}
}

func (h *handler) GetUserByID(c echo.Context) error {
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
