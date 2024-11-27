package admin

import (
	"LibSystem/global"
	"LibSystem/internal/api/controller"
	"LibSystem/internal/repository/dao"
	"LibSystem/internal/service"
	"LibSystem/middle"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur *UserRouter) InitApiRouter(router *gin.RouterGroup) {
	userCtl := controller.NewUserController(service.NewUserService(
		dao.NewUserDao(global.DB),
	))
	privateUserrouter := router.Group("user")
	privateUserrouter.Use(middle.VerifyJWT())
	{
		privateUserrouter.GET("/:id", userCtl.GetById)
		privateUserrouter.GET("/username", userCtl.GetByUsername)
		privateUserrouter.GET("", userCtl.GetList)
		privateUserrouter.POST("", userCtl.AddUser)
		privateUserrouter.PUT("/editPassword", userCtl.EditPassword)
		privateUserrouter.PUT("", userCtl.UpdateUser)
		privateUserrouter.DELETE("", userCtl.DeleteUser)
	}
}
