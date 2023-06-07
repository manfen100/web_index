package logic

import (
	"context"

	"openai-project/opneai/internal/svc"
	"openai-project/opneai/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OpneaiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOpneaiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpneaiLogic {
	return &OpneaiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OpneaiLogic) Opneai(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
