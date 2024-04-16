package v1

import (
	"context"

	"github.com/gogoclouds/gogo-services/admin-service/internal/model"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/page"
)

type IRoleServer interface {
	List(ctx context.Context, req *RoleListRequest) (*page.Data[*model.Role], error)
	GetDetails(ctx context.Context, req *RoleRequest) (*RoleResponse, error)
	Add(ctx context.Context, req *RoleCreateRequest) error
	Update(ctx context.Context, req *RoleUpdateRequest) error
	Delete(ctx context.Context, req *RoleDeleteRequest) error
	UpdateStatus(ctx context.Context, id int64, status bool) error
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
