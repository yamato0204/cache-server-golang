package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/yamato0204/cache-server-golang/internal/usecase"
)

type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(userUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

func (h *UserHandler) UserRegister(c echo.Context) error {
	ctx := c.Request().Context()
	userId, err := h.userUsecase.RegisterData(ctx)
	fmt.Println(userId)
	if err != nil {
		return err
	}
	return nil
}
