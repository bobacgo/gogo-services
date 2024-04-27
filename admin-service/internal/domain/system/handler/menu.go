package handler

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/gogoclouds/gogo-services/admin-service/api/system/v1"
	"github.com/gogoclouds/gogo-services/framework/web/r"
	"github.com/gogoclouds/gogo-services/framework/web/r/errs"
)

type MenuApi struct {
	svc v1.IMenuServer
}

func NewMenuServer(svc v1.IMenuServer) *MenuApi {
	return &MenuApi{svc: svc}
}

func (h *MenuApi) List(ctx *gin.Context) {
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

func (h *MenuApi) TreeList(ctx *gin.Context) {
	list, err := h.svc.TreeList(ctx)
	if err != nil {
		r.Reply(ctx, err)
		return
	}
	r.Reply(ctx, list)
}

func (h *MenuApi) Details(ctx *gin.Context) {
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

func (h *MenuApi) Add(ctx *gin.Context) {
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

func (h *MenuApi) Update(ctx *gin.Context) {
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

func (h *MenuApi) UpdateHidden(ctx *gin.Context) {
	req := new(v1.MenuUpdateHiddenRequest)
	if err := ctx.ShouldBind(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	if err := h.svc.UpdateHidden(ctx, req); err != nil {
		r.Reply(ctx, err)
		return
	}
	r.Reply(ctx, nil)
}

func (h *MenuApi) Delete(ctx *gin.Context) {
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
