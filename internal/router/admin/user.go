package admin

import (
	"LibSystem/global"
	"LibSystem/internal/api/controller"
	"LibSystem/internal/repository/dao"
	"LibSystem/internal/service"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur *UserRouter) InitApiRouter(router *gin.RouterGroup) {
	userrouter := router.Group("user")
	userCtl := controller.NewUserController(service.NewUserService(
		dao.NewUserDao(global.DB),
	))
	{
		userrouter.GET("/:id", userCtl.GetById)
		userrouter.GET("/username", userCtl.GetByUsername)
		userrouter.GET("", userCtl.GetList)
		userrouter.POST("", userCtl.AddUser)
		userrouter.PUT("/editPassword", userCtl.EditPassword)
		userrouter.PUT("", userCtl.UpdateUser)
		userrouter.DELETE("", userCtl.DeleteUser)
	}
}
