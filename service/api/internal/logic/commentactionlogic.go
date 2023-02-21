package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"mini-douyin/service/api/internal/svc"
	"mini-douyin/service/api/internal/types"
	"mini-douyin/service/interact/rpc/interact"
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
	res, err := l.svcCtx.InteractRpc.CommentAction(l.ctx, &interact.DouyinCommentActionRequest{
		Token:       req.Token,
		VideoId:     req.VideoId,
		ActionType:  req.ActionType,
		CommentText: &req.CommentText,
		CommentId:   &req.CommentId,
	})
	if err != nil {
		return nil, err
	}
	if req.ActionType == 1 {
		return &types.Douyin_comment_action_response{
			StatusCode: res.StatusCode,
			StatusMsg:  *res.StatusMsg,
			Comment:    *res.Comment,
		}, nil
	}

	if req.ActionType == 2 {
		return &types.Douyin_comment_action_response{
			StatusCode: res.StatusCode,
			StatusMsg:  *res.StatusMsg,
			Comment:    *res.Comment,
		}, nil
	}

	var StatusMsg = "ActionType错误"
	return &types.Douyin_comment_action_response{
		StatusCode: res.StatusCode,
		StatusMsg:  StatusMsg,
	}, nil
}
