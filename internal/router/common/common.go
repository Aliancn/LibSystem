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
	infoCtl := controller.NewInfoController(service.NewInfoService(dao.NewInfoDao(global.DB), dao.NewBookDao(global.DB), dao.NewUserDao(global.DB), dao.NewPaperDao(global.DB), dao.NewBorrowDao(global.DB)))
	publicRouter := router.Group("")
	{
		publicRouter.POST("/register", userCtl.Register)
		publicRouter.POST("/login", userCtl.Login)
		publicRouter.POST("/logout", userCtl.Logout)
		publicRouter.GET("/info", infoCtl.GetInfo)
	}
}
