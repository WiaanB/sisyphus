package logging

import (
	"log/slog"

	"github.com/WiaanB/sisyphus/internal/logging/service"
)

var Logger *Logging

type Logging struct {
	Service *slog.Logger
}

func New() {
	svc := service.NewLoggingService()

	Logger = &Logging{
		Service: svc,
	}
}
