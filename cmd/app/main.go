package main

import (
	"log"
	"ssr/config"
	"ssr/internal/app"
	docs "ssr/swagger"
)

// @description Student-Supervisor Relationships

// @contact.name Ivakhnenko Dmitry, Moscow
// @contact.url github.com/ivaahn/ssr
// @contact.email ivahnencko01@gmail.com

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl /api/v1/auth/login
func main() {
	cfg, err := config.NewConfig()

	docs.SwaggerInfo.Version = cfg.App.Version
	docs.SwaggerInfo.Title = cfg.App.Name
	docs.SwaggerInfo.Host = cfg.HTTP.Host
	docs.SwaggerInfo.BasePath = cfg.HTTP.BasePath
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
