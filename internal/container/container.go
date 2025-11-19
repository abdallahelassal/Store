package container

import (
	"github.com/abdallahelassal/Store/internal/middleware"
	"github.com/abdallahelassal/Store/internal/modules/user/handler"
	"github.com/abdallahelassal/Store/internal/modules/user/repository"
	"github.com/abdallahelassal/Store/internal/modules/user/service"
	"github.com/abdallahelassal/Store/pkg"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Container struct {
	UserHandler *handler.UserHandler
	AuthMiddleware *middleware.AuthMiddleware
	Logger			*zap.Logger
}

func NewContainer(db *gorm.DB,jwtService *pkg.JWTService,logger *zap.Logger)*Container{
	userRepo := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepo,jwtService)

	userHandler := handler.NewUserHandler(userService)

	authMiddleware := middleware.NewAuthMiddleware(jwtService)
	return &Container{
		UserHandler: userHandler,
		AuthMiddleware: authMiddleware,
		Logger: logger,
	}
}