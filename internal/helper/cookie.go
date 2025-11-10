package helper

import (
	"time"

	"github.com/carfdev/carfdev-gateway/internal/config"
	"github.com/gin-gonic/gin"
)

type Cookie interface {
	SetCookie(ctx *gin.Context, name, value string)
	RemoveCookie(ctx *gin.Context, name string)
}

type cookie struct {
	domain string
	secure bool
	expiry time.Duration
}

func NewCookie(cfg *config.Config) Cookie {
	secure := cfg.Env == "production"
	expiry := 30 * 24 * time.Hour
	return &cookie{
		domain: cfg.Domain,
		secure: secure,
		expiry: expiry,
	}
}

func (c *cookie) SetCookie(ctx *gin.Context, name, value string) {
	ctx.SetCookie(
		name,
		value,
		int(c.expiry.Seconds()),
		"/",
		c.domain,
		c.secure,
		true,
	)
}

func (c *cookie) RemoveCookie(ctx *gin.Context, name string) {
	ctx.SetCookie(
		name,
		"",
		-1,
		"/",
		c.domain,
		c.secure,
		true,
	)
}
