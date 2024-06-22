package main

import (
	"fmt"
	"time"

	httpserver "github.com/dragon-huang0403/todo-go/internal/http/server"
	"github.com/dragon-huang0403/todo-go/pkg/config"
)

type AppConfig struct {
	HTTPServer httpserver.Config `koanf:"http_server" validate:"required"`
	Operation  OperationConfig   `koanf:"operation" validate:"required"`
}

func (AppConfig) Default() AppConfig {
	return AppConfig{
		HTTPServer: httpserver.Config{}.Default(),
		Operation:  OperationConfig{}.Default(),
	}
}

func getAppConfig(configFile string) (*AppConfig, error) {
	conf := config.Config{}.Default()
	conf.ConfigFile = configFile

	appConfig := AppConfig{}
	err := config.GetConfig(conf, &appConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to get config: %w", err)
	}

	return &appConfig, nil
}

type OperationConfig struct {
	LogLevel        string        `koanf:"log_level" validate:"required,oneof=debug info warn error"`
	ShutdownTimeout time.Duration `koanf:"shutdown_timeout" validate:"required"`
}

func (OperationConfig) Default() OperationConfig {
	return OperationConfig{
		LogLevel:        "info",
		ShutdownTimeout: 10 * time.Second,
	}
}
