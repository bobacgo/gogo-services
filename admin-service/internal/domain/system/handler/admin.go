package handler

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/gogoclouds/gogo-services/admin-service/api/system/v1"
	"github.com/gogoclouds/gogo-services/admin-service/internal/domain/system/service"
	"github.com/gogoclouds/gogo-services/admin-service/internal/model"
	"github.com/gogoclouds/gogo-services/admin-service/internal/router/middleware"
	"github.com/gogoclouds/gogo-services/common-lib/app/security"
	"github.com/gogoclouds/gogo-services/common-lib/web/r"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/errs"
)

type adminServer struct {
	svc *service.AdminService
}

func NewAdminServer(svc *service.AdminService) v1.AdminServer {
	return &adminServer{svc: svc}
}

func (h *adminServer) Register(ctx *gin.Context) {
	req := new(v1.AdminRegisterRequest)
	if err := ctx.ShouldBind(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	if err := req.Password.Decrypt(); err != nil {
		r.Reply(ctx, err)
		return
	}
	err := h.svc.Register(ctx, req)
	r.Reply(ctx, err)
}

func (h *adminServer) Login(ctx *gin.Context) {
	req := new(v1.AdminLoginRequest)
	if err := ctx.ShouldBind(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}

	if err := req.Password.Decrypt(); err != nil {
		r.Reply(ctx, err)
		return
	}
	data, err := h.svc.Login(ctx, req)
	if err != nil {
		r.Reply(ctx, err)
		return
	}
	r.Reply(ctx, data)
}

func (h *adminServer) Logout(ctx *gin.Context) {
	username := security.GetUsername(ctx)
	err := h.svc.Logout(ctx, username)
	r.Reply(ctx, err)
}

// RefreshToken 请求头携带rToken
func (h *adminServer) RefreshToken(ctx *gin.Context) {
	req := new(v1.AdminRefreshTokenRequest)
	if err := ctx.ShouldBind(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	auth := middleware.GetAuthHeader(ctx)
	if auth == "" {
		r.Reply(ctx, errs.BadRequest.WithDetails("Authorization header is empty"))
		return
	}
	req.RToken = auth
	data, err := h.svc.RefreshToken(ctx, req)
	if err != nil {
		r.Reply(ctx, err)
		return
	}
	r.Reply(ctx, data)
}

func (h *adminServer) GetSelfInfo(ctx *gin.Context) {
	username := security.GetUsername(ctx)
	data, err := h.svc.GetAdminInfo(ctx, username)
	if err != nil {
		r.Reply(ctx, err)
		return
	}
	r.Reply(ctx, data)
}

func (h *adminServer) List(ctx *gin.Context) {
	req := new(v1.ListRequest)
	if err := ctx.ShouldBind(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	data, err := h.svc.List(ctx, req)
	if err != nil {
		r.Reply(ctx, err)
		return
	}
	r.Reply(ctx, data)
}

func (h *adminServer) GetItem(ctx *gin.Context) {
	userID := security.GetUserIntID(ctx)
	data, err := h.svc.GetItem(ctx, userID)
	if err != nil {
		r.Reply(ctx, err)
		return
	}
	r.Reply(ctx, data)
}

func (h *adminServer) Update(ctx *gin.Context) {
	req := new(model.Admin)
	if err := ctx.ShouldBind(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	err := h.svc.Update(ctx, req)
	r.Reply(ctx, err)
}

func (h *adminServer) UpdatePassword(ctx *gin.Context) {
	req := new(v1.UpdatePasswordRequest)
	if err := ctx.ShouldBind(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}

	if err := req.Password.Decrypt(); err != nil {
		r.Reply(ctx, err)
		return
	}
	err := h.svc.UpdatePassword(ctx, req)
	r.Reply(ctx, err)
}

func (h *adminServer) Delete(ctx *gin.Context) {
	userID := security.GetUserIntID(ctx)
	err := h.svc.Delete(ctx, userID)
	r.Reply(ctx, err)
}

func (h *adminServer) UpdateStatus(ctx *gin.Context) {
	req := new(v1.UpdateStatusRequest)
	if err := ctx.ShouldBind(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	err := h.svc.UpdateStatus(ctx, req.ID, req.Status)
	r.Reply(ctx, err)
}

func (h *adminServer) UpdateRole(ctx *gin.Context) {
	req := new(v1.UpdateRoleRequest)
	if err := ctx.ShouldBind(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	err := h.svc.UpdateRole(ctx, req.ID, req.Roles)
	r.Reply(ctx, err)
}

func (h *adminServer) GetRoleList(ctx *gin.Context) {
	userID := security.GetUserIntID(ctx)
	data, err := h.svc.GetRoleList(ctx, userID)
	if err != nil {
		r.Reply(ctx, err)
		return
	}
	r.Reply(ctx, data)
}
