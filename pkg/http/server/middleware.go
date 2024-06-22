package httpserver

import (
	"context"

	"github.com/dragon-huang0403/todo-go/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

const ctxKey = "context-key"

func TransformContext(c echo.Context) context.Context {
	ctx, ok := c.Get(ctxKey).(context.Context)
	if !ok {
		return context.Background()
	}

	return ctx
}

func LogMiddleware(pctx context.Context) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogRemoteIP: true,
		LogMethod:   true,
		LogURI:      true,
		LogStatus:   true,
		LogHost:     true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			zapLogger := logger.LoggerFromContext(pctx)
			ctx := c.Request().Context()
			ctx = logger.ContextWithLogger(ctx, zapLogger)
			c.Set(ctxKey, ctx)

			logger.Info(
				ctx,
				"Request",
				zap.Int("status", v.Status),
				zap.String("method", v.Method),
				zap.String("uri", v.URI),
				zap.String("remote_ip", v.RemoteIP),
				zap.String("host", v.Host))

			return nil
		},
	})
}
