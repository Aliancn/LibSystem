package user

import (
	"LibSystem/global"
	"LibSystem/internal/api/controller"
	"LibSystem/internal/repository/dao"
	"LibSystem/internal/service"
	"LibSystem/middle"

	"github.com/gin-gonic/gin"
)

type UBorrowRouter struct {
}

func (br *UBorrowRouter) InitApiRouter(router *gin.RouterGroup) {
	borrowCtl := controller.NewBorrowController(service.NewBorrowService(dao.NewBorrowDao(global.DB), dao.NewBookDao(global.DB)))
	// borrowRouter := router.Group("/borrows")
	privateBorrowRouter := router.Group("/borrows")
	privateBorrowRouter.Use(middle.VerifyJWT())
	{
		privateBorrowRouter.POST("", borrowCtl.Borrow)
		privateBorrowRouter.PUT("/:id", borrowCtl.Return)
		privateBorrowRouter.GET("", borrowCtl.GetByUser)
	}
}
