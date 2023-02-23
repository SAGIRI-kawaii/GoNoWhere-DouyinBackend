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
	resComment := types.Douyin_comment{
		Id: res.Comment.Id,
		User: types.Douyin_user_info{
			ID:              res.Comment.User.Id,
			Name:            res.Comment.User.Name,
			FollowCount:     *res.Comment.User.FollowCount,
			FollowerCount:   *res.Comment.User.FollowerCount,
			IsFollow:        res.Comment.User.IsFollow,
			Avatar:          *res.Comment.User.Avatar,
			BackgroundImage: *res.Comment.User.BackgroundImage,
			Signature:       *res.Comment.User.Signature,
			TotalFavorited:  *res.Comment.User.TotalFavorited,
			WorkCount:       *res.Comment.User.WorkCount,
			FavoriteCount:   *res.Comment.User.FavoriteCount,
		},
		Content:    res.Comment.Content,
		CreateDate: res.Comment.CreateDate,
	}
	if req.ActionType == 1 {
		return &types.Douyin_comment_action_response{
			StatusCode: res.StatusCode,
			StatusMsg:  *res.StatusMsg,
			Comment:    resComment,
		}, nil
	}

	if req.ActionType == 2 {
		return &types.Douyin_comment_action_response{
			StatusCode: res.StatusCode,
			StatusMsg:  *res.StatusMsg,
			Comment:    resComment,
		}, nil
	}

	var StatusMsg = "ActionType错误"
	return &types.Douyin_comment_action_response{
		StatusCode: res.StatusCode,
		StatusMsg:  StatusMsg,
	}, nil
}
