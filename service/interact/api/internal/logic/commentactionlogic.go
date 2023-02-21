package logic

import (
	"context"

	"mini-douyin/service/interact/api/internal/svc"
	"mini-douyin/service/interact/api/internal/types"
	"mini-douyin/service/interact/rpc/interact"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentActionLogic) CommentAction(req *types.Douyin_comment_action_request) (resp *types.Douyin_comment_action_response, err error) {
	// todo: add your logic here and delete this line
	var s = int64(req.CommentId)
	res, err := l.svcCtx.InteractRpc.CommentAction(l.ctx, &interact.DouyinCommentActionRequest{
		Token:       req.Token,
		VideoId:     int64(req.VideoId),
		ActionType:  int32(req.ActionType),
		CommentText: &req.CommentText,
		CommentId:   &s,
	})
	if err != nil {
		return nil, err
	}
	if req.ActionType == 1 {
		return &types.Douyin_comment_action_response{
			StatusCode: int(res.StatusCode),
			StatusMsg:  *res.StatusMsg,
			Comment:    *res.Comment,
		}, nil
	}

	if req.ActionType == 2 {
		return &types.Douyin_comment_action_response{
			StatusCode: int(res.StatusCode),
			StatusMsg:  *res.StatusMsg,
			// Comment:    *res.Comment,
		}, nil
	}

	var StatusMsg = "ActionType错误"
	return &types.Douyin_comment_action_response{
		StatusCode: int(res.StatusCode),
		StatusMsg:  StatusMsg,
	}, nil
}
