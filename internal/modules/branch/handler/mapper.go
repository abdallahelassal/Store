package handler

import (
	"github.com/abdallahelassal/Store/internal/modules/branch/domain"
	"github.com/abdallahelassal/Store/internal/modules/branch/dtos/response"
)

func ToBranchResponse(branch *domain.Branch) *response.BranchResponse {
	return &response.BranchResponse{
		UUID: branch.UUID.String(),
		Name: branch.Name,
	}
}
func ToBranchListResponse(branches []*domain.Branch,total int64) *response.BranchListResponse{
	list := make([]*response.BranchResponse, 0, len(branches))

	for _, v := range branches{
		list = append(list,ToBranchResponse(v))
	}
	return &response.BranchListResponse{
		Data:  list,
		Total: total,
	}

}