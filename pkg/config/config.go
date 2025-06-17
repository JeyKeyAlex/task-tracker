package config

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	WebDir string `env:"WEB_DIR,default=./web"`
	Port   string `env:"TODO_PORT,default=:7540"`
	DbFile string `env:"TODO_DBFILE,default=./pkg/db/scheduler.db"`
}

func NewConfig() *Config {
	// Загружаем переменные окружения из .env файла
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: .env file not found or couldn't be loaded")
	}

	cfg := &Config{}

	// Создаем контекст
	ctx := context.Background()

	// Парсим переменные окружения в структуру Config
	if err := envconfig.Process(ctx, cfg); err != nil {
		log.Printf("Failed to parse env config: %v", err)
	}

	return cfg
}
