package handler

import (
	"errors"
	"net/http"

	"Go_Practice/api/schema"
	"Go_Practice/usecase/interactor"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserUC interactor.IUserUseCase
}

func NewUserHandler(userUC interactor.IUserUseCase) *UserHandler {
	return &UserHandler{UserUC: userUC}
}

func (h *UserHandler) FindById(c echo.Context) error {

	var id string
	if err := echo.PathParamsBinder(c).MustString("user-id", &id).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.UserUC.FindById(id)
	if err != nil {
		switch {
		case errors.Is(err, interactor.ErrUserNotFound):
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, schema.UserResFromEntity(res))
}


