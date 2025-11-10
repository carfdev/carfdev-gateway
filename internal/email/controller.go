package email

import (
	"net/http"

	"github.com/carfdev/carfdev-gateway/internal/helper"
	"github.com/gin-gonic/gin"
)

type EmailController struct {
	service EmailService
	res     helper.Response
}

func NewEmailController(service EmailService, res helper.Response) *EmailController {
	return &EmailController{service: service, res: res}
}

func (c *EmailController) SendContact(ctx *gin.Context) {

	var req SendContactRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.res.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request body: "+err.Error())
		return
	}

	resp, err := c.service.SendContact(ctx.Request.Context(), req)
	if err != nil {
		c.res.ErrorResponse(ctx, http.StatusBadGateway, "Unable to send contact email at this time.")
		return
	}

	c.res.SuccessResponse(ctx, resp.Status, resp.Message)
}
