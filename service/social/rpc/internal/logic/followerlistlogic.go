package logic

import (
	"context"

	"mini-douyin/service/social/rpc/follow"
	"mini-douyin/service/social/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowerListLogic {
	return &FollowerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowerListLogic) FollowerList(in *follow.DouyinRelationFollowerListRequest) (*follow.DouyinRelationFollowerListResponse, error) {
	// todo: add your logic here and delete this line

	return &follow.DouyinRelationFollowerListResponse{}, nil
}
