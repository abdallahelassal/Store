package routes

import (
	"github.com/abdallahelassal/Store/internal/middleware"
	"github.com/abdallahelassal/Store/internal/modules/branch/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, h *handler.BranchHandler, authMW *middleware.AuthMiddleware) {
	branches := rg.Group("branches")
	branches.Use(authMW.RequireAuth())
	{
		branches.POST("", authMW.RequireRole("admin"), h.CreateBranch)
		branches.GET("", h.ListBranches)
		branches.GET("/:uuid", h.GetBranch)
		branches.DELETE("/:uuid", authMW.RequireRole("admin"), h.DeleteBranch)
	}
}
