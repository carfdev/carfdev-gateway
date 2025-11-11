package main

import (
	"log"

	"github.com/carfdev/carfdev-gateway/internal/config"
	"github.com/carfdev/carfdev-gateway/internal/server"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	cfg := config.LoadConfig()

	s := server.NewHTTPServer(cfg)

	if err := s.Start(); err != nil {
		log.Fatalf("Server stopped with error: %v", err)
	}
}
