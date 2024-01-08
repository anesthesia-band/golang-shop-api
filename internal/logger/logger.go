package logger

import (
	"log/slog"
	"os"

	"github.com/anesthesia-band/golang-shop-api/internal/constants"
)

type serverLogger interface {
	Info(s string, args ...any)
	Debug(s string, args ...any)
	Warn(s string, args ...any)
	Error(s string, args ...any)
}

func SetUp(env string) serverLogger {
	var log *slog.Logger

	switch env {
	case constants.Local:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case constants.Prod:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
