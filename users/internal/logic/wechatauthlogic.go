package logic

import (
	"context"

	"users/users/internal/svc"
	"users/users/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type WechatAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWechatAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) WechatAuthLogic {
	return WechatAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WechatAuthLogic) WechatAuth(req types.WechatAuthRequest) (resp *types.WechatAuthResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
