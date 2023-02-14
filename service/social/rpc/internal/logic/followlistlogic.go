package logic

import (
	"context"

	"mini-douyin/service/social/rpc/follow"
	"mini-douyin/service/social/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowListLogic {
	return &FollowListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowListLogic) FollowList(in *follow.DouyinRelationFollowListRequest) (*follow.DouyinRelationFollowListResponse, error) {
	// todo: add your logic here and delete this line

	return &follow.DouyinRelationFollowListResponse{}, nil
}
