package handler

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/gogoclouds/gogo-services/admin-service/api/system/v1"
	"github.com/gogoclouds/gogo-services/admin-service/internal/domain/system/service"
	"github.com/gogoclouds/gogo-services/common-lib/web/r"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/errs"
)

type menuServer struct {
	svc *service.MenuService
}

var _ v1.MenuServer = (*menuServer)(nil)

func NewMenuServer(svc *service.MenuService) v1.MenuServer {
	return &menuServer{svc: svc}
}

func (h *menuServer) List(ctx *gin.Context) {
	req := new(v1.MenuListRequest)
	if err := ctx.ShouldBindUri(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	if err := ctx.ShouldBind(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	list, err := h.svc.List(ctx, req)
	if err != nil {
		r.Reply(ctx, err)
		return
	}
	r.Reply(ctx, list)
}

func (h *menuServer) TreeList(ctx *gin.Context) {
	list, err := h.svc.TreeList(ctx)
	if err != nil {
		r.Reply(ctx, err)
		return
	}
	r.Reply(ctx, list)
}

func (h *menuServer) Details(ctx *gin.Context) {
	req := new(v1.MenuRequest)
	if err := ctx.ShouldBindUri(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	data, err := h.svc.GetDetails(ctx, req)
	if err != nil {
		r.Reply(ctx, err)
		return
	}
	r.Reply(ctx, data)
}

func (h *menuServer) Add(ctx *gin.Context) {
	req := new(v1.MenuCreateRequest)
	if err := ctx.ShouldBind(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	if err := h.svc.Add(ctx, req); err != nil {
		r.Reply(ctx, err)
		return
	}
	r.Reply(ctx, nil)
}

func (h *menuServer) Update(ctx *gin.Context) {
	req := new(v1.MenuUpdateRequest)
	if err := ctx.ShouldBind(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	if err := h.svc.Update(ctx, req); err != nil {
		r.Reply(ctx, err)
		return
	}
	r.Reply(ctx, nil)
}

func (h *menuServer) UpdateHidden(ctx *gin.Context) {
	req := new(v1.MenuUpdateHiddenRequest)
	if err := ctx.ShouldBind(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	if err := h.svc.UpdateHidden(ctx, req.ID, req.Hidden); err != nil {
		r.Reply(ctx, err)
		return
	}
	r.Reply(ctx, nil)
}

func (h *menuServer) Delete(ctx *gin.Context) {
	req := new(v1.MenuDeleteRequest)
	if err := ctx.ShouldBind(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	if err := h.svc.Delete(ctx, req); err != nil {
		r.Reply(ctx, err)
		return
	}
	r.Reply(ctx, nil)
}
