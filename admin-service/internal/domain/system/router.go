package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/admin-service/internal/domain/system/handler"
	"github.com/gogoclouds/gogo-services/admin-service/internal/domain/system/repo"
	"github.com/gogoclouds/gogo-services/admin-service/internal/domain/system/service"
	"github.com/gogoclouds/gogo-services/admin-service/internal/router/middleware"
	"github.com/gogoclouds/gogo-services/common-lib/app"
)

func Register(app *app.App, r gin.IRouter) {
	db := app.Opts.DB
	rdb := app.Opts.Redis

	authRouter := r.Group("")
	authRouter.Use(middleware.Auth())
	{
		adminRepo := repo.NewAdminRepo(db)
		menuRepo := repo.NewMenuRepo(db)
		adminServer := handler.NewAdminServer(service.NewAdminService(rdb, adminRepo, adminRepo, menuRepo))
		r.POST("/admin/register", adminServer.Register)                      // 用户注册
		r.POST("/admin/login", adminServer.Login)                            // 登录以后返回token
		r.GET("/admin/refreshToken", adminServer.RefreshToken)               // 刷新token
		authRouter.GET("/admin/logout", adminServer.Logout)                  // 登出功能
		authRouter.POST("/admin/logout", adminServer.Logout)                 // 登出功能 (不推荐)
		authRouter.GET("/admin/info", adminServer.GetSelfInfo)               // 获取当前登录用户信息
		authRouter.GET("/admin/list", adminServer.List)                      // 根据用户名或姓名分页获取用户列表
		authRouter.GET("/admin/:id", adminServer.GetItem)                    // 获取指定用户信息
		authRouter.POST("/admin/update/:id", adminServer.Update)             // 修改指定用户信息
		authRouter.POST("/admin/updatePassword", adminServer.UpdatePassword) // 修改指定用户密码
		authRouter.POST("/admin/delete/:id", adminServer.Delete)             // 删除指定用户信息
		authRouter.POST("/admin/updateStatus/:id", adminServer.UpdateStatus) // 修改帐号状态
		authRouter.POST("/admin/role/update", adminServer.UpdateRole)        // 给用户分配角色
		authRouter.GET("/admin/role/:adminId", adminServer.GetRoleList)      // 获取指定用户的角色
	}
	{
		roleHandler := handler.NewRoleServer(service.NewRoleService(repo.NewRoleRepo(db)))
		authRouter.POST("/role/create", roleHandler.Add)                    // 添加角色
		authRouter.POST("/role/update/:id", roleHandler.Update)             // 修改角色
		authRouter.POST("/role/delete", roleHandler.Delete)                 // 批量删除角色
		authRouter.GET("/role/listAll", roleHandler.ListAll)                // 获取所有角色
		authRouter.GET("/role/list", roleHandler.List)                      // 根据角色名称分页获取角色列表
		authRouter.POST("/role/updateStatus/:id", roleHandler.UpdateStatus) // 修改角色状态
		authRouter.GET("/role/listMenu/:roleId", roleHandler.Update)        // 获取角色相关菜单
		authRouter.GET("/role/listResource/:roleId", roleHandler.Update)    // 获取角色相关资源
		authRouter.POST("/role/allocMenu", roleHandler.Update)              // 给角色分配菜单
		authRouter.POST("/role/allocResource", roleHandler.Update)          // 给角色分配资源
	}
	{
		menuHandler := handler.NewMenuServer(service.NewMenuService(repo.NewMenuRepo(db)))
		authRouter.GET("/menu/list/:parentId", menuHandler.List)
		authRouter.GET("/menu/treeList", menuHandler.TreeList)
		authRouter.GET("/menu/:id", menuHandler.Details)
		authRouter.POST("/menu/create", menuHandler.Add)
		authRouter.POST("/menu/update/:id", menuHandler.Update)
		authRouter.POST("/menu/updateHidden/:id", menuHandler.UpdateHidden)
		authRouter.POST("/menu/delete/:id", menuHandler.Delete)
	}
}
