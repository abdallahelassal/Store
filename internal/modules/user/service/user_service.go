package service

import (
	"context"
	"log"

	"github.com/abdallahelassal/Store/internal/modules/user/domain"
	"github.com/abdallahelassal/Store/internal/modules/user/dtos/request"
	"github.com/abdallahelassal/Store/internal/modules/user/dtos/response"
	"github.com/abdallahelassal/Store/internal/modules/user/repository"
	"github.com/abdallahelassal/Store/pkg"
	"github.com/abdallahelassal/Store/pkg/utils"
)

type userService struct{
	repo		repository.UserRepository
	jwtService	*pkg.JWTService
}

func NewUserService(repo repository.UserRepository, jwtService *pkg.JWTService) UserService{
	return&userService{
		repo:	repo,
		jwtService: jwtService,
	}
}

func (s *userService) Register(ctx context.Context,req request.RegisterRequest)(*response.AuthResponse,error){
	existing , err := s.repo.GetByEmail(ctx,req.Email)
	if err == nil && existing != nil {
		return nil , domain.ErrEmailAlreadyExits
	}

	hashedPassword , err := utils.HashPassword(req.Password)
	if err != nil {
		log.Printf("error hashing password %v", err)
		return nil , err
	}

	user := &domain.User{
		User_name: req.Name,
		Email: req.Email,
		Password: hashedPassword,
		Role: domain.RoleUser,
	}
	if err := s.repo.Create(ctx,user); err != nil {
		log.Printf("error create user %v",err)
		return  nil , err
	}
	return s.generateAuthResponse(user)
}

func (s *userService) Login(ctx context.Context,req request.LoginRequest)(*response.LoginResponse,error){
	user , err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, domain.ErrInvalidCredintials
	}
	if !utils.ComparePassword(req.Password,user.Password){
		return nil , domain.ErrInvalidCredintials
	}

	accessToken, err := s.jwtService.GenerateToken(user.UUID,user.Email,string(user.Role))
	if err != nil {
		return nil, err 
	}
	refreshToken, err := s.jwtService.GenerateRefreshToken(user.UUID)
	if err != nil {
		return nil, err 
	}
	return &response.LoginResponse{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
		TokenType: "Bearer",
		ExpiresIn: 3600,
		User: s.toUserResponse(user),
	},nil 
}

func (c *userService) GetProfile(ctx context.Context,userID uint)(*response.UserResponse,error){
	user , err := c.repo.GetByID(ctx,userID)
	if err != nil {
		return nil , domain.ErrInvalidCredintials
	}
	return &response.UserResponse{
		Name: user.User_name,
		Email: user.Email,
		Role: string(user.Role),
	},nil 
}

func (c *userService) UpdateProfile(ctx context.Context, userID uint , req request.UpdateProfileRequest)(*response.UserResponse, error){
	user , err := c.repo.GetByID(ctx, userID)
	if err != nil {
		return nil , err 
	}

	if req.Name != ""{
		user.User_name = req.Name
	}

	if req.Email != ""{
		exiting , _ := c.repo.GetByEmail(ctx,req.Email)
		if exiting != nil && exiting.ID != user.ID {
			return nil , domain.ErrEmailAlreadyExits
		}
		user.Email = req.Email
	}
	if err := c.repo.Update(ctx,user); err != nil {
		return nil , err 
	}

	return c.toUserResponse(user),nil 
}

func (s *userService) RefreshToken(ctx context.Context,refreshToken string)(*response.AuthResponse, error){
	claims , err := s.jwtService.ValidateToken(refreshToken)
	if err != nil {
		return nil , domain.ErrInvalidToken
	}



	user , err := s.repo.GetByUUID(ctx, claims.UserID)
	if  err != nil {
		return nil , err 
	}
	return s.generateAuthResponse(user)

}

func (s *userService) ListUsers(ctx context.Context, limit , offset int)([]*response.UserResponse, error){
	user , err := s.repo.List(ctx,limit,offset)
	if err != nil {
		return nil ,err
	}

	response := make([]*response.UserResponse,len(user))
	
	for i, u := range user {
		response[i] = s.toUserResponse(u)
	}
	return response ,nil
}

func (s *userService) GetByID(ctx context.Context, userID uint)(*response.UserResponse,error){
	user , err := s.repo.GetByID(ctx,userID)
	if err != nil {
		return nil, err 
	}
	return s.toUserResponse(user),nil
}

func (s *userService) DeleteUser(ctx context.Context , userID uint)error{
	return s.repo.Delete(ctx,userID)
}

func (s *userService) generateAuthResponse(user *domain.User)(*response.AuthResponse,error){
	accessToken , err := s.jwtService.GenerateToken(user.UUID,user.Email,string(user.Role))
	if err != nil {
		return  nil , err
	}
	refreshToken , err := s.jwtService.GenerateRefreshToken(user.UUID)
	if err != nil {
		return nil ,err
	}
	return &response.AuthResponse{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
		TokenType: "Bearer",
		ExpiresIn: 3600,
		User: s.toUserResponse(user),
	},nil 
}
func (s *userService) toUserResponse(user *domain.User)*response.UserResponse{
	return &response.UserResponse{
		Name: user.User_name,
		Email: user.Email,
		Role: string(user.Role),
	}
}