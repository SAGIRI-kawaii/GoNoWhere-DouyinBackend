package logic

import (
	"context"

	"mini-douyin/service/video/rpc/internal/svc"
	"mini-douyin/service/video/rpc/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishListLogic) PublishList(in *video.DouyinPublishListRequest) (*video.DouyinPublishListResponse, error) {
	// todo: add your logic here and delete this line

	return &video.DouyinPublishListResponse{}, nil
}
