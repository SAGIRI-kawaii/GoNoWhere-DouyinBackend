package logic

import (
	"context"

	"mini-douyin/service/interact/rpc/interact"
	"mini-douyin/service/interact/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// {
//     "status_code": "string",
//     "status_msg": "string",
//     "video_list": [
//         {
//             "id": 0,
//             "author": {
//                 "id": 0,
//                 "name": "string",
//                 "follow_count": 0,
//                 "follower_count": 0,
//                 "is_follow": true,
//                 "avatar": "string",
//                 "background_image": "string",
//                 "signature": "string",
//                 "total_favorited": "string",
//                 "work_count": 0,
//                 "favorite_count": 0
//             },
//             "play_url": "string",
//             "cover_url": "string",
//             "favorite_count": 0,
//             "comment_count": 0,
//             "is_favorite": true,
//             "title": "string"
//         }
//     ]
// }

func (l *FavoriteListLogic) FavoriteList(in *interact.DouyinFavoriteListRequest) (*interact.DouyinFavoriteListResponse, error) {
	// todo: add your logic here and delete this line
	// FindList方法查DB返回的list是 Favorite  类型
	list, err := l.svcCtx.FavoriteModel.FindFavoritesListByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	// 定义最终返回的结果favorite_list,
	var FavoritesList []*interact.Video = make([]*interact.Video, 0)
	for _, item := range list {
		// todo: 通过video_id查Video信息补充Video的除Author的字段
		// todo: 通过uid查User信息补充Author字段

		FavoritesList = append(FavoritesList, &interact.Video{
			Id: item.VideoId, //被点赞的视频的video_id  ，还是自增id？
			// Author:     ,
			// PlayUrl:    ,
			// CreateDate: timetostring,
			// CoverUrl,
			// FavoriteCount,
			// CommentCount,
			// IsFavorite    ,
			// Title
		})

	}

	// FavoriteList []*DouyinFavorite
	var a string = "获取该用户点赞的视频列表success"
	return &interact.DouyinFavoriteListResponse{
		StatusCode: int32(0),
		StatusMsg:  &a,
		VideoList:  FavoritesList,
	}, nil
}
