package v1

import (
	"context"

	"github.com/gogoclouds/gogo-services/admin-service/internal/domain/system/dto"
	"github.com/gogoclouds/gogo-services/admin-service/internal/model"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/page"
)

type IMenuServer interface {
	List(ctx context.Context, req *MenuListRequest) (*page.Data[*model.Menu], error)
	TreeList(ctx context.Context) ([]*dto.MenuNode, error)
	GetDetails(ctx context.Context, req *MenuRequest) (*model.Menu, error)
	Add(ctx context.Context, req *MenuCreateRequest) error
	Update(ctx context.Context, req *MenuUpdateRequest) error
	UpdateHidden(ctx context.Context, ID int64, hidden *bool) error
	Delete(ctx context.Context, req *MenuDeleteRequest) error
}

type MenuListRequest struct {
	page.Query
	ParentID *int64 `json:"parentId" uri:"parentId"`
}

type MenuListResponse struct {
	page.Data[*model.Menu]
}

type MenuRequest struct {
	ID int64 `json:"id" uri:"id"`
}

type MenuCreateRequest struct {
	ParentID int64   `json:"parentId"`                  // 父级ID
	Title    *string `json:"title" validate:"required"` // 菜单名称
	Level    *int32  `json:"level"`                     // 菜单级数
	Sort     *int32  `json:"sort"`                      // 菜单排序
	Name     string  `json:"name" validate:"required"`  // 前端名称
	Icon     *string `json:"icon" validate:"required"`  // 前端图标
	Hidden   bool    `json:"hidden"`                    // 前端隐藏
}

type MenuCreateResponse struct {
}

type MenuUpdateRequest struct {
	ID       int64   `json:"id"`
	ParentID *int64  `json:"parentId"` // 父级ID
	Title    *string `json:"title"`    // 菜单名称
	Level    *int32  `json:"level"`    // 菜单级数
	Sort     *int32  `json:"sort"`     // 菜单排序
	Name     string  `json:"name"`     // 前端名称
	Icon     *string `json:"icon"`     // 前端图标
	Hidden   *bool   `json:"hidden"`   // 前端隐藏
}

type MenuUpdateResponse struct {
}

type MenuUpdateHiddenRequest struct {
	ID     int64 `json:"id"`
	Hidden *bool `json:"hidden" form:"hidden" validate:"required"`
}

type MenuDeleteRequest struct {
	ID int64 `json:"id" uri:"id"`
}

type MenuDeleteResponse struct {
}
