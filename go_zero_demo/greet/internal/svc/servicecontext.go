package svc

import (
	"go_zero_demo/greet/internal/config"
	"go_zero_demo/greet/model"
)

type ServiceContext struct {
	Config config.Config
	//UserModel model.UserModel
	OauthUserFansModel model.OauthUserFansModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	//conn := sqlx.N(c.Mongo.Datasource)
	//conn := sqlx.NewMysql(c.Mysql.DataSource)
	//var clientManager = syncx.()

	println()
	return &ServiceContext{
		Config: c,
		//UserModel: model.NewUserModel(conn, c.CacheRedis),
		OauthUserFansModel: model.NewOauthUserFansModel(c.Mongo.Datasource, "oauth", "oauth_user_fans"),
	}
}
