package repository

import (
	"context"

	"github.com/abdallahelassal/Store/internal/modules/branch/domain"
)

type BranchRepository interface {
	// Define methods for branch repository here
	CreateBranch(ctx context.Context, branch *domain.Branch) error
	DeleteBranch(ctx context.Context, branchUUID string) error
	GetBranch(ctx context.Context , branchUUID string) (*domain.Branch, error)
	GetBranchByName(ctx context.Context, name string) (*domain.Branch,error)
	ListBranches(ctx context.Context, limit,offset int) ([]*domain.Branch,int64,error)
}