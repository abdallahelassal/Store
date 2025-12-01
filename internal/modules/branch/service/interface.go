package service

import (
	"context"

	"github.com/abdallahelassal/Store/internal/modules/branch/domain"
	"github.com/abdallahelassal/Store/internal/modules/branch/dtos/request"
	
)

type BranchService interface {
	CreateBranch(ctx context.Context,req request.CreateBranchRequest) error
	DeleteBranch(ctx context.Context,uuid string)error
	GetBranch(ctx context.Context,uuid string)(*domain.Branch,error)
	ListBranches(ctx context.Context,limit,offset int)([]*domain.Branch,int64,error)
}
