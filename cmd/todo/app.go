package main

import (
	"context"
	"time"

	"github.com/dragon-huang0403/todo-go/internal/controller"
	"github.com/dragon-huang0403/todo-go/internal/db"
	httpserver "github.com/dragon-huang0403/todo-go/internal/http/server"
	"github.com/dragon-huang0403/todo-go/internal/store"
	"github.com/dragon-huang0403/todo-go/pkg/logger"
	"github.com/dragon-huang0403/todo-go/pkg/validator"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func Start(ctx context.Context, config AppConfig, validator *validator.Validator) error {
	wg, ctx := errgroup.WithContext(ctx)

	db := db.New()
	store := store.New(db)
	controller := controller.New(store)

	wg.Go(func() error {
		return httpserver.Start(ctx, config.HTTPServer, controller, validator)
	})

	<-ctx.Done()
	logger.Info(ctx, "shutting down application")

	// Graceful shutdown
	stopCh := make(chan struct{}, 1)
	go func() {
		if err := wg.Wait(); err != nil {
			logger.Error(ctx, "Failed to shutdown application", zap.Error(err))
		}
		close(stopCh)
	}()

	select {
	case <-stopCh:
		logger.Info(ctx, "Application gracefully down")
	case <-time.After(config.Operation.ShutdownTimeout):
		logger.Warn(ctx, "Application shutdown timeout")
	}

	return nil
}
