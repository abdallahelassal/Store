package repository

import (
	"context"

	"github.com/abdallahelassal/Store/internal/modules/branch/domain"
	"gorm.io/gorm"
)

type branchRepository struct {
	db *gorm.DB
}



func NewBranchRepository(db *gorm.DB)BranchRepository{
	return &branchRepository{db: db}
}


func (b *branchRepository) CreateBranch(ctx context.Context,branch *domain.Branch)error{
	return b.db.WithContext(ctx).Create(branch).Error
}

func (b *branchRepository) DeleteBranch(ctx context.Context,branchUUID string)error{
	return b.db.WithContext(ctx).
	Where("uuid = ?", branchUUID).
	Delete(&domain.Branch{}).Error
}

func (b *branchRepository) GetBranch(ctx context.Context,branchUUID string)(*domain.Branch,error){
	var branch domain.Branch
	err := b.db.WithContext(ctx).
	Where("uuid = ?", branchUUID).
	First(&branch).Error
	if err != nil{
		if err == gorm.ErrRecordNotFound{
			return nil,domain.ErrBranchNotFound
		}
		return nil,err
	}
	return &branch,nil
}

func (b *branchRepository) GetBranchByName(ctx context.Context,name string)(*domain.Branch,error){
	var branch domain.Branch
	if err := b.db.WithContext(ctx).
	Where("name = ?", name).
	First(&branch).Error;err != nil {
		return nil , domain.ErrBranchNotFound
	}
	return &branch,nil
}


func (b *branchRepository) ListBranches(ctx context.Context,limit,offset int)([]*domain.Branch,int64,error){
	var branches []*domain.Branch
	var total int64
	if err := b.db.WithContext(ctx).
	Model(&domain.Branch{}).
	Count(&total).Error; err != nil {
		return nil,0, err
	}

	err := b.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Order("created_at DESC").
		Find(&branches).Error
	if err != nil{
		return nil,0,err
	}	

	return branches,total, nil
}
