package logic

import (
	"context"

	"mini-douyin/service/api/internal/svc"
	"mini-douyin/service/api/internal/types"
	"mini-douyin/service/interact/rpc/interact"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteListLogic) FavoriteList(req *types.Douyin_favorite_list_request) (resp *types.Douyin_favorite_list_response, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.InteractRpc.FavoriteList(l.ctx, &interact.DouyinFavoriteListRequest{
		UserId: int64(req.UserId),
		Token:  req.Token,
	})
	if err != nil {
		return nil, err
	}
	var FavoriteList []types.Douyin_video = make([]types.Douyin_video, 0)
	for _, item := range res.VideoList {
		FavoriteList = append(FavoriteList, types.Douyin_video{
			Id: item.Id,
			Author: types.Douyin_user_info{
				ID:              item.Author.Id,
				Name:            item.Author.Name,
				FollowCount:     *item.Author.FollowCount,
				FollowerCount:   *item.Author.FollowerCount,
				IsFollow:        item.Author.IsFollow,
				Avatar:          *item.Author.Avatar,
				BackgroundImage: *item.Author.BackgroundImage,
				Signature:       *item.Author.Signature,
				TotalFavorited:  *item.Author.TotalFavorited,
				WorkCount:       *item.Author.WorkCount,
				FavoriteCount:   *item.Author.FavoriteCount,
			},
			PlayUrl:       item.PlayUrl,
			CoverUrl:      item.CoverUrl,
			FavoriteCount: item.FavoriteCount,
			CommentCount:  item.CommentCount,
			IsFavorite:    item.IsFavorite,
			Title:         item.Title,
		})
	}
	var a = "获取点赞列表success"
	return &types.Douyin_favorite_list_response{
		StatusCode: res.StatusCode,
		StatusMsg:  a,
		VideoList:  FavoriteList,
	}, nil
}
