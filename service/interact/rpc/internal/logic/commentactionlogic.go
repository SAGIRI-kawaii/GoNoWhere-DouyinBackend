package logic

import (
	"context"

	"mini-douyin/common/jwtx"
	model "mini-douyin/service/interact/model/comments"
	"mini-douyin/service/interact/rpc/interact"
	"mini-douyin/service/interact/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type CommentActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentActionLogic) CommentAction(in *interact.DouyinCommentActionRequest) (*interact.DouyinCommentActionResponse, error) {
	// todo: add your logic here and delete this line

	// UserId := in.Token
	// String转int64
	claims, err := jwtx.ParseToken(in.Token)
	UserId := claims.UserID
	if err != nil {
		return nil, err
	}

	// action_type：1-发布评论，2-删除评论
	// comment_text：用户填写的评论内容，在action_type=1的时候使用
	// comment_id：要删除的评论id，在action_type=2的时候使用
	ActionType := in.ActionType
	if ActionType == 1 {
		newComment := model.Comments{
			UserId:  UserId,
			VideoId: in.VideoId,
			Content: *in.CommentText,
		}
		l.svcCtx.CommentModel.Insert(l.ctx, &newComment)
		err = l.svcCtx.VideoModel.AddCommentByVideoId(l.ctx, in.VideoId)
		if err != nil {
			return nil, status.Error(100, "数据库操作出错")
		}

		/**
		返回Content类型
		*/
		newCommentReturn := interact.Comment{
			Content: *in.CommentText,
		}
		var a string = "发布评论成功"
		return &interact.DouyinCommentActionResponse{

			StatusCode: int32(0),
			StatusMsg:  &a,
			Comment:    &newCommentReturn,
		}, nil

	} else if ActionType == 2 {
		CommentId := *in.CommentId
		err := l.svcCtx.CommentModel.Delete(l.ctx, CommentId)
		if err != nil {
			return nil, status.Error(100, "error")
		}

		err = l.svcCtx.VideoModel.ReduceCommentByVideoId(l.ctx, in.VideoId)
		if err != nil {
			return nil, status.Error(100, "数据库操作出错")
		}

		var a string = "删除评论成功"
		return &interact.DouyinCommentActionResponse{
			StatusCode: int32(0),
			StatusMsg:  &a,
		}, nil
	}
	var a string = "ActionType输入异常"
	return &interact.DouyinCommentActionResponse{
		StatusCode: int32(0),
		StatusMsg:  &a,
	}, nil

}
