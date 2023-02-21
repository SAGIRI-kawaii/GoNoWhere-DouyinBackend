package logic

import (
	"context"

	"mini-douyin/service/interact/api/internal/svc"
	"mini-douyin/service/interact/api/internal/types"
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
		Token:  req.Token,
		UserId: int64(req.UserId),
	})
	if err != nil {
		return nil, err
	}

	var VideoList []*types.Douyin_video = make([]*types.Douyin_video, 0)

	for _, item := range res.VideoList {

		var author = &types.Douyin_user{
			ID:               int(item.Author.Id),
			Name:             item.Author.Name,
			FollowCount:      int(*item.Author.FollowCount),
			FollowerCount:    int(*item.Author.FollowerCount),
			IsFollow:         item.Author.IsFollow, //待查
			Avatar:           *item.Author.Avatar,
			Background_image: *item.Author.BackgroundImage,
			Signature:        *item.Author.Signature,
			TotalFavorited:   int(*item.Author.TotalFavorited),
			WorkCount:        int(*item.Author.WorkCount),
			FavoriteCount:    int(*item.Author.FavoriteCount),
		}

		VideoList = append(VideoList, &types.Douyin_video{
			Id:            int(item.Id),
			Author:        *author,
			PlayUrl:       item.PlayUrl,
			CoverUrl:      item.CoverUrl,
			FavoriteCount: int(item.FavoriteCount),
			CommentCount:  int(item.CommentCount),
			IsFavorite:    item.IsFavorite,
			Titlestring:   item.Title,
		})
	}

	var a string = "获取点赞列表success"
	return &types.Douyin_favorite_list_response{
		StatusCode: int(res.StatusCode),
		StatusMsg:  a,
		VideoList:  VideoList,
	}, nil
}
