package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env            string        `yaml:"env"             env-required:"true"`
	ContextTimeout time.Duration `yaml:"context_timeout" env-required:"true"`
	DatabasePath   string        `yaml:"database_path"   env-required:"true"`
	Host           string        `yaml:"host"            env-required:"true"`
	Port           string        `yaml:"port"            env-required:"true"`
}

func MustLoad() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("ERROR: CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("ERROR: %v", err)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	return &cfg
}
