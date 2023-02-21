package logic

import (
	"context"

	"mini-douyin/service/api/internal/svc"
	"mini-douyin/service/api/internal/types"
	"mini-douyin/service/message/rpc/douyinrelationservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageChatLogic {
	return &MessageChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageChatLogic) MessageChat(req *types.Douyin_message_chat_request) (resp *types.Douyin_message_chat_response, err error) {

	res, err := l.svcCtx.MessageRpc.Chat(l.ctx, &douyinrelationservice.DouyinMessageChatRequest{
		Token:      req.Token,
		ToUserId:   req.ToUserID,
		PreMsgTime: req.PreMsgTime,
	})

	if err != nil {
		return nil, err
	}

	MessageList := make([]types.Douyin_message, 0)
	for _, item := range res.MessageList {
		MessageList = append(MessageList, types.Douyin_message{
			ID:         item.Id,
			ToUserID:   item.ToUserId,
			FromUserID: item.FromUserId,
			Content:    item.Content,
			CreateTime: *item.CreateTime,
		})
	}

	return &types.Douyin_message_chat_response{
		StatusCode:  res.StatusCode,
		StatusMsg:   *res.StatusMsg,
		MessageList: MessageList,
	}, nil
}
