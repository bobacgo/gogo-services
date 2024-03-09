package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/admin-service/internal/model"
	"github.com/gogoclouds/gogo-services/common-lib/app/security"
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

type UsernamePasswd struct {
	Username string              `json:"username" validate:"required,lte=20"` // 用户名
	Password security.Ciphertext `json:"password" validate:"required"`        // 密码
}

type AdminRegisterRequest struct {
	UsernamePasswd
	Icon  string `json:"icon"`                                // 头像
	Email string `json:"email" validate:"omitempty,required"` // 邮箱
	Note  string `json:"note"`                                // 备注
}

type AdminLoginRequest struct {
	UsernamePasswd
}

type AdminLoginResponse struct {
	Token string `json:"token"`
}

type AdminPwdErr struct {
	DecrCount int64 `json:"decrCount"` // 剩余次数
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
	UsernamePasswd
	Password    security.Ciphertext `json:"oldPassword" validate:"required"` // 密码
	NewPassword security.Ciphertext `json:"newPassword" validate:"required"` // 新密码
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
