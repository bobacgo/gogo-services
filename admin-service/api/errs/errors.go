package errs

import (
	"github.com/gogoclouds/gogo-services/framework/web/r/errs"
	"github.com/gogoclouds/gogo-services/framework/web/r/status"
)

// Token

var BadRequest = errs.BadRequest

var (
	TokenGenerateErr = status.New(401001, "令牌生成出错")
	TokenMiss        = status.New(401002, "令牌缺失")
	TokenInvalid     = status.New(401003, "令牌无效")
	TokenExpired     = status.New(401004, "令牌过期")
	TokenOut         = status.New(401005, "令牌失效") // Token 被覆盖(已登出)
)

// admin

var (
	AdminUsernameDuplicated   = status.New(400001, "用户名重复")
	AdminUnUsernameDuplicated = status.New(400002, "用户名重复")
	AdminLoginFail            = status.New(400003, "登录名或密码不正确")
	AdminNotFound             = status.New(400004, "找不到该用户")
	AdminLoginForbidden       = status.New(400005, "帐号已被禁用")
	AdminOldPwdErr            = status.New(400006, "旧密码错误")

	AdminEmailDuplicated = status.New(400008, "邮箱已经存在")
)

// menu

var (
	MenuNotFound = status.New(400100, "找不到该菜单")
)

// role
