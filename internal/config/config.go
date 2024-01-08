package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-required:"true"`
	Storage    string `yaml:"storage_path" env-required:"true"`
	HttpServer `yaml:"http_server"`
}

type HttpServer struct {
	Address string        `yaml:"address" env-default:"localhost:8080"`
	Timeout time.Duration `yaml:"idle_timeout" env-default:"4s"`
}

func MustLoad() Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "./config/config.yaml"
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Cannot read config: %s", err)
	}

	return cfg
}
