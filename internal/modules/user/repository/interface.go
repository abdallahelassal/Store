package repository

import (
	"context"

	"github.com/abdallahelassal/Store/internal/modules/user/domain"
	
) 

type UserRepository interface{
	Create(ctx context.Context,user *domain.User) error
	GetByID(ctx context.Context, id uint)(*domain.User,error)
	GetByEmail(ctx context.Context,email string)(*domain.User,error)
	Update(ctx context.Context,user *domain.User) error
	Delete(ctx context.Context,id uint)error
	List(ctx context.Context, limit, offset int) ([]*domain.User, error)
	Count(ctx context.Context) (int64, error)
	GetByUUID(ctx context.Context, uuid string) (*domain.User, error) 
}