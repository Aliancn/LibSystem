package admin

import (
	"LibSystem/global"
	"LibSystem/internal/api/controller"
	"LibSystem/internal/repository/dao"
	"LibSystem/internal/service"
	"LibSystem/middle"

	"github.com/gin-gonic/gin"
)

type PaperRouter struct {
}

func (pr *PaperRouter) InitApiRouter(router *gin.RouterGroup) {
	paperCtl := controller.NewPaperController(service.NewPaperService(dao.NewPaperDao(global.DB)))
	privatePaperRouter := router.Group("/papers")
	privatePaperRouter.Use(middle.VerifyJWT())
	{
		privatePaperRouter.DELETE("/:id", paperCtl.DeletePaper)
		privatePaperRouter.PUT("", paperCtl.UpdatePaper)
	}
}
