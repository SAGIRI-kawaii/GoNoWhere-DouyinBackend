package logic

import (
	"context"

	"mini-douyin/service/video/rpc/internal/svc"
	"mini-douyin/service/video/rpc/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishActionLogic {
	return &PublishActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishActionLogic) PublishAction(in *video.DouyinPublishActionRequest) (*video.DouyinPublishActionResponse, error) {
	// todo: add your logic here and delete this line

	return &video.DouyinPublishActionResponse{}, nil
}
