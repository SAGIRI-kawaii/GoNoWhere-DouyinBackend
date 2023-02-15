package logic

import (
	"context"

	"mini-douyin/service/video/api/internal/svc"
	"mini-douyin/service/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishActionLogic {
	return &PublishActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishActionLogic) PublishAction(req *types.Duyin_publish_action_request) (resp *types.Duyin_publish_action_response, err error) {
	// todo: add your logic here and delete this line

	return
}
