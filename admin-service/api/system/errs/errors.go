package errs

import (
	"github.com/gogoclouds/gogo-services/common-lib/web/r/status"
)

// admin

var (
	AdminUsernameDuplicated   = status.New(400001, "用户名重复")
	AdminUnUsernameDuplicated = status.New(400002, "用户名重复")
	AdminLoginFail            = status.New(400003, "登录名或密码不正确")
	AdminNotFound             = status.New(400004, "找不到该用户")
	AdminLoginForbidden       = status.New(400005, "帐号已被禁用")
	AdminOldPwdError          = status.New(400006, "旧密码错误")
)

// menu

// role