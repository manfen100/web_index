// Code generated by goctl. DO NOT EDIT.
package model

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type oauthUserFansModel interface {
	Insert(ctx context.Context, data *OauthUserFans) error
	FindOne(ctx context.Context, id string) (*OauthUserFans, error)
	Update(ctx context.Context, data *OauthUserFans) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, id string) (int64, error)
}

type defaultOauthUserFansModel struct {
	conn *mon.Model
}

func newDefaultOauthUserFansModel(conn *mon.Model) *defaultOauthUserFansModel {
	return &defaultOauthUserFansModel{conn: conn}
}

func (m *defaultOauthUserFansModel) Insert(ctx context.Context, data *OauthUserFans) error {
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.CreateAt = time.Now()
		data.UpdateAt = time.Now()
	}

	_, err := m.conn.InsertOne(ctx, data)
	return err
}

func (m *defaultOauthUserFansModel) FindOne(ctx context.Context, id string) (*OauthUserFans, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data OauthUserFans

	err = m.conn.FindOne(ctx, &data, bson.M{"_id": oid})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultOauthUserFansModel) Update(ctx context.Context, data *OauthUserFans) (*mongo.UpdateResult, error) {
	data.UpdateAt = time.Now()

	res, err := m.conn.UpdateOne(ctx, bson.M{"_id": data.ID}, bson.M{"$set": data})
	return res, err
}

func (m *defaultOauthUserFansModel) Delete(ctx context.Context, id string) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, ErrInvalidObjectId
	}

	res, err := m.conn.DeleteOne(ctx, bson.M{"_id": oid})
	return res, err
}
