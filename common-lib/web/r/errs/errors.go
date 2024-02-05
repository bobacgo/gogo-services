package errs

import (
	"github.com/gogoclouds/gogo-services/common-lib/web/r/status"
)

var UsernameNotFound = status.New(50034, "用户名不存在")
var BadRequest = status.New(400, "请求参数错误")
