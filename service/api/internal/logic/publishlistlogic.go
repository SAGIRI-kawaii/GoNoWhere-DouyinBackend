package logic

import (
	"context"

	"mini-douyin/service/api/internal/svc"
	"mini-douyin/service/api/internal/types"
	"mini-douyin/service/video/rpc/videoservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishListLogic) PublishList(req *types.Douyin_publish_list_request) (resp *types.Douyin_publish_list_response, err error) {
	res, err := l.svcCtx.VideoRpc.PublishList(l.ctx, &videoservice.DouyinPublishListRequest{
		Token:  req.Token,
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	var VideoList []types.Douyin_video = make([]types.Douyin_video, 0)
	for _, item := range res.VideoList {
		VideoList = append(VideoList, types.Douyin_video{
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

	return &types.Douyin_publish_list_response{
		StatusCode: res.StatusCode,
		StatusMsg:  *res.StatusMsg,
		VideoList:  VideoList,
	}, nil
}
