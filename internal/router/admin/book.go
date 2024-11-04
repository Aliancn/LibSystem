package admin

import (
	"LibSystem/global"
	"LibSystem/internal/api/controller"
	"LibSystem/internal/repository/dao"
	"LibSystem/internal/service"
	"github.com/gin-gonic/gin"
)

type BookRouter struct {
}

func (br *BookRouter) InitApiRouter(router *gin.RouterGroup) {
	bookCtl := controller.NewBookController(service.NewBookService(dao.NewBookDao(global.DB)))
	bookRouter := router.Group("/books")
	{
		bookRouter.POST("", bookCtl.AddBook)
		bookRouter.DELETE("/:id", bookCtl.DeleteBook)
		bookRouter.PUT("", bookCtl.UpdateBook)
	}
}
