package router

import (
	"net/http"

	"github.com/abdallahelassal/Store/config"
	"github.com/abdallahelassal/Store/internal/container"
	"github.com/abdallahelassal/Store/internal/middleware"
	userRoutes "github.com/abdallahelassal/Store/internal/modules/user/routes"
	"github.com/gin-gonic/gin"
)

func SetupRouter(cont *container.Container, cfg *config.Config) *gin.Engine {
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()

	r.Use(middleware.Recovery())
	r.Use(middleware.Logger(cont.Logger))
	r.Use(middleware.Cors())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "server is runing",
		})
	})
	// Register module routes. Passing a root group so routes like
	// `/auth/register` are available at the top-level.
	userRoutes.RegisterRoutes(r.Group("/"), cont.UserHandler, cont.AuthMiddleware)
	return r
}
