package handler

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/gogoclouds/gogo-services/admin-service/api/system/v1"
)

type adminServer struct{}

func NewAdminServer() v1.AdminServer {
	return &adminServer{}
}

func (h *adminServer) Register(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *adminServer) Login(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *adminServer) Logout(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *adminServer) RefreshToken(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *adminServer) GetAdminInfo(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *adminServer) List(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *adminServer) GetItem(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *adminServer) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *adminServer) UpdatePassword(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *adminServer) Delete(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *adminServer) UpdateStatus(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *adminServer) UpdateRole(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *adminServer) GetRoleList(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
