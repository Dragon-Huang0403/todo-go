package httpserver

import (
	"context"

	"github.com/dragon-huang0403/todo-go/internal/controller"
	"github.com/dragon-huang0403/todo-go/internal/http/server/handler"
	httpserver "github.com/dragon-huang0403/todo-go/pkg/http/server"
	"github.com/dragon-huang0403/todo-go/pkg/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//	@title			Todo Server API
//	@version		1.0.0
//	@description	Todo Server API
//	@schemes		http
//	@host			localhost:8080
//	@BasePath		/

func NewServer(ctx context.Context, ctl *controller.Controller, validator *validator.Validator) *echo.Echo {
	e := echo.New()
	e.Validator = validator

	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.Recover(), httpserver.LogMiddleware(ctx))

	handler := handler.New(ctl)

	// Add routes
	router := e.Group("")
	addRoutes(router, handler)

	return e
}
