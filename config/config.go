package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type (
	// Config -.
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		Pg   `yaml:"postgres"`
		Auth `yaml:"auth"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// Pg -.
	Pg struct {
		DSN string `env-required:"true" yaml:"dsn" env:"PG_DSN"`
	}

	Auth struct {
		TokenExp   int    `env-required:"true" yaml:"token_exp" env:"AUTH_TOKEN_EXP"`
		SigningKey string `env-required:"true" yaml:"signing_key" env:"AUTH_SIGNING_KEY"`
	}
)

const (
	testConfigPath = "./config/config_test.yml"
	prodConfigPath = "/config.yml"
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	var configPath string

	switch os.Getenv("DEPLOY_MODE") {
	case "PROD":
		configPath = prodConfigPath
	default:
		configPath = testConfigPath
	}

	cfg := &Config{}
	err := cleanenv.ReadConfig(configPath, cfg)

	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
