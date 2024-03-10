package errs

import (
	"github.com/gogoclouds/gogo-services/common-lib/web/r/codes"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/status"
)

var BadRequest = status.New(codes.BadRequest, "请求参数错误")
var InternalError = status.New(codes.InternalServerError, "服务器繁忙")
