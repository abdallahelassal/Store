package service

import (
	"context"
	"log"

	"github.com/abdallahelassal/Store/internal/modules/branch/domain"
	"github.com/abdallahelassal/Store/internal/modules/branch/dtos/request"
	"github.com/abdallahelassal/Store/internal/modules/branch/repository"
)

type branchService struct {
	repo repository.BranchRepository
}

func NewBranchService(repo repository.BranchRepository)BranchService{
	return&branchService{repo:repo}
}

func (b *branchService) CreateBranch(ctx context.Context, req request.CreateBranchRequest)error{
	exiting , err := b.repo.GetBranchByName(ctx,req.Name)
	if err == nil && exiting != nil {
		return domain.ErrBranchAlreadyExiting
	}
	
	branch := &domain.Branch{
		Name: req.Name,
	}
	if err := b.repo.CreateBranch(ctx,branch);err != nil {
		log.Printf("faild create branch")
		return  err
	}
	return nil
}

func (b *branchService) DeleteBranch(ctx context.Context, uuid string)error{
	return b.repo.DeleteBranch(ctx,uuid)
}

func (b *branchService) GetBranch(ctx context.Context, uuid string)(*domain.Branch, error){
	branch, err := b.repo.GetBranch(ctx,uuid)
	if err != nil {
		return nil , domain.ErrBranchNotFound
	}
	return branch,nil
}

func (b *branchService) ListBranches(ctx context.Context, limit,offset int)([]*domain.Branch,int64, error){
	branches,total , err := b.repo.ListBranches(ctx, limit, offset)
	if err != nil {
		return nil,0, err
	}
	return branches,total, nil
}