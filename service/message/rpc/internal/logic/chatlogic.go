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

	println("TIMETIME:")
	println(in.PreMsgTime)
	tm := time.UnixMilli(in.PreMsgTime)
	println(tm.String())

	messages, err := l.svcCtx.MessageModel.FindLatestMsg(l.ctx, res1.UserID, tm.String())
	if err != nil {
		return nil, status.Error(100, "消息列表获取失败")
	}
	msg := "sucess"
	var ansMsg []*message.Message
	for i := 0; i < len(*messages); i++ {
		time1 := (*messages)[i].CreateAt.Time.String()

		t, _ := time.ParseInLocation("2006-01-02 15:04:05", time1[:19], time.Local)
		println("TIME:")
		println(time1[:19])
		println(t.String())
		println(t.UnixMilli())
		//t.UnixMilli() to string
		time2 := t.UnixMilli()
		ansMsg = append(ansMsg, &message.Message{
			Id:         (*messages)[i].Id,
			ToUserId:   (*messages)[i].UserId.Int64,
			FromUserId: (*messages)[i].ToUserId.Int64,
			Content:    (*messages)[i].Content.String,
			CreateTime: &time2,
		})
	}
	return &message.DouyinMessageChatResponse{
		StatusCode:  0,
		StatusMsg:   &msg,
		MessageList: ansMsg,
	}, nil
}
