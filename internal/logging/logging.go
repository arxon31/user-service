package logging

import (
	"github.com/arxon31/user-service/internal/config"
	"log/slog"
	"os"
)

func GetLogger(cfg *config.Config) *slog.Logger {
	if *cfg.IsDebug {
		logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug}))
		return logger
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo}))
	return logger
}
