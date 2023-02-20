package logic

import (
	"context"

	"mini-douyin/service/video/api/internal/svc"
	"mini-douyin/service/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedLogic) Feed(req *types.Douyin_feed_request) (resp *types.Douyin_feed_response, err error) {
	// todo: add your logic here and delete this line

	return
}
