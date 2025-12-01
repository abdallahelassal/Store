package container

import (
	"github.com/abdallahelassal/Store/internal/middleware"
	"github.com/abdallahelassal/Store/internal/modules/user/handler"
	BranchHandler "github.com/abdallahelassal/Store/internal/modules/branch/handler"
	BranchRepository "github.com/abdallahelassal/Store/internal/modules/branch/repository"
	BranchService "github.com/abdallahelassal/Store/internal/modules/branch/service"
	"github.com/abdallahelassal/Store/internal/modules/user/repository"
	"github.com/abdallahelassal/Store/internal/modules/user/service"
	"github.com/abdallahelassal/Store/pkg"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Container struct {
	UserHandler *handler.UserHandler
	BranchHandler *BranchHandler.BranchHandler
	AuthMiddleware *middleware.AuthMiddleware
	Logger			*zap.Logger
}

func NewContainer(db *gorm.DB,jwtService *pkg.JWTService,logger *zap.Logger)*Container{
	userRepo := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepo,jwtService)

	userHandler := handler.NewUserHandler(userService)

	// Branch Module
	branchRepo := BranchRepository.NewBranchRepository(db)
	branchService := BranchService.NewBranchService(branchRepo)
	branchHandler := BranchHandler.NewBranchHandler(branchService)

	authMiddleware := middleware.NewAuthMiddleware(jwtService)
	return &Container{
		UserHandler: userHandler,
		BranchHandler: branchHandler,
		AuthMiddleware: authMiddleware,
		Logger: logger,
	}
}