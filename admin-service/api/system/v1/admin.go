package v1

import "github.com/gin-gonic/gin"

type AdminServer interface {
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

type AdminRequest struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	Icon     string `json:"icon"`     // 头像
	Email    string `json:"email"`    // 邮箱
	Note     string `json:"note"`     // 备注
}

type AdminResponse struct {
}
