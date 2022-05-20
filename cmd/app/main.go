package main

import (
	"log"
	"ssr/config"
	"ssr/internal/app"
)

// @title SSR Application
// @version 0.1.0
// @description Student-Supervisor Relationships

// @contact.name Ivakhnenko Dmitry, Moscow
// @contact.url github.com/ivaahn/ssr
// @contact.email ivahnencko01@gmail.com

// @host localhost:8080
// @BasePath /
// @schemes http

// @securitydefinitions.apikey Auth
// @in header
// @name Authorization
func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
