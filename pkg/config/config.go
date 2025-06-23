package config

import (
	"context"
	"errors"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	WebDir string `env:"WEB_DIR,default=./web"`
	Port   string `env:"TODO_PORT,default=:7540"`
	DbFile string `env:"TODO_DBFILE,default=./pkg/db/scheduler.db"`
}

func NewConfig() (*Config, error) {
	// Загружаем переменные окружения из .env файла
	if err := godotenv.Load(".env"); err != nil {
		err = errors.New("warning: .env file not found or couldn't be loaded")
		return nil, err
	}

	cfg := &Config{}

	// Создаем контекст
	ctx := context.Background()

	// Парсим переменные окружения в структуру Config
	if err := envconfig.Process(ctx, cfg); err != nil {
		err = fmt.Errorf("failed to parse env config: %v", err)
		return nil, err
	}

	return cfg, nil
}
