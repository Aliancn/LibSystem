package commonpath

import (
	"LibSystem/global"
	"LibSystem/internal/api/controller"
	"LibSystem/internal/repository/dao"
	"LibSystem/internal/service"
	"github.com/gin-gonic/gin"
)

type CommonRouter struct{}

func (cr *CommonRouter) InitApiRouter(router *gin.RouterGroup) {

	userCtl := controller.NewUserController(service.NewUserService(dao.NewUserDao(global.DB)))
	publicRouter := router.Group("")
	{
		publicRouter.POST("/register", userCtl.Register)
		publicRouter.POST("/login", userCtl.Login)
		publicRouter.POST("/logout", userCtl.Logout)
	}
}
