package config

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	DSN  string
	Port string
}

func Load() (*Config, error) {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		return nil, fmt.Errorf("failed to load 'DSN'")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Printf("failed to load 'PORT', falling back")
		port = "8080"
	}

	return &Config{
		DSN:  dsn,
		Port: port,
	}, nil
}
