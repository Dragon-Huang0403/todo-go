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
	if logLevel == zapcore.DebugLevel {
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	zapConfig.Level = zap.NewAtomicLevelAt(logLevel)

	logger, err := zapConfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, fmt.Errorf("failed to build logger: %w", err)
	}

	ctx = ContextWithLogger(ctx, logger)
	return ctx, nil
}

func Flush(ctx context.Context) error {
	logger := LoggerFromContext(ctx)
	return logger.Sync()
}

func ContextWithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, ctxKey{}, logger)
}

func LoggerFromContext(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value(ctxKey{}).(*zap.Logger)
	if !ok {
		fmt.Println("failed to get logger from context")
		return zap.NewNop()
	}
	return logger
}

func Debug(ctx context.Context, msg string, fields ...zapcore.Field) {
	LoggerFromContext(ctx).Debug(msg, fields...)
}

func Info(ctx context.Context, msg string, fields ...zapcore.Field) {
	LoggerFromContext(ctx).Info(msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...zapcore.Field) {
	LoggerFromContext(ctx).Warn(msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...zapcore.Field) {
	LoggerFromContext(ctx).Error(msg, fields...)
}

func Fatal(ctx context.Context, msg string, fields ...zapcore.Field) {
	LoggerFromContext(ctx).Fatal(msg, fields...)
}

func Level(ctx context.Context) zapcore.Level {
	return LoggerFromContext(ctx).Level()
}
