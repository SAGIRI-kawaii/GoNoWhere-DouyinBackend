package logic

import (
	"context"

	"mini-douyin/service/api/internal/svc"
	"mini-douyin/service/api/internal/types"
	"mini-douyin/service/message/rpc/douyinrelationservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageActionLogic {
	return &MessageActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageActionLogic) MessageAction(req *types.Douyin_message_action_request) (resp *types.Douyin_message_action_response, err error) {

	res, err := l.svcCtx.MessageRpc.Action(l.ctx, &douyinrelationservice.DouyinRelationActionRequest{
		Token:      req.Token,
		ToUserId:   req.ToUserID,
		ActionType: req.ActionType,
		Content:    req.Content,
	})

	if err != nil {
		return nil, err
	}

	return &types.Douyin_message_action_response{
		StatusCode: res.StatusCode,
		StatusMsg:  *res.StatusMsg,
	}, nil
}
