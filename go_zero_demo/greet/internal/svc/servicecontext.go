package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go_zero_demo/greet/internal/config"
	"go_zero_demo/greet/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	conn := sqlx.NewMysql(c.Mysql.DataSource)
	println()
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
	}
}
