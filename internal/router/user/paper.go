package user

import (
	"LibSystem/global"
	"LibSystem/internal/api/controller"
	"LibSystem/internal/repository/dao"
	"LibSystem/internal/service"
	"LibSystem/middle"

	"github.com/gin-gonic/gin"
)

type UPaperRouter struct {
}

func (pr *UPaperRouter) InitApiRouter(router *gin.RouterGroup) {
	paperCtl := controller.NewPaperController(
		service.NewPaperService(dao.NewPaperDao(global.DB)),
		service.NewInfoService(dao.NewInfoDao(global.DB), dao.NewBookDao(global.DB), dao.NewUserDao(global.DB), dao.NewPaperDao(global.DB), dao.NewBorrowDao(global.DB)),
	)
	publicPaperRouter := router.Group("/papers")
	privatePaperRouter := router.Group("/papers")
	privatePaperRouter.Use(middle.VerifyJWT())
	{
		publicPaperRouter.GET("", paperCtl.GetPaperList)
		privatePaperRouter.GET("/:id", paperCtl.GetPaperById)
		privatePaperRouter.GET("/title", paperCtl.GetPaperByTitle)
		privatePaperRouter.POST("/upload", paperCtl.AddPaper)
		privatePaperRouter.GET("/download/:id", paperCtl.DownloadPaper)
	}
}
