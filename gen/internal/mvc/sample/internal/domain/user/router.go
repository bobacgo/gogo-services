package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/gen/internal/mvc/sample/internal/domain/user/handler"
	"github.com/gogoclouds/gogo-services/gen/internal/mvc/sample/internal/domain/user/repo"
	"github.com/gogoclouds/gogo-services/gen/internal/mvc/sample/internal/domain/user/service"
)

func Register(r gin.IRouter) {
	userGp := r.Group("/user")
	{
		userHandler := handler.NewUserServer(service.NewUserService(repo.NewUserRepo(nil)))
		userGp.POST("/list", userHandler.List)
		userGp.POST("/details", userHandler.Details)
		userGp.POST("/add", userHandler.Add)
		userGp.POST("/update", userHandler.Update)
		userGp.POST("/delete", userHandler.Delete)
	}
}
