package server

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
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

	corsCfg := cors.Config{
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	if cfg.Env == "development" {
		corsCfg.AllowAllOrigins = true
	} else {
		corsCfg.AllowOrigins = []string{cfg.Client}
	}

	r.Use(cors.New(corsCfg))

	api := r.Group("/v1")

	// Module Register
	email.RegisterRoutes(api.Group("/email"), cfg)

	return &HTTPServer{
		engine: r,
		config: cfg,
	}
}

func (s *HTTPServer) Start() error {
	log.Printf("Starting HTTP server on port %s", s.config.Port)
	return s.engine.Run(":" + s.config.Port)
}
