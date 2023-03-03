package logic

import (
	"context"
	"errors"
	"go_zero_demo/greet/internal/svc"
	"go_zero_demo/greet/internal/types"
	"go_zero_demo/greet/model"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginReply, err error) {

	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("参数错误")
	}

	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, 1)

	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errors.New("用户不存在")
	default:
		return nil, errors.New("系统错误")

	}

	if userInfo.Password != req.Password {
		return nil, errors.New("密码错误")
	}

	return &types.LoginReply{
		Name: "111",
	}, nil

}
