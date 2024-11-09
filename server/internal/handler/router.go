package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/yamato0204/cache-server-golang/internal/handler/schema"
)

type Router struct {
	baseHandler *BaseHandler
}

func NewRouter(baseHandler *BaseHandler) *Router {
	return &Router{baseHandler: baseHandler}
}

func (r *Router) RegisterRoutes(e *echo.Echo) {
	schema.RegisterHandlers(e, r.baseHandler)
}
