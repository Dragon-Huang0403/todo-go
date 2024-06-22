package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dragon-huang0403/todo-go/pkg/logger"
	"github.com/dragon-huang0403/todo-go/pkg/validator"
	"go.uber.org/zap"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "", "config file path")
	flag.Parse()
}

func main() {
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	appConfig, err := getAppConfig(configFile)
	if err != nil {
		log.Fatalf("failed to get config: %v", err)
	}

	validator := validator.New()
	err = validator.Validate(appConfig)
	if err != nil {
		log.Fatalf("failed to validate config: %v", err)
	}

	ctx, err = logger.Init(ctx, appConfig.Operation.LogLevel)
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}

	logger.Info(ctx, "config loaded", zap.Any("config", appConfig))
	logger.Info(ctx, "starting todo app")

	// Start the todo app
	err = Start(ctx, *appConfig, validator)
	if err != nil {
		logger.Error(ctx, "failed to start todo app", zap.Error(err))
	}

	logger.Info(ctx, "todo app stopped")
}
