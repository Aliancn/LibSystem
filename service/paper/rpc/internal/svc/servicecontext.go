package svc

import (
	"LibSystem/service/paper/model"
	"LibSystem/service/paper/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	PaperModel model.PaperModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("postgres", c.Pgsql.DataSource)
	return &ServiceContext{
		Config:     c,
		PaperModel: model.NewPaperModel(conn, c.CacheRedis),
	}
}
