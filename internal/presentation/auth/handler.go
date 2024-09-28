package auth

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"

	authApp "github/tkuramot/echo-practice/internal/application/auth"
	"github/tkuramot/echo-practice/internal/presentation/settings"
)

type Handler struct {
	registerUserUseCase *authApp.RegisterUserUseCase
	loginUserUseCase    *authApp.LoginUserUseCase
}

func NewHandler(
	registerUserUseCase *authApp.RegisterUserUseCase,
	loginUserUseCase *authApp.LoginUserUseCase,
) *Handler {
	return &Handler{
		registerUserUseCase: registerUserUseCase,
		loginUserUseCase:    loginUserUseCase,
	}
}

func (h *Handler) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	var params registerUserParams
	if err := c.Bind(&params); err != nil {
		return settings.ReturnStatusBadRequest(c, err)
	}

	dto := authApp.RegisterUserUseCaseInputDto{
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

func (h *Handler) LoginUser(c echo.Context) error {
	ctx := c.Request().Context()
	var params loginUserParams
	if err := c.Bind(&params); err != nil {
		// TODO: error message
		return settings.ReturnStatusBadRequest(c, err)
	}

	dto := authApp.LoginUserUseCaseInputDto{
		Email:    params.Email,
		Password: params.Password,
	}
	user, err := h.loginUserUseCase.Run(ctx, dto)
	if err != nil {
		// TODO error message
		return settings.ReturnStatusUnauthorized(c, err)
	}

	sess, err := session.Get(settings.SessionKey, c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values[settings.SessionUserIDKey] = user.ID
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return settings.ReturnStatusInternalServerError(c, err)
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
