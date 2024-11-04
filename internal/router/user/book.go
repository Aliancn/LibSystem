package user

import (
	"LibSystem/global"
	"LibSystem/internal/api/controller"
	"LibSystem/internal/repository/dao"
	"LibSystem/internal/service"
	"github.com/gin-gonic/gin"
)

type UBookRouter struct {
}

func (br *UBookRouter) InitApiRouter(router *gin.RouterGroup) {
	bookCtl := controller.NewBookController(service.NewBookService(dao.NewBookDao(global.DB)))
	bookRouter := router.Group("/books")
	{
		bookRouter.GET("", bookCtl.GetBookList)
		bookRouter.GET("/:id", bookCtl.GetBookById)
		bookRouter.GET("/title", bookCtl.GetBookByTitle)
	}
}
