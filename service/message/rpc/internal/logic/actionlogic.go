package logic

import (
	"context"
	"database/sql"
	"google.golang.org/grpc/status"
	"mini-douyin/common/jwtx"
	"mini-douyin/service/message/model"
	"mini-douyin/service/message/rpc/internal/svc"
	"mini-douyin/service/message/rpc/message"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActionLogic {
	return &ActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActionLogic) Action(in *message.DouyinRelationActionRequest) (*message.DouyinRelationActionResponse, error) {
	// todo: add your logic here and delete this line
	if in.ActionType == 1 {
		res1, err := jwtx.ParseToken2Uid("a", in.Token)
		if err != nil || res1 == 0 {
			return nil, status.Error(100, "Token解析失败")
		}
		newMessage := follows.Messages{
			UserId:   sql.NullInt64{Valid: true, Int64: res1},
			ToUserId: sql.NullInt64{Valid: true, Int64: in.ToUserId},
			Content:  sql.NullString{Valid: true, String: in.Content},
			CreateAt: sql.NullTime{Valid: true, Time: time.Now()},
		}
		_, err = l.svcCtx.MessageModel.Insert(l.ctx, &newMessage)
		if err != nil {
			return nil, status.Error(100, "消息发送失败")
		}
		msg := "sucess"
		return &message.DouyinRelationActionResponse{
			StatusCode: 0,
			StatusMsg:  &msg,
		}, nil
	}

	msg := "error"
	return &message.DouyinRelationActionResponse{
		StatusCode: 1,
		StatusMsg:  &msg,
	}, nil
}
