package main

import (
	"learn/internal/app"
	"learn/internal/config"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load '.env' file: %v", err)
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configs: %v", err)
	}

	ap, err := app.New(cfg)
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}
	defer ap.DB.Close()

	log.Printf("starting server at port: %s", ap.Cfg.Port)
	if err := ap.Run(); err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}
