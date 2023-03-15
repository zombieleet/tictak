package logger

import (
	"context"
	"golang.org/x/exp/slog"
	"io"
	"os"
	"reflect"
)

type LoggerType string

const (
	JSON LoggerType = "json"
	TEXT LoggerType = "text"
)

type LoggerOptions struct {
	LogLevel slog.Level
	LogType  LoggerType
}

type Logger struct {
	slogger *slog.Logger
	level   slog.Level
}

func NewLogger(loggerOptions LoggerOptions) *Logger {

	loggerOpts := slog.HandlerOptions{
		Level: loggerOptions.LogLevel,
	}

	if reflect.ValueOf(loggerOptions.LogType).IsZero() {
		loggerOptions.LogType = TEXT
	}

	var slogger *slog.Logger

	handlerMaps := map[LoggerType]interface{}{
		JSON: loggerOpts.NewJSONHandler,
		TEXT: loggerOpts.NewTextHandler,
	}[loggerOptions.LogType]

	if handler, ok := handlerMaps.(func(io.Writer) *slog.JSONHandler); ok {
		slogger = slog.New(handler(os.Stdout))
	}

	if handler, ok := handlerMaps.(func(io.Writer) *slog.TextHandler); ok {
		slogger = slog.New(handler(os.Stdout))
	}

	return &Logger{
		slogger: slogger,
		level:   loggerOptions.LogLevel,
	}
}

func (logger *Logger) Log(msg string) {
	logger.slogger.Info(msg)
}

func (logger *Logger) LogWithCtx(ctx context.Context, msg string, group string, kv []any) {
	logger.slogger.WithGroup(group).InfoCtx(ctx, msg, kv...)
}

func (logger *Logger) NetworkError(err error) {
	logger.slogger.WithGroup("NETWORK_ERROR").Error("Error:", err)
}
