package logic

import (
	"context"

	"mini-douyin/service/message/rpc/internal/svc"
	"mini-douyin/service/message/rpc/message"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatLogic {
	return &ChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChatLogic) Chat(in *message.DouyinMessageChatRequest) (*message.DouyinMessageChatResponse, error) {
	// todo: add your logic here and delete this line

	return &message.DouyinMessageChatResponse{}, nil
}
