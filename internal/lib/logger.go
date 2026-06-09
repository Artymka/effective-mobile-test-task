package lib

import (
	"log/slog"
	"os"
)

type Logger struct {
	logger *slog.Logger
}

func NewLogger() *Logger {
	return &Logger{slog.New(slog.NewTextHandler(os.Stdout, nil))}
}

func (l *Logger) Error(op string, err error) {
	l.logger.Error("", slog.Attr{"op", slog.AnyValue(op)}, slog.Attr{"err", slog.AnyValue(err)})
}
func (l *Logger) ErrorMsg(op string, msg string) {
	l.logger.Error(msg, slog.Attr{"op", slog.AnyValue(op)})
}

func (l *Logger) Info(op string, msg string) {
	l.logger.Info(msg, slog.Attr{"op", slog.AnyValue(op)})
}

func (l *Logger) Debug(op string, err error, msg string) {
	l.logger.Debug(msg, slog.Attr{"op", slog.AnyValue(op)}, slog.Attr{"err", slog.AnyValue(err)})
}
