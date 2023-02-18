package logic

import (
	"context"
	"mini-douyin/service/video/api/internal/svc"
	"mini-douyin/service/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishListLogic) PublishList(req *types.Douyin_publish_list_request) (resp *types.Douyin_publish_list_response, err error) {
	// todo: add your logic here and delete this line

	return
}
