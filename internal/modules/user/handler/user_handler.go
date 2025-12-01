package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/abdallahelassal/Store/internal/middleware"
	"github.com/abdallahelassal/Store/internal/modules/user/dtos/request"
	"github.com/abdallahelassal/Store/internal/modules/user/service"
	"github.com/abdallahelassal/Store/pkg"
	"github.com/abdallahelassal/Store/pkg/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid request", err.Error())
		return
	}

	result, err := h.service.Register(c.Request.Context(), req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "failed register", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "user registered successfully", result)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req request.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid request", err.Error())
		return
	}

	result, err := h.service.Login(c.Request.Context(), req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "login failed", err.Error())
		return
	}

	pkg.SetCookie(c,"access_token",result.AccessToken,time.Hour)
	utils.SuccessResponse(c, http.StatusOK, "login successfull", result)
}

func (h *UserHandler) RefreshToken(c *gin.Context){
	var req request.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c,http.StatusBadRequest,"invalid request", err.Error())
		return
	}
	if req.RefreshToken == ""{
		utils.ErrorResponse(c,http.StatusBadRequest,"refresh token is required","")
		return
	}
	result , err := h.service.RefreshToken(c.Request.Context(),req.RefreshToken)
	if err != nil {
		utils.ErrorResponse(c,http.StatusUnauthorized,"token refresh failed", err.Error())
		return
	}
	utils.SuccessResponse(c,http.StatusOK,"token refreshed successfully",result)
}

func (h *UserHandler) ListUsers(c *gin.Context){
	limitStr := c.DefaultQuery("limit","10")
	offsetStr := c.DefaultQuery("offset","0")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		utils.ErrorResponse(c,http.StatusBadRequest,"invalid limit value", "")
		return
	}
	offset , err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		utils.ErrorResponse(c,http.StatusBadRequest,"invalid offset value", "")
		return
	}
	if limit > 100 {
		limit = 100
	}
	users , err := h.service.ListUsers(c.Request.Context(),limit,offset)
	if err != nil {
		utils.ErrorResponse(c,http.StatusInternalServerError,"failed to fetch users",err.Error())
		return
	}
	utils.SuccessResponse(c,http.StatusOK,"users retrieved successfully",users)
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	authUserInterface, exists := c.Get("auth_user")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "user not authenticated", "")
		return
	}
	authUser, ok := authUserInterface.(middleware.AuthenticatedUser)

	if !ok {
		utils.ErrorResponse(c, http.StatusInternalServerError, "invalid user data", "")
		return
	}

	user, err := h.service.GetByUUID(c.Request.Context(), authUser.UUID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "user not found", "")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "user retrieved successfully", user)
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
		authUserInterface, exists := c.Get("auth_user")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "user not authenticated", "")
		return
	}
	authUser, ok := authUserInterface.(middleware.AuthenticatedUser)

	if !ok {
		utils.ErrorResponse(c, http.StatusInternalServerError, "invalid user data", "")
		return
	}
	var req request.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "invalid request", err.Error())
		return
	}
	profile, err := h.service.UpdateProfile(c.Request.Context(), authUser.UUID, req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "update failed", err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "profile updated successfully", profile)
}

func (h *UserHandler) DeleteUser(c *gin.Context){
	authUserInterface , exists := c.Get("auth_user")
	if !exists {
		utils.ErrorResponse(c,http.StatusUnauthorized,"unauthorized","")
		return
	}
	authUser , ok := authUserInterface.(middleware.AuthenticatedUser)
	if !ok{
		utils.ErrorResponse(c,http.StatusInternalServerError,"invalid user data", "")
		return
	}
	if err := h.service.DeleteUser(c.Request.Context(),authUser.UUID); err != nil{
		utils.ErrorResponse(c,http.StatusInternalServerError,"failed to delete user", err.Error())
		return
	}
	utils.SuccessResponse(c,http.StatusOK,"user deleted successfully",nil)
}
func (h *UserHandler) MoveUser(c *gin.Context){
	var req request.UpdateUserBranchRequest	
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c,http.StatusBadRequest,"invalid request", err.Error())
		return
	}
	if err := h.service.UpdateUserBranch(c.Request.Context(), req.UserUUID, req.NewBranchUUID); err != nil {
		utils.ErrorResponse(c,http.StatusInternalServerError,"failed to update user branch", err.Error())
		return
	}
	utils.SuccessResponse(c,http.StatusOK,"user branch updated successfully",nil)
}