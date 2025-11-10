package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/carfdev/carfdev-gateway/internal/nats"
	"github.com/gin-gonic/gin"
)

const SubjectCheckAccess = "users.check_access"

type CheckAccessRequest struct {
	Token string `json:"access_token"`
}

type CheckAccessResponse struct {
	ID   string `json:"id"`
	Role string `json:"role"`
}

func CheckAccess(nc *nats.NatsClient, token string, timeout time.Duration) (*CheckAccessResponse, error) {
	req := CheckAccessRequest{Token: token}
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resp, err := nc.RequestWithContext(ctx, SubjectCheckAccess, data)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, errors.New("access check timed out")
		}
		return nil, errors.New("access denied")
	}

	var checkRes CheckAccessResponse
	if err := json.Unmarshal(resp, &checkRes); err != nil {
		return nil, err
	}

	return &checkRes, nil
}

func AuthMiddleware(nc *nats.NatsClient, requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid Authorization header"})
			c.Abort()
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")

		checkRes, err := CheckAccess(nc, token, 3*time.Second)
		if err != nil {
			status := http.StatusUnauthorized
			if err.Error() == "access check timed out" {
				status = http.StatusGatewayTimeout
			}
			c.JSON(status, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if len(requiredRoles) > 0 && !slices.Contains(requiredRoles, strings.ToLower(checkRes.Role)) {
			c.JSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
			c.Abort()
			return
		}

		c.Set("userID", checkRes.ID)
		c.Set("role", checkRes.Role)

		c.Next()
	}
}
