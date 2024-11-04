package admin

import (
	"LibSystem/global"
	"LibSystem/internal/api/controller"
	"LibSystem/internal/repository/dao"
	"LibSystem/internal/service"
	"github.com/gin-gonic/gin"
)

type PaperRouter struct {
}

func (pr *PaperRouter) InitApiRouter(router *gin.RouterGroup) {
	paperCtl := controller.NewPaperController(service.NewPaperService(dao.NewPaperDao(global.DB)))
	paperRouter := router.Group("/papers")
	{
		paperRouter.DELETE("/:id", paperCtl.DeletePaper)
		paperRouter.PUT("", paperCtl.UpdatePaper)
	}
}
