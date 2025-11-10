package email

import (
	"github.com/carfdev/carfdev-gateway/internal/config"
	"github.com/carfdev/carfdev-gateway/internal/helper"
	"github.com/carfdev/carfdev-gateway/internal/nats"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, cfg *config.Config) {

	url := cfg.NatsUrl

	nc, err := nats.NewNatsClient(url)
	if err != nil {
		panic("failed to connect to NATS: " + err.Error())
	}
	service := NewEmailService(nc)
	response := helper.NewResponse()
	controller := NewEmailController(service, response)

	rg.POST("/send-contact", controller.SendContact)
}
