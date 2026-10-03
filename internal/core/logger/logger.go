package core_logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

type Logger struct {
	*slog.Logger
	file *os.File
}

type loggerContext struct{}

var loggerContextKey = loggerContext{}
var defaultLogger = Logger{Logger: slog.Default()}

// создает логгер который пишет в stdout и logfile
func NewLogger(config *Config) (*Logger, error) {
	if err := os.MkdirAll(config.LogFolder, 0o755); err != nil {
		return nil, fmt.Errorf("failed to create log folder: %w", err)
	}

	timestamp := time.Now().UTC().Format("2006-01-02T15-04-05.000000")
	logFilePath := filepath.Join(config.LogFolder, timestamp+".log")

	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, fmt.Errorf("failed to create log file: %w", err)
	}

	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.Level(config.LogLevel),
	}

	fileHandler := slog.NewJSONHandler(logFile, opts)
	stdoutHandler := slog.NewTextHandler(os.Stdout, opts)
	multiHandler := slog.NewMultiHandler(fileHandler, stdoutHandler)
	logger := slog.New(multiHandler)

	return &Logger{
		Logger: logger,
		file:   logFile,
	}, nil
}

// Закрывает файл в который пишет логгер, вызывать дефером обязательно
func (l *Logger) Close() error {
	if l.file != nil {
		l.file.Close()
	}

	return nil
}

// Добавляет поля в логгер
func (l *Logger) With(fields ...any) *Logger {
	return &Logger{
		Logger: l.Logger.With(fields...),
		file:   l.file,
	}
}

// Добавляет логгер в контекст
func WithLogger(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, loggerContextKey, logger)
}

// Достает логгер из контекста
func FromContext(ctx context.Context) *Logger {
	logger, ok := ctx.Value(loggerContextKey).(*Logger)
	if !ok {
		return &defaultLogger
	}

	return logger
}
