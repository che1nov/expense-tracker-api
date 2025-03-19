package config

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func InitLogger() {
	handler := slog.NewTextHandler(os.Stdout, nil)
	Logger = slog.New(handler)
}
