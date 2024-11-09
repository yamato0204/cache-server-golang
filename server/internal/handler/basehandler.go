package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/yamato0204/cache-server-golang/internal/handler/api"
)

type BaseHandler struct {
	userHandler *api.UserHandler
}

func NewBaseHandler(userHandler *api.UserHandler) *BaseHandler {
	return &BaseHandler{userHandler: userHandler}
}

func (h *BaseHandler) UserRegister(ctx echo.Context) error {
	return h.userHandler.UserRegister(ctx)
}
