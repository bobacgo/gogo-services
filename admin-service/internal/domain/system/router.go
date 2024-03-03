package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/admin-service/internal/domain/system/handler"
	"github.com/gogoclouds/gogo-services/admin-service/internal/domain/system/repo"
	"github.com/gogoclouds/gogo-services/admin-service/internal/domain/system/service"
	"github.com/gogoclouds/gogo-services/common-lib/app"
)

func Register(app *app.App, r gin.IRouter) {
	db := app.Opts.DB
	rdb := app.Opts.Redis

	{
		r := r.Group("/admin")
		adminRepo := repo.NewAdminRepo(db)
		adminServer := handler.NewAdminServer(service.NewAdminService(rdb, adminRepo, adminRepo))
		r.POST("/register", adminServer.Register)              // 用户注册
		r.POST("/login", adminServer.Login)                    // 登录以后返回token
		r.POST("/logout", adminServer.Logout)                  // 登出功能
		r.GET("/refreshToken", adminServer.RefreshToken)       // 刷新token
		r.GET("/info", adminServer.GetAdminInfo)               // 获取当前登录用户信息
		r.GET("/list", adminServer.List)                       // 根据用户名或姓名分页获取用户列表
		r.GET("/{id}", adminServer.GetItem)                    // 获取指定用户信息
		r.POST("/update/{id}", adminServer.Update)             // 修改指定用户信息
		r.POST("/updatePassword", adminServer.UpdatePassword)  // 修改指定用户密码
		r.POST("/delete/{id}", adminServer.Delete)             // 删除指定用户信息
		r.POST("/updateStatus/{id}", adminServer.UpdateStatus) // 修改帐号状态
		r.POST("/role/update", adminServer.UpdateRole)         // 给用户分配角色
		r.GET("/role/{adminId}", adminServer.GetRoleList)      // 获取指定用户的角色
	}
	{
		r := r.Group("/role")
		roleHandler := handler.NewRoleServer(service.NewRoleService(repo.NewRoleRepo(db)))
		r.POST("/create", roleHandler.Add)                     // 添加角色
		r.POST("/update/{id}", roleHandler.Update)             // 修改角色
		r.POST("/delete", roleHandler.Delete)                  // 批量删除角色
		r.GET("/listAll", roleHandler.ListAll)                 // 获取所有角色
		r.GET("/list", roleHandler.List)                       // 根据角色名称分页获取角色列表
		r.POST("/updateStatus/{id}", roleHandler.UpdateStatus) // 修改角色状态
		r.GET("/listMenu/{roleId}", roleHandler.Update)        // 获取角色相关菜单
		r.GET("/listResource/{roleId}", roleHandler.Update)    // 获取角色相关资源
		r.POST("/allocMenu", roleHandler.Update)               // 给角色分配菜单
		r.POST("/allocResource", roleHandler.Update)           // 给角色分配资源
	}
	{
		r := r.Group("/menu")
		menuHandler := handler.NewMenuServer(service.NewMenuService(repo.NewMenuRepo(db)))
		r.GET("/list/{parentId}", menuHandler.List)
		r.GET("/treeList", menuHandler.TreeList)
		r.GET("/{id}", menuHandler.Details)
		r.POST("/create", menuHandler.Add)
		r.POST("/update/{id}", menuHandler.Update)
		r.POST("/updateHidden/{id}", menuHandler.UpdateHidden)
		r.POST("/delete/{id}", menuHandler.Delete)
	}
}
