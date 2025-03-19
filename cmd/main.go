package main

import (
	"expense-tracker-api/internal/config"
	database "expense-tracker-api/internal/db"
	"expense-tracker-api/internal/routes"
	"log/slog"
	"net/http"
)

func main() {
	config.InitLogger()
	logger := config.Logger

	cfg := config.LoadConfig()
	logger.Info("Loaded configuration")

	database.Connect(cfg.DBName)
	logger.Info("Connected to database")

	logger.Info("Starting server...", slog.String("port", cfg.Port))
	err := http.ListenAndServe(":"+cfg.Port, routes.SetupRoutes())
	if err != nil {
		logger.Error("Failed to start server", slog.String("error", err.Error()))
	}
}
