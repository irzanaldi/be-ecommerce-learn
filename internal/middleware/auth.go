package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/irzan-aldi/be-ecommerce/internal/config"
	"github.com/irzan-aldi/be-ecommerce/pkg/response"
	"net/http"
	"strings"
)

func JWTAuth(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Ambil header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, http.StatusUnauthorized, "missing token", nil)
			c.Abort()
			return
		}

		// 2. Extract "Bearer xxx"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Error(c, http.StatusUnauthorized, "invalid token format", nil)
			c.Abort()
			return
		}

		// 3. Parse & validate
		token, err := jwt.Parse(parts[1], func(t *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTSecret), nil
		})
		if err != nil || !token.Valid {
			response.Error(c, http.StatusUnauthorized, "invalid or expired token", nil)
			c.Abort()
			return
		}

		// 4. Extract claims → set ke context
		claims := token.Claims.(jwt.MapClaims)
		c.Set("userID", uint(claims["user_id"].(float64)))
		c.Set("email", claims["email"].(string))
		c.Set("role", claims["role"].(string))

		c.Next()
	}
}
