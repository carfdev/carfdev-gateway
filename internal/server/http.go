package server

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/carfdev/carfdev-gateway/internal/config"
	"github.com/carfdev/carfdev-gateway/internal/email"
)

type HTTPServer struct {
	engine *gin.Engine
	config *config.Config
}

func NewHTTPServer(cfg *config.Config) *HTTPServer {

	gin.SetMode(cfg.GinMode)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	if err := r.SetTrustedProxies(nil); err != nil {
		log.Printf("Warning: could not set trusted proxies: %v", err)
	}

	api := r.Group("/api")

	// Module Register
	email.RegisterRoutes(api.Group("/email"), cfg)

	return &HTTPServer{
		engine: r,
		config: cfg,
	}
}

func (s *HTTPServer) Start() error {
	return s.engine.Run(":" + s.config.Port)
}
