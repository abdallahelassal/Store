package repository

import(
	"context"
	"github.com/abdallahelassal/Store/internal/modules/user/domain"
	"gorm.io/gorm"
)

type userRepository struct{
	db 		*gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository{
	return&userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *domain.User)error{
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) GetByID(ctx context.Context,id uint)(*domain.User,error){
	var user domain.User
	err := r.db.WithContext(ctx).First(&user,id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			return nil,domain.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) GetByEmail(ctx context.Context,email string)(*domain.User,error){
	var user domain.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil{
		if err == gorm.ErrRecordNotFound{
			return nil,domain.ErrUserNotFound
		}
		return nil,err
	}
	return &user,nil
}
func (r *userRepository) Update(ctx context.Context,user *domain.User)error{
	return r.db.WithContext(ctx).Save(user).Error
}
func (r *userRepository) Delete(ctx context.Context,uuid string)error{
	return r.db.WithContext(ctx).Delete(&domain.User{}, uuid).Error
}
func (r *userRepository) List(ctx context.Context, limit,offset int)([]*domain.User,error){
	var users []*domain.User
	err := r.db.WithContext(ctx).
	Limit(limit).
	Offset(offset).
	Order("created_at DESC").
	Find(&users).Error
	return users,err
}

func (r *userRepository) Count(ctx context.Context)(int64,error){
	var count int64
	err := r.db.WithContext(ctx).Model(&domain.User{}).Count(&count).Error
	return count,err
}
func (r *userRepository) GetByUUID(ctx context.Context, uuid string) (*domain.User, error) {
    var user domain.User
    if err := r.db.WithContext(ctx).Where("uuid = ?", uuid).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *userRepository) UpdateUserBranch(ctx context.Context, userUUID, newBranchUUID string) error {
	return r.db.WithContext(ctx).
		Model(&domain.User{}).
		Where("uuid = ?", userUUID).
		Update("branch_uuid", newBranchUUID).Error
}