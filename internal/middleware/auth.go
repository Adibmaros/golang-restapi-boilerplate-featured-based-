package middleware

import (
	"golang-restapi-big-structure/internal/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtService jwt.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "token tidak tersedia",
			})
			return
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "token tidak valid",
			})
			return
		}

		claims, err := jwtService.ValidateToken(parts[1])

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "token tidak valid",
			})
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)
	}
}

func RequiredRoles(roles ...jwt.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, exist := c.Get("role")

		if !exist {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "kamu tidak punya akses",
			})
			return
		}

		userRole, ok := value.(jwt.Role)

		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "ada yang salah",
			})
			return
		}

		for _, role := range roles {
			if userRole == role {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "kamu tidak punya akses",
		})
	}
}
