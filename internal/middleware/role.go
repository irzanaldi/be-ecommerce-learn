package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/irzan-aldi/be-ecommerce/pkg/response"
	"net/http"
)

func RequireRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.GetString("role") // dari JWT middleware

		for _, r := range roles {
			if r == userRole {
				c.Next()
				return
			}
		}

		response.Error(c, http.StatusForbidden, "access denied", nil)
		c.Abort()
	}
}
