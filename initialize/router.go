package initialize

import (
	"LibSystem/internal/router"
	"LibSystem/middle"
	"github.com/gin-gonic/gin"
)

func routerInit() *gin.Engine {
	r := gin.Default()
	allRouter := router.AllRouter
	// 注册路由
	common := r.Group("")
	{
		allRouter.CommonRouter.InitApiRouter(common)
	}
	user := r.Group("")
	user.Use(middle.VerifyJWTUser())
	{
		allRouter.UPaperRouter.InitApiRouter(user)
		allRouter.UBookRouter.InitApiRouter(user)
		allRouter.UBorrowRouter.InitApiRouter(user)
	}
	admin := r.Group("/admin")
	admin.Use(middle.VerifyJWTAdmin())
	{
		allRouter.UserRouter.InitApiRouter(admin)
		allRouter.BookRouter.InitApiRouter(admin)
		allRouter.PaperRouter.InitApiRouter(admin)
		allRouter.BorrowRouter.InitApiRouter(admin)
	}
	return r
}
