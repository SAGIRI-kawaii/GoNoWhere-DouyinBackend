package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"mini-douyin/service/api/internal/svc"
	"mini-douyin/service/api/internal/types"
	"mini-douyin/service/video/rpc/videoservice"
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

func (l *PublishActionLogic) PublishAction(req *types.Douyin_publish_action_request) (resp *types.Douyin_publish_action_response, err error) {
	res, err := l.svcCtx.VideoRpc.PublishAction(l.ctx, &videoservice.DouyinPublishActionRequest{
		Token: req.Token,
		Data:  req.Data,
		Title: req.Title,
	})
	if err != nil {
		return nil, err
	}
	return &types.Douyin_publish_action_response{
		StatusCode: res.StatusCode,
		StatusMsg:  *res.StatusMsg,
	}, nil
}
