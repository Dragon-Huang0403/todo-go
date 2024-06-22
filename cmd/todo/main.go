package main

import (
	"flag"
	"log"

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
	appConfig, err := getAppConfig(configFile)
	if err != nil {
		log.Fatalf("failed to get config: %v", err)
	}

	validator := validator.New()
	err = validator.Validate(appConfig)
	if err != nil {
		log.Fatalf("failed to validate config: %v", err)
	}

	logger, err := logger.New(appConfig.Operation.LogLevel)
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}

	logger.Info("config", zap.Any("config", appConfig))
}
