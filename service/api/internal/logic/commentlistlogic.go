package logic

import (
	"context"

	"mini-douyin/service/api/internal/svc"
	"mini-douyin/service/api/internal/types"
	"mini-douyin/service/interact/rpc/interact"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentListLogic) CommentList(req *types.Douyin_comment_list_request) (resp *types.Douyin_comment_list_response, err error) {
	// todo: add your logic here and delete this line

	res, err := l.svcCtx.InteractRpc.CommentList(l.ctx, &interact.DouyinCommentListRequest{
		Token:   req.Token,
		VideoId: int64(req.VideoId),
	})
	if err != nil {
		return nil, err
	}

	var CommentsList []types.Douyin_comment = make([]types.Douyin_comment, 0)

	for _, item := range res.CommentList {
		CommentsList = append(CommentsList, types.Douyin_comment{
			Id: item.Id,
			User: types.Douyin_user_info{
				ID:              item.User.Id,
				Name:            item.User.Name,
				FollowCount:     *item.User.FollowCount,
				FollowerCount:   *item.User.FollowerCount,
				IsFollow:        item.User.IsFollow,
				Avatar:          *item.User.Avatar,
				BackgroundImage: *item.User.BackgroundImage,
				Signature:       *item.User.Signature,
				TotalFavorited:  *item.User.TotalFavorited,
				WorkCount:       *item.User.WorkCount,
				FavoriteCount:   *item.User.FavoriteCount,
			},
			Content:    item.Content,
			CreateDate: item.CreateDate,
		})
	}

	var a = "获取评论列表success"
	return &types.Douyin_comment_list_response{
		StatusCode:  int(res.StatusCode),
		StatusMsg:   a,
		CommentList: CommentsList,
	}, nil
}
