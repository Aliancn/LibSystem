package admin

import (
	"LibSystem/global"
	"LibSystem/internal/api/controller"
	"LibSystem/internal/repository/dao"
	"LibSystem/internal/service"
	"github.com/gin-gonic/gin"
)

type BorrowRouter struct {
}

func (br *BorrowRouter) InitApiRouter(router *gin.RouterGroup) {
	borrowCtl := controller.NewBorrowController(service.NewBorrowService(dao.NewBorrowDao(global.DB), dao.NewBookDao(global.DB)))
	borrowRouter := router.Group("/borrows")
	{
		borrowRouter.GET("", borrowCtl.GetAll)
		borrowRouter.DELETE("/:id", borrowCtl.Delete)
	}
}
