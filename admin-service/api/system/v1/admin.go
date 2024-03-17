package v1

import (
	"time"

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
	GetSelfInfo(ctx *gin.Context)
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
	Icon  string `json:"icon"`                             // 头像
	Email string `json:"email" validate:"omitempty,email"` // 邮箱
	Note  string `json:"note"`                             // 备注
}

type AdminLoginRequest struct {
	UsernamePasswd
}

type AdminLoginResponse struct {
	Token  string `json:"token"`
	RToken string `json:"rToken"`
}

type AdminPwdErr struct {
	DecrCount int64 `json:"decrCount"` // 剩余次数
}

type AdminRefreshTokenRequest struct {
	AToken string `json:"aToken" form:"aToken" validate:"required"`
	RToken string `json:"rToken" header:"authorization"`
}

type UserInfo struct {
	Username string   `json:"username"`
	NickName string   `json:"nickName"`
	Menus    []any    `json:"menus"`
	Icon     string   `json:"icon"`
	Roles    []string `json:"roles"`
}

type AdminListRequest struct {
	Keyword string `json:"keyword" form:"keyword"`
	page.Query
}

type AdminListResponse struct {
	page.Data[*model.Admin]
}

type AdminRequest struct {
	ID int64 `json:"id" uri:"id"`
}

type AdminResponse struct {
	ID         int64      `json:"id"`
	Username   string     `json:"username"`
	Icon       *string    `json:"icon"`      // 头像
	Email      *string    `json:"email"`     // 邮箱
	Nickname   *string    `json:"nickname"`  // 昵称
	Note       *string    `json:"note"`      // 备注信息
	Status     bool       `json:"status"`    // 帐号启用状态：0->禁用；1->启用
	LoginTime  *time.Time `json:"loginTime"` // 最后登录时间
	CreateTime *time.Time `son:"createTime"`
	UpdateTime *time.Time `json:"updateTime"`
}

type AdminUpdateRequest struct {
	ID       int64   `gorm:"-" json:"id"`
	Icon     *string `json:"icon"`                             // 头像
	Email    *string `json:"email" validate:"omitempty,email"` // 邮箱
	Nickname *string `json:"nickname"`                         // 昵称
	Note     *string `json:"note"`                             // 备注信息
}

type UpdatePasswordRequest struct {
	Username    string              `json:"username" validate:"required,lte=20"` // 用户名
	Password    security.Ciphertext `json:"oldPassword" validate:"required"`     // 密码
	NewPassword security.Ciphertext `json:"newPassword" validate:"required"`     // 新密码
}

type UpdateStatusRequest struct {
	ID     int64 `json:"id"`     // ID
	Status bool  `json:"status"` // 状态
}

type UpdateRoleRequest struct {
	ID    int64   `json:"id"`    // ID
	Roles []int64 `json:"roles"` // 角色
}
