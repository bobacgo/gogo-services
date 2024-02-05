package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/page"
	"github.com/gogoclouds/gogo-services/gen/internal/mvc/sample/internal/model"
)

type UserServer interface {
	List(ctx *gin.Context)
	Details(ctx *gin.Context)
	Add(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type UserListRequest struct {
	page.Query
}

type UserListResponse struct {
	page.Data[*model.SysUser]
}

type UserRequest struct {
	ID uint32 `json:"id"`
}

type UserResponse struct {
	*model.SysUser
}

type UserCreateRequest struct {
}

type UserCreateResponse struct {
}

type UserUpdateRequest struct {
	ID uint32 `json:"id"`
}

type UserUpdateResponse struct {
}

type UserDeleteRequest struct {
	ID uint32 `json:"id"`
}

type UserDeleteResponse struct {
}
