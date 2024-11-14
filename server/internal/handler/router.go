package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/yamato0204/cache-server-golang/internal/handler/middleware"
	"github.com/yamato0204/cache-server-golang/internal/handler/schema"
)

type Router struct {
	baseHandler           *BaseHandler
	transactionMiddleware *middleware.TransactionMiddleware
}

func NewRouter(baseHandler *BaseHandler, transactionMiddleware *middleware.TransactionMiddleware) *Router {
	return &Router{baseHandler: baseHandler, transactionMiddleware: transactionMiddleware}
}

func (r *Router) RegisterRoutes(e *echo.Echo) {
	e.Use(r.transactionMiddleware.Intercept)
	schema.RegisterHandlers(e, r.baseHandler)
}
