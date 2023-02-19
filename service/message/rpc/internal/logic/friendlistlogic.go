package logic

import (
	"context"

	"mini-douyin/service/message/rpc/internal/svc"
	"mini-douyin/service/message/rpc/message"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendListLogic) FriendList(in *message.DouyinRelationFriendListRequest) (*message.DouyinRelationFriendListResponse, error) {
	// todo: add your logic here and delete this line

	return &message.DouyinRelationFriendListResponse{}, nil
}
