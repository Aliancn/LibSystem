package user

import (
	"LibSystem/global"
	"LibSystem/internal/api/controller"
	"LibSystem/internal/repository/dao"
	"LibSystem/internal/service"
	"LibSystem/middle"

	"github.com/gin-gonic/gin"
)

type UBookRouter struct {
}

func (br *UBookRouter) InitApiRouter(router *gin.RouterGroup) {
	bookCtl := controller.NewBookController(service.NewBookService(dao.NewBookDao(global.DB)))
	publicBookRouter := router.Group("/books")
	privateBookRouter := router.Group("/books")
	privateBookRouter.Use(middle.VerifyJWT())
	{
		publicBookRouter.GET("", bookCtl.GetBookList)
		privateBookRouter.GET("/:id", bookCtl.GetBookById)
		privateBookRouter.GET("/title", bookCtl.GetBookByTitle)
	}
}
