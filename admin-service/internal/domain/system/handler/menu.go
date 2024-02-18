package handler

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	v1 "github.com/gogoclouds/gogo-services/admin-service/api/system/v1"
	"github.com/gogoclouds/gogo-services/admin-service/internal/domain/system/service"
	"github.com/gogoclouds/gogo-services/common-lib/web/r"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/errs"
)

type MenuServer struct {
	svc *service.MenuService
}

func NewMenuServer(svc *service.MenuService) v1.MenuServer {
	return &MenuServer{svc: svc}
}

func (h *MenuServer) List(ctx *gin.Context) {
	req := new(v1.MenuListRequest)
	if err := ctx.ShouldBind(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	list, err := h.svc.List(ctx, req)
	if err != nil {
		slog.ErrorContext(ctx, "List error", slog.String("err", err.Error()))
		r.Reply(ctx, err)
		return
	}
	r.Reply(ctx, list)
}

func (h *MenuServer) TreeList(ctx *gin.Context) {
	list, err := h.svc.TreeList(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "List error", slog.String("err", err.Error()))
		r.Reply(ctx, err)
		return
	}
	r.Reply(ctx, list)
}

func (h *MenuServer) Details(ctx *gin.Context) {
	req := new(v1.MenuRequest)
	if err := ctx.ShouldBind(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	data, err := h.svc.GetDetails(ctx, req)
	if err != nil {
		slog.ErrorContext(ctx, "Details error", slog.String("err", err.Error()))
		r.Reply(ctx, err)
		return
	}
	r.Reply(ctx, data)
}

func (h *MenuServer) Add(ctx *gin.Context) {
	req := new(v1.MenuCreateRequest)
	if err := ctx.ShouldBind(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	if err := h.svc.Add(ctx, req); err != nil {
		slog.ErrorContext(ctx, "Add error", slog.String("err", err.Error()))
		r.Reply(ctx, err)
		return
	}
	r.Reply(ctx, nil)
}

func (h *MenuServer) Update(ctx *gin.Context) {
	req := new(v1.MenuUpdateRequest)
	if err := ctx.ShouldBind(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	if err := h.svc.Update(ctx, req); err != nil {
		slog.ErrorContext(ctx, "Update error", slog.String("err", err.Error()))
		r.Reply(ctx, err)
		return
	}
	r.Reply(ctx, nil)
}

func (h *MenuServer) UpdateHidden(ctx *gin.Context) {
	req := new(v1.MenuUpdateHiddenRequest)
	if err := ctx.ShouldBind(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	if err := h.svc.UpdateHidden(ctx, req.ID, req.Hidden == 1); err != nil {
		slog.ErrorContext(ctx, "Update error", slog.String("err", err.Error()))
		r.Reply(ctx, err)
		return
	}
	r.Reply(ctx, nil)
}

func (h *MenuServer) Delete(ctx *gin.Context) {
	req := new(v1.MenuDeleteRequest)
	if err := ctx.ShouldBind(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	if err := h.svc.Delete(ctx, req); err != nil {
		slog.ErrorContext(ctx, "Delete error", slog.String("err", err.Error()))
		r.Reply(ctx, err)
		return
	}
	r.Reply(ctx, nil)
}
