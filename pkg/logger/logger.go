package logger

import (
	"log/slog"
	"os"
	"sync"
)

type Logger struct {
	slog *slog.Logger
	mu   sync.Mutex
}

var AppLogger *Logger

func init() {
	file, err := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		slog.Error(err.Error())
	}
	handler := slog.NewTextHandler(file, &slog.HandlerOptions{
		Level:     slog.LevelError,
		AddSource: true,
	})
	AppLogger = &Logger{slog: slog.New(handler)}
}

func (l *Logger) ErrorLogger(errMsg string, args ...any) {
	l.slog.Error(errMsg, args...)
}
