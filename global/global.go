package global

import (
	"LibSystem/config"
	"LibSystem/logger"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	Config *config.AllConfig // 全局Config
	Log    logger.ILog
	DB     *gorm.DB
	Redis  *redis.Client
)
