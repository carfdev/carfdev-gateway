package helper

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Response interface {
	ErrorResponse(ctx *gin.Context, status int, message string)
	SuccessResponse(ctx *gin.Context, status int, data any)
}

type response struct{}

func NewResponse() Response {
	return &response{}
}

func (r *response) ErrorResponse(ctx *gin.Context, status int, message string) {
	ctx.JSON(status, gin.H{
		"status":    status,
		"error":     message,
		"timestamp": time.Now().Unix(),
	})
}

func (r *response) SuccessResponse(ctx *gin.Context, status int, data any) {
	ctx.JSON(status, gin.H{
		"status":    status,
		"data":      data,
		"timestamp": time.Now().Unix(),
	})
}
