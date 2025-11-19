package service

import (
	"context"

	"github.com/abdallahelassal/Store/internal/modules/user/dtos/request"
	"github.com/abdallahelassal/Store/internal/modules/user/dtos/response"
)

type UserService interface {
	Register(ctx context.Context,req request.RegisterRequest)(*response.AuthResponse,error)
	Login(ctx context.Context,req request.LoginRequest)(*response.LoginResponse,error)
	GetProfile(ctx context.Context,UserID uint)(*response.UserResponse,error)
	UpdateProfile(ctx context.Context,UserID string,req request.UpdateProfileRequest)(*response.UserResponse,error)
	RefreshToken(ctx context.Context,refreshToken string)(*response.AuthResponse,error)
	ListUsers(ctx context.Context,limit , offset int)([]*response.UserResponse,error)
	GetByID(ctx context.Context,id uint)(*response.UserResponse, error)
	GetByUUID(ctx context.Context,uuid string)(*response.UserResponse,error)
	DeleteUser(ctx context.Context,uuid string)error
}