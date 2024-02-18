package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/admin-service/internal/model"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/page"
)

type AdminServer interface {

	// Register
	// @Summary 注册
	// @Description 注册
	// @Tags 用户
	// @Param {object} AdminRegisterRequest
	// @Success 200 {object} AdminResponse
	// @Router /api/v1/sys-user [get]
	// @Security Bearer
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
	RefreshToken(ctx *gin.Context)
	GetAdminInfo(ctx *gin.Context)
	List(ctx *gin.Context)
	GetItem(ctx *gin.Context)
	Update(ctx *gin.Context)
	UpdatePassword(ctx *gin.Context)
	Delete(ctx *gin.Context)
	UpdateStatus(ctx *gin.Context)
	UpdateRole(ctx *gin.Context)
	GetRoleList(ctx *gin.Context)
}

type AdminRegisterRequest struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	Icon     string `json:"icon"`     // 头像
	Email    string `json:"email"`    // 邮箱
	Note     string `json:"note"`     // 备注
}

type AdminLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AdminLoginResponse struct {
	Token string `json:"token"`
}

type UserInfo struct {
	Username string   `json:"username"`
	NickName string   `json:"nickName"`
	Menus    string   `json:"menus"`
	Icon     string   `json:"icon"`
	Roles    []string `json:"roles"`
}

type ListRequest struct {
	Keyword string `json:"keyword" form:"keyword"`
	page.Query
}

type ListResponse struct {
	page.Data[*model.Admin]
}

type UpdatePasswordRequest struct {
	Username    string `json:"username"`    // 用户名
	Password    string `json:"oldPassword"` // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

type UpdateStatusRequest struct {
	ID     int64 `json:"id"`     // ID
	Status bool  `json:"status"` // 状态
}

type UpdateRoleRequest struct {
	ID    int64   `json:"id"`    // ID
	Roles []int64 `json:"roles"` // 角色
}

type AdminResponse struct {
	model.Admin
}