package logic

import (
	"context"

	"mini-douyin/service/message/rpc/internal/svc"
	"mini-douyin/service/message/rpc/message"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActionLogic {
	return &ActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActionLogic) Action(in *message.DouyinRelationActionRequest) (*message.DouyinRelationActionResponse, error) {
	// todo: add your logic here and delete this line

	return &message.DouyinRelationActionResponse{}, nil
}
