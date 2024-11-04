package initialize

import (
	"LibSystem/config"
	"LibSystem/global"
	"LibSystem/logger"
	"github.com/gin-gonic/gin"
)

func GlobalInit() *gin.Engine {
	// 配置文件初始化
	global.Config = config.InitLoadConfig()
	// Log初始化
	global.Log = logger.NewLogger(global.Config.Log.Level, global.Config.Log.FilePath)
	// Gorm初始化
	global.DB = InitDatabase(global.Config.DataSource.Dsn())
	// Redis初始化
	global.Redis = initRedis()
	// Router初始化
	router := routerInit()
	return router
}
