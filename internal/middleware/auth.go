package middleware

import (
	"net/http"
	"strings"

	"github.com/abdallahelassal/Store/internal/modules/user/domain"
	"github.com/abdallahelassal/Store/pkg"
	"github.com/abdallahelassal/Store/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	service *pkg.JWTService
}

type AuthenticatedUser struct {
    UUID  string
    Email string
    Role  domain.Role
}

func NewAuthMiddleware(jwtService *pkg.JWTService)*AuthMiddleware{
	return &AuthMiddleware{service: jwtService}
}

func (m *AuthMiddleware) RequireAuth()gin.HandlerFunc{
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized,"authorization header required","")
			c.Abort()
			return 
		}

		parts := strings.Fields(authHeader)
		if len(parts) != 2 || parts[0] != "Bearer"{
			utils.ErrorResponse(c,http.StatusUnauthorized,"invalid authorization format", "")
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims , err := m.service.ValidateToken(tokenString)
		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized,"invalid or expired token",err.Error())
			c.Abort()
			return 
		}
		authUser := AuthenticatedUser{
			UUID: claims.UserID,
			Email: claims.Email,
			Role: claims.Role,
		}

		c.Set("auth_user",authUser)

		c.Next()
	}
}
func (m *AuthMiddleware) RequireRole(roles ...string)gin.HandlerFunc{
	return func(c *gin.Context) {
		userInterface , exists := c.Get("auth_user")
		if !exists{
			utils.ErrorResponse(c,http.StatusForbidden,"role not found","")
			c.Abort()
			return 
		}

		authUser := userInterface.(AuthenticatedUser)
		allowed := false
		for _ , role := range roles {
			if string(authUser.Role) == role {
				allowed = true
				break
			}
		}
		if !allowed {
			utils.ErrorResponse(c,http.StatusForbidden,"insufficient permissions","")
			c.Abort()
			return 
		}
		c.Next()
	}
}