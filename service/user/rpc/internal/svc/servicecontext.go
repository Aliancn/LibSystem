package svc

import (
	"LibSystem/service/user/model"
	"LibSystem/service/user/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	_ "github.com/lib/pq"
)

type ServiceContext struct {
	Config config.Config

	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("postgres",c.Pgsql.DataSource)
	return &ServiceContext{
		Config: c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
	}
}
