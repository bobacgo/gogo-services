package handler

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	v1 "github.com/gogoclouds/gogo-services/admin-service/api/system/v1"
	"github.com/gogoclouds/gogo-services/common-lib/web/r"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/errs"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/page"
)

type RoleApi struct {
	svc v1.IRoleServer
}

func NewRoleServer(svc v1.IRoleServer) *RoleApi {
	return &RoleApi{svc: svc}
}

func (h *RoleApi) List(ctx *gin.Context) {
	req := new(v1.RoleListRequest)
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

func (h *RoleApi) Details(ctx *gin.Context) {
	req := new(v1.RoleRequest)
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

func (h *RoleApi) Add(ctx *gin.Context) {
	req := new(v1.RoleCreateRequest)
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

func (h *RoleApi) Update(ctx *gin.Context) {
	req := new(v1.RoleUpdateRequest)
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

func (h *RoleApi) Delete(ctx *gin.Context) {
	req := new(v1.RoleDeleteRequest)
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

func (h *RoleApi) ListAll(ctx *gin.Context) {
	req := v1.RoleListRequest{
		Query: page.Query{
			PageNum:  -1,
			PageSize: -1,
		},
	}
	list, err := h.svc.List(ctx, &req)
	if err != nil {
		slog.ErrorContext(ctx, "List error", slog.String("err", err.Error()))
		r.Reply(ctx, err)
		return
	}
	r.Reply(ctx, list)
}

func (h *RoleApi) UpdateStatus(ctx *gin.Context) {
	req := new(v1.RoleUpdateStatusRequest)
	if err := ctx.ShouldBind(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	err := h.svc.UpdateStatus(ctx, req.ID, req.Status)
	if err != nil {
		slog.ErrorContext(ctx, "List error", slog.String("err", err.Error()))
		r.Reply(ctx, err)
		return
	}
	r.Reply(ctx, nil)
}
