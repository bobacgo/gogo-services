package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/admin-service/internal/model"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/page"
)

type MenuServer interface {
	List(ctx *gin.Context)
	TreeList(ctx *gin.Context)
	Details(ctx *gin.Context)
	Add(ctx *gin.Context)
	Update(ctx *gin.Context)
	UpdateHidden(ctx *gin.Context)
	Delete(ctx *gin.Context)
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
