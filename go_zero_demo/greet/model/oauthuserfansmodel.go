package model

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ OauthUserFansModel = (*customOauthUserFansModel)(nil)

type (
	// OauthUserFansModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOauthUserFansModel.
	OauthUserFansModel interface {
		oauthUserFansModel
	}

	customOauthUserFansModel struct {
		*defaultOauthUserFansModel
	}
)

// NewOauthUserFansModel returns a model for the mongo.
func NewOauthUserFansModel(url, db, collection string) OauthUserFansModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customOauthUserFansModel{
		defaultOauthUserFansModel: newDefaultOauthUserFansModel(conn),
	}
}
