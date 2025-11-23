package main

import (
	"log"

	"github.com/abdallahelassal/Store/config"
	"github.com/abdallahelassal/Store/database"
	"github.com/abdallahelassal/Store/internal/container"
	"github.com/abdallahelassal/Store/internal/server"

	"github.com/abdallahelassal/Store/internal/router"
	"github.com/abdallahelassal/Store/pkg/logger"
	"github.com/abdallahelassal/Store/pkg"
	
)

func main() {
	cfg := config.LoadConfig()

	appLogger := logger.NewLogger(cfg.Environment)


	db  := database.NewConnection(cfg, appLogger)

	jwtService := pkg.NewJWTService(cfg.JWTConfig,cfg.JWTConfig.AccessExpiration,cfg.JWTConfig.RefreshExpiration)

	cont := container.NewContainer(db.DB,jwtService,appLogger)

	r := router.SetupRouter(cont,cfg)

	srv := server.NewServer(cfg.ServerConfig.PORT,r,appLogger)

	if err := srv.Start();err != nil {
		log.Fatal("faild start server",err)
	}
	sqlDB , _ := db.DB.DB()
	sqlDB.Close()
}
