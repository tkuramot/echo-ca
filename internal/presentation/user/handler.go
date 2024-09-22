package user

import (
	"errors"
	"github.com/labstack/echo/v4"
	userApp "github/tkuramot/echo-practice/internal/application/user"
	errDomain "github/tkuramot/echo-practice/internal/domain/error"
	"github/tkuramot/echo-practice/internal/presentation/settings"
)

type handler struct {
	findUserUseCase *userApp.FindUserUseCase
	saveUserUseCase *userApp.SaveUserUseCase
}

func NewHandler(
	findUserUseCase *userApp.FindUserUseCase,
	saveUserUseCase *userApp.SaveUserUseCase,
) *handler {
	return &handler{
		findUserUseCase: findUserUseCase,
		saveUserUseCase: saveUserUseCase,
	}
}

func (h *handler) GetUserByID(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	dto, err := h.findUserUseCase.Run(ctx, id)
	if err != nil {
		if errors.Is(err, errDomain.NotFountError) {
			return settings.ReturnStatusNotFound(c, err)
		}
		return settings.ReturnStatusInternalServerError(c, err)
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
