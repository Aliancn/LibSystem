package initialize

import (
	"LibSystem/global"
	"LibSystem/internal/model"
)

func InitTable() {
	// 初始化表
	// 例如:
	global.DB.AutoMigrate(&model.User{})
	global.DB.AutoMigrate(&model.Book{})
	global.DB.AutoMigrate(&model.Paper{})
	global.DB.AutoMigrate(&model.Borrow{})
}
