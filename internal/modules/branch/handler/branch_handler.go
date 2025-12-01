package handler

import (
	"net/http"
	"strconv"

	"github.com/abdallahelassal/Store/internal/modules/branch/dtos/request"
	"github.com/abdallahelassal/Store/internal/modules/branch/service"
	"github.com/abdallahelassal/Store/pkg/utils"
	"github.com/gin-gonic/gin"
)

type BranchHandler struct {
	service service.BranchService
}

func NewBranchHandler(service service.BranchService)*BranchHandler{
	return &BranchHandler{service: service}
}

func (h *BranchHandler) CreateBranch(c *gin.Context){
	var req request.CreateBranchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid request", err.Error())
		return
	}

	 err := h.service.CreateBranch(c.Request.Context(),req)
	if err != nil {
		utils.ErrorResponse(c,http.StatusBadRequest,"failed create branch", err.Error())
		return
	}
	utils.SuccessResponse(c,http.StatusCreated,"branch created successfully", nil)
}

func (h *BranchHandler) DeleteBranch(c *gin.Context){
	uuid := c.Param("uuid")
	err := h.service.DeleteBranch(c.Request.Context(),uuid)
	if err != nil {
		utils.ErrorResponse(c,http.StatusBadRequest,"failed delete branch", err.Error())
		return
	}
	utils.SuccessResponse(c,http.StatusOK,"branch deleted successfully", nil)
}

func (h *BranchHandler) GetBranch(c *gin.Context){
	uuid := c.Param("uuid")
	branch, err := h.service.GetBranch(c.Request.Context(),uuid)
	if err != nil {
		utils.ErrorResponse(c,http.StatusBadRequest,"failed get branch", err.Error())
		return
	}
	res := ToBranchResponse(branch)
	utils.SuccessResponse(c,http.StatusOK,"branch fetched successfully", res)
}

func (h *BranchHandler) ListBranches(c *gin.Context){
	limitStr := c.DefaultQuery("limit","10")
	offsetStr := c.DefaultQuery("offset","0")

	limit , err := strconv.Atoi(limitStr)
	if err != nil {
		utils.ErrorResponse(c,http.StatusBadRequest,"invalid limit", err.Error())
		return
	}
	offset , err := strconv.Atoi(offsetStr)
	if err != nil {
		utils.ErrorResponse(c,http.StatusBadRequest,"invalid offset", err.Error())
		return
	}
	if limit > 100 {
		limit = 100
	}
	branch , total , err := h.service.ListBranches(c.Request.Context(),limit,offset)
	if err != nil {
		utils.ErrorResponse(c,http.StatusBadRequest,"failed list branches", err.Error())
		return
	}
	res := ToBranchListResponse(branch,total)
	utils.SuccessResponse(c,http.StatusOK,"branches listed successfully", res)
}