package initialize

import (
	"LibSystem/global"
	"errors"
	"time"

	"golang.org/x/time/rate"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	GormToManyRequestError = errors.New("gorm: to many request")
)

func InitDatabase(dsn string) *gorm.DB {
	var ormLogger logger.Interface
	if global.Config.Server.Level == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,  // DSN data source name
		PreferSimpleProtocol: true, // 禁用复杂协议
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  //设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) //打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	// 慢日志中间件
	SlowQueryLog(db)
	// 限流器中间件
	GormRateLimiter(db, rate.NewLimiter(500, 1000))

	return db
}

// SlowQueryLog 慢查询日志
func SlowQueryLog(db *gorm.DB) {
	err := db.Callback().Query().Before("*").Register("slow_query_start", func(d *gorm.DB) {
		now := time.Now()
		d.Set("start_time", now)
	})
	if err != nil {
		panic(err)
	}

	err = db.Callback().Query().After("*").Register("slow_query_end", func(d *gorm.DB) {
		now := time.Now()
		start, ok := d.Get("start_time")
		if ok {
			duration := now.Sub(start.(time.Time))
			// 一般认为 200 Ms 为Sql慢查询
			if duration > time.Millisecond*200 {
				global.Log.Error("慢查询", "SQL:", d.Statement.SQL.String())
			}
		}
	})
	if err != nil {
		panic(err)
	}
}

// GormRateLimiter Gorm限流器 此限流器不能终止GORM查询链。
func GormRateLimiter(db *gorm.DB, r *rate.Limiter) {
	err := db.Callback().Query().Before("*").Register("RateLimitGormMiddleware", func(d *gorm.DB) {
		if !r.Allow() {
			d.AddError(GormToManyRequestError)
			global.Log.Error(GormToManyRequestError.Error())
			return
		}
	})
	if err != nil {
		panic(err)
	}
}
