package admin

import (
	"LibSystem/global"
	"LibSystem/internal/api/controller"
	"LibSystem/internal/repository/dao"
	"LibSystem/internal/service"
	"LibSystem/middle"

	"github.com/gin-gonic/gin"
)

type BorrowRouter struct {
}

func (br *BorrowRouter) InitApiRouter(router *gin.RouterGroup) {
	borrowCtl := controller.NewBorrowController(service.NewBorrowService(dao.NewBorrowDao(global.DB), dao.NewBookDao(global.DB)))
	privateBorrowRouter := router.Group("/borrows")
	privateBorrowRouter.Use(middle.VerifyJWT())
	{
		privateBorrowRouter.GET("", borrowCtl.GetAll)
		privateBorrowRouter.DELETE("/:id", borrowCtl.Delete)
	}
}
