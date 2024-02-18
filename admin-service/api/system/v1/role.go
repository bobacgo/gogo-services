package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/admin-service/internal/model"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/page"
)

type RoleServer interface {
	List(ctx *gin.Context)
	Details(ctx *gin.Context)
	Add(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	ListAll(ctx *gin.Context)
	UpdateStatus(ctx *gin.Context)
}

type RoleListRequest struct {
	page.Query
}

type RoleListResponse struct {
	page.Data[*model.Role]
}

type RoleRequest struct {
	ID int64 `json:"id"`
}

type RoleResponse struct {
	*model.Role
}

type RoleCreateRequest struct {
}

type RoleCreateResponse struct {
}

type RoleUpdateRequest struct {
	ID int64 `json:"id"`
}

type RoleUpdateStatusRequest struct {
	ID     int64 `json:"id"`
	Status bool  `json:"status"`
}

type RoleUpdateResponse struct {
}

type RoleDeleteRequest struct {
	IDs []int64 `json:"id"`
}

type RoleDeleteResponse struct {
}