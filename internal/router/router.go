package router

import (
	"LibSystem/internal/router/admin"
	commonpath "LibSystem/internal/router/common"
	"LibSystem/internal/router/user"
)

type RouterGroup struct {
	admin.UserRouter
	commonpath.CommonRouter
	admin.BookRouter
	admin.PaperRouter
	admin.BorrowRouter
	user.UBookRouter
	user.UPaperRouter
	user.UBorrowRouter
}

var AllRouter = new(RouterGroup)
