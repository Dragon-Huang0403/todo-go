package httpserver

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/dragon-huang0403/todo-go/internal/controller"
	"github.com/dragon-huang0403/todo-go/pkg/logger"
	"github.com/dragon-huang0403/todo-go/pkg/validator"
	"go.uber.org/zap"
)

// Start will block until the server is shutdown
// And will start graceful shutdown when the context is done
func Start(ctx context.Context, config Config, ctl *controller.Controller, validator *validator.Validator) error {
	server := NewServer(ctx, ctl, validator)

	// start server
	go func() {
		logger.Info(ctx, "server is starting", zap.String("addr_port", config.AddrPort))
		err := server.Start(config.AddrPort)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error(ctx, "Failed to start http server", zap.Error(err))
		}
	}()

	// wait for context done
	<-ctx.Done()

	// graceful shutdown server
	stopCh := make(chan struct{}, 1)

	go func() {
		logger.Info(ctx, "Shutting down server")
		err := server.Shutdown(ctx)
		if err != nil {
			logger.Error(ctx, "Failed to shutdown http server", zap.Error(err))
		}
		close(stopCh)
	}()

	select {
	case <-stopCh:
		logger.Info(ctx, "Server gracefully shutdown")
	case <-time.After(config.ShutdownTimeout):
		logger.Warn(ctx, "Server force shutdown")
		if err := server.Close(); err != nil {
			logger.Error(ctx, "Failed to force shutdown http server", zap.Error(err))
		}
	}

	return nil
}
