package user

import (
	"LibSystem/global"
	"LibSystem/internal/api/controller"
	"LibSystem/internal/repository/dao"
	"LibSystem/internal/service"
	"github.com/gin-gonic/gin"
)

type UPaperRouter struct {
}

func (pr *UPaperRouter) InitApiRouter(router *gin.RouterGroup) {
	paperCtl := controller.NewPaperController(service.NewPaperService(dao.NewPaperDao(global.DB)))
	paperRouter := router.Group("/papers")
	{
		paperRouter.GET("", paperCtl.GetPaperList)
		paperRouter.GET("/:id", paperCtl.GetPaperById)
		paperRouter.GET("/title", paperCtl.GetPaperByTitle)
		paperRouter.POST("/upload", paperCtl.AddPaper)
		paperRouter.GET("/download/:id", paperCtl.DownloadPaper)
	}
}
