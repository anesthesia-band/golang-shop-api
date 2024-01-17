package main

import (
	"fmt"
	"os"

	"github.com/anesthesia-band/golang-shop-api/internal/config"
	"github.com/anesthesia-band/golang-shop-api/internal/logger"
	"github.com/anesthesia-band/golang-shop-api/internal/storage"
	"github.com/anesthesia-band/golang-shop-api/internal/storage/models/goods"
)

func main() {
	cfg := config.MustLoad()

	logger := logger.SetUp(cfg.Env)

	storage, err := storage.Init(cfg.Storage)
	if err != nil {
		logger.Error("unable to init storage", err)
		os.Exit(1)
	}

	result, err := goods.GetAll(storage, true)
	if err != nil {
		logger.Error("unable to insert good", err)
		os.Exit(1)
	}
	fmt.Println(result)

	logger.Info("Server started")
	logger.Debug("Debug mode enabled")
}
