package service

import (
	"log/slog"
	"os"
)

func NewLoggingService() *slog.Logger {
	jh := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	})

	logger := slog.New(jh)
	return logger
}
