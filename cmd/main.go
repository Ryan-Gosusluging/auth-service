package main

import (
	"log"

	"github.com/Ryan-Gosusluging/auth-service/config"
	"github.com/Ryan-Gosusluging/auth-service/internal/app"
)

// @title Auth Service API
// @version 1.0
// @description API for auth service
// @host localhost:3001
// @BasePath /
func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}