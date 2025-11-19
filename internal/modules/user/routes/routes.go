package routes

import (
	"github.com/abdallahelassal/Store/internal/middleware"
	"github.com/abdallahelassal/Store/internal/modules/user/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup,h *handler.UserHandler,authMW *middleware.AuthMiddleware) {
	auth := rg.Group("auth")
	{
		auth.POST("/register",h.Register)
		auth.POST("/login",h.Login)
		auth.POST("/refresh",h.RefreshToken)
	}
	rg.PUT("/porfile",authMW.RequireAuth(),h.UpdateProfile)
	

	users := rg.Group("users")
	users.Use(authMW.RequireAuth(),authMW.RequireRole("admin"))
	{
		users.GET("",h.ListUsers)
		users.GET("/:id",h.GetProfile)
		users.DELETE("/:id",h.DeleteUser)
	}
}