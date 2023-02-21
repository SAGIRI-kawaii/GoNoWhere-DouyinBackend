package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"mini-douyin/common/jwtx"
	"mini-douyin/service/message/rpc/internal/svc"
	"mini-douyin/service/message/rpc/message"
	"time"
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
	res1, err := jwtx.ParseToken(in.Token)
	if err != nil {
		return nil, status.Error(100, "Token解析失败")
	}
	println(time.Unix(in.PreMsgTime, 0).Format("2006-01-02 03:04:05"))
	tm := time.Unix(in.PreMsgTime, 0)
	messages, err := l.svcCtx.MessageModel.FindLatestMsg(l.ctx, res1.UserID, tm.Format("2006-01-02 03:04:05"))
	if err != nil {
		return nil, status.Error(100, "消息列表获取失败")
	}
	msg := "sucess"
	var ansMsg []*message.Message
	for i := 0; i < len(*messages); i++ {
		time := (*messages)[i].CreateAt.Time.String()
		ansMsg = append(ansMsg, &message.Message{
			Id:         (*messages)[i].Id,
			ToUserId:   (*messages)[i].UserId.Int64,
			FromUserId: (*messages)[i].ToUserId.Int64,
			Content:    (*messages)[i].Content.String,
			CreateTime: &time,
		})
	}
	return &message.DouyinMessageChatResponse{
		StatusCode:  0,
		StatusMsg:   &msg,
		MessageList: ansMsg,
	}, nil
}
