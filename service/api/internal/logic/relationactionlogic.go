package logic

import (
	"context"

	"mini-douyin/service/api/internal/svc"
	"mini-douyin/service/api/internal/types"
	"mini-douyin/service/social/rpc/followclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationActionLogic {
	return &RelationActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationActionLogic) RelationAction(req *types.Douyin_relation_action_request) (resp *types.Douyin_relation_action_response, err error) {
	res, err := l.svcCtx.SocialRpc.FollowAction(l.ctx, &followclient.DouyinRelationActionRequest{
		Token:      req.Token,
		ToUserId:   req.ToUserId,
		ActionType: req.ActionType,
	})

	if err != nil {
		return nil, err
	}

	return &types.Douyin_relation_action_response{
		StatusCode: res.StatusCode,
		StatusMsg:  res.StatusMsg,
	}, nil
}
