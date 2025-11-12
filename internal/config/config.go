package config

import (
	"os"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Port    string
	GinMode string
	Env     string
	Domain  string
	NatsUrl string
	Client  string
}

func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = gin.DebugMode
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	domain := os.Getenv("DOMAIN")
	if domain == "" {
		domain = "localhost:8080"
	}

	url := os.Getenv("NATS_URL")

	if url == "" {
		url = "nats://localhost:4222"
	}

	client := os.Getenv("CLIENT_URL")
	if client == "" {
		client = "http://localhost:8080"
	}

	return &Config{
		Port:    port,
		GinMode: mode,
		Env:     env,
		Domain:  domain,
		NatsUrl: url,
		Client:  client,
	}
}
