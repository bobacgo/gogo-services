package errs

import (
	"github.com/gogoclouds/gogo-services/common-lib/web/r/codes"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/status"
)

// admin

var (
	AdminNotFound = status.New(codes.Code(400034), "角色不存在")
)

// menu

// role
