package admin

import (
	"LibSystem/global"
	"LibSystem/internal/api/controller"
	"LibSystem/internal/repository/dao"
	"LibSystem/internal/service"
	"LibSystem/middle"

	"github.com/gin-gonic/gin"
)

type BookRouter struct {
}

func (br *BookRouter) InitApiRouter(router *gin.RouterGroup) {
	bookCtl := controller.NewBookController(service.NewBookService(dao.NewBookDao(global.DB)))
	privateBookRouter := router.Group("/books")
	privateBookRouter.Use(middle.VerifyJWT())
	{
		privateBookRouter.POST("", bookCtl.AddBook)
		privateBookRouter.DELETE("/:id", bookCtl.DeleteBook)
		privateBookRouter.PUT("", bookCtl.UpdateBook)
	}
}
