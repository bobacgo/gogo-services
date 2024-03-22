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

type MenuResponse struct {
	*model.Menu
}

type MenuCreateRequest struct {
}

type MenuCreateResponse struct {
}

type MenuUpdateRequest struct {
	ID int64 `json:"id"`
}

type MenuUpdateResponse struct {
}

type MenuUpdateHiddenRequest struct {
	ID     int64 `json:"id"`
	Hidden int   `json:"hidden" form:"hidden"`
}

type MenuDeleteRequest struct {
	ID int64 `json:"id"`
}

type MenuDeleteResponse struct {
}
