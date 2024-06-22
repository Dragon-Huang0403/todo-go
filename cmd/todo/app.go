package main

import (
	"context"

	httpserver "github.com/dragon-huang0403/todo-go/internal/http/server"
	"github.com/dragon-huang0403/todo-go/pkg/validator"
)

func Start(ctx context.Context, config AppConfig, validator *validator.Validator) error {
	err := httpserver.Start(ctx, config.HTTPServer, validator)
	return err
}
