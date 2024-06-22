package main

import "time"

type Config struct {
	Operation OperationConfig `koanf:"operation"`
}

func (Config) Default() Config {
	return Config{
		Operation: OperationConfig{}.Default(),
	}
}

type OperationConfig struct {
	LogLevel        string        `koanf:"log_level"`
	ShutdownTimeout time.Duration `koanf:"shutdown_timeout"`
}

func (OperationConfig) Default() OperationConfig {
	return OperationConfig{
		LogLevel:        "info",
		ShutdownTimeout: 10 * time.Second,
	}
}
