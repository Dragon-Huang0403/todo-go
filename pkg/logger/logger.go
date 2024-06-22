package logger

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ctxKey struct{}

func Init(ctx context.Context, _logLevel string) (context.Context, error) {
	logLevel, err := zapcore.ParseLevel(_logLevel)
	if err != nil {
		return nil, fmt.Errorf("failed to parse log level: %w", err)
	}

	zapConfig := zap.NewProductionConfig()
	zapConfig.Level = zap.NewAtomicLevelAt(logLevel)
	logger, err := zapConfig.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build logger: %w", err)
	}

	ctx = context.WithValue(ctx, ctxKey{}, logger)
	return ctx, nil
}

func loggerFromContext(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value(ctxKey{}).(*zap.Logger)
	if !ok {
		fmt.Println("failed to get logger from context")
		return zap.NewNop()
	}
	return logger
}

func Debug(ctx context.Context, msg string, fields ...zapcore.Field) {
	loggerFromContext(ctx).Debug(msg, fields...)
}

func Info(ctx context.Context, msg string, fields ...zapcore.Field) {
	loggerFromContext(ctx).Info(msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...zapcore.Field) {
	loggerFromContext(ctx).Warn(msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...zapcore.Field) {
	loggerFromContext(ctx).Error(msg, fields...)
}

func Fatal(ctx context.Context, msg string, fields ...zapcore.Field) {
	loggerFromContext(ctx).Fatal(msg, fields...)
}

func Level(ctx context.Context) zapcore.Level {
	return loggerFromContext(ctx).Level()
}
