package handler

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/gogoclouds/gogo-services/admin-service/api/system/v1"
	"github.com/gogoclouds/gogo-services/admin-service/internal/config"
	"github.com/gogoclouds/gogo-services/admin-service/internal/router/middleware"
	"github.com/gogoclouds/gogo-services/framework/app/security"
	"github.com/gogoclouds/gogo-services/framework/web/r"
	"github.com/gogoclouds/gogo-services/framework/web/r/errs"
)

type AdminApi struct {
	svc v1.IAdminServer
}

func NewAdminServer(svc v1.IAdminServer) *AdminApi {
	return &AdminApi{svc: svc}
}

func (h *AdminApi) Register(ctx *gin.Context) {
	req := new(v1.AdminRegisterRequest)
	if err := ctx.ShouldBind(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	if err := req.Password.Decrypt(config.Cfg.Security); err != nil {
		r.Reply(ctx, err)
		return
	}
	err := h.svc.Register(ctx, req)
	r.Reply(ctx, err)
}

func (h *AdminApi) Login(ctx *gin.Context) {
	req := new(v1.AdminLoginRequest)
	if err := ctx.ShouldBind(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}

	if err := req.Password.Decrypt(config.Cfg.Security); err != nil {
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

func (h *AdminApi) Logout(ctx *gin.Context) {
	username := security.GetUsername(ctx)
	err := h.svc.Logout(ctx, &v1.AdminLogoutRequest{Username: username})
	r.Reply(ctx, err)
}

// RefreshToken 请求头携带rToken
func (h *AdminApi) RefreshToken(ctx *gin.Context) {
	req := new(v1.AdminRefreshTokenRequest)
	if err := ctx.ShouldBind(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}

	req.RToken = middleware.TrimToken(ctx.GetString(middleware.AuthHeader))
	if req.RToken == "" {
		r.Reply(ctx, errs.BadRequest.WithDetails("Authorization header is empty"))
		return
	}
	data, err := h.svc.RefreshToken(ctx, req)
	if err != nil {
		r.Reply(ctx, err)
		return
	}
	r.Reply(ctx, data)
}

func (h *AdminApi) GetSelfInfo(ctx *gin.Context) {
	username := security.GetUsername(ctx)
	data, err := h.svc.GetAdminInfo(ctx, &v1.AdminInfoRequest{Username: username})
	if err != nil {
		r.Reply(ctx, err)
		return
	}
	r.Reply(ctx, data)
}

func (h *AdminApi) List(ctx *gin.Context) {
	req := new(v1.AdminListRequest)
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

func (h *AdminApi) GetItem(ctx *gin.Context) {
	req := new(v1.AdminRequest)
	if err := ctx.ShouldBindUri(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	data, err := h.svc.GetItem(ctx, req)
	if err != nil {
		r.Reply(ctx, err)
		return
	}
	r.Reply(ctx, data)
}

func (h *AdminApi) Update(ctx *gin.Context) {
	req := new(v1.AdminUpdateRequest)
	if err := ctx.ShouldBindUri(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	if err := ctx.ShouldBind(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	err := h.svc.Update(ctx, req)
	r.Reply(ctx, err)
}

func (h *AdminApi) UpdatePassword(ctx *gin.Context) {
	req := new(v1.UpdatePasswordRequest)
	if err := ctx.ShouldBind(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}

	if err := req.Password.Decrypt(config.Cfg.Security); err != nil {
		r.Reply(ctx, err)
		return
	}
	err := h.svc.UpdatePassword(ctx, req)
	r.Reply(ctx, err)
}

func (h *AdminApi) Delete(ctx *gin.Context) {
	req := new(v1.AdminRequest)
	if err := ctx.ShouldBindUri(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	err := h.svc.Delete(ctx, req)
	r.Reply(ctx, err)
}

func (h *AdminApi) UpdateStatus(ctx *gin.Context) {
	req := new(v1.AdminUpdateStatusRequest)
	if err := ctx.ShouldBindUri(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	if err := ctx.ShouldBind(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	err := h.svc.UpdateStatus(ctx, req)
	r.Reply(ctx, err)
}

func (h *AdminApi) UpdateRole(ctx *gin.Context) {
	req := new(v1.AdminUpdateRoleRequest)
	if err := ctx.ShouldBind(req); err != nil {
		r.Reply(ctx, errs.BadRequest.WithDetails(err))
		return
	}
	err := h.svc.UpdateRole(ctx, req)
	r.Reply(ctx, err)
}

func (h *AdminApi) GetRoleList(ctx *gin.Context) {
	userID := security.GetUserIntID(ctx)
	data, err := h.svc.GetRoleList(ctx, &v1.AdminRequest{ID: userID})
	if err != nil {
		r.Reply(ctx, err)
		return
	}
	r.Reply(ctx, data)
}
