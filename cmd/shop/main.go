package main

import (
	"os"

	"github.com/anesthesia-band/golang-shop-api/internal/config"
	"github.com/anesthesia-band/golang-shop-api/internal/logger"
	"github.com/anesthesia-band/golang-shop-api/internal/storage"
)

func main() {
	cfg := config.MustLoad()

	logger := logger.SetUp(cfg.Env)

	_, err := storage.Init(cfg.Storage)
	if err != nil {
		logger.Error("unable to init storage", err)
		os.Exit(1)
	}

	logger.Info("Server started")
	logger.Debug("Debug mode enabled")
}
