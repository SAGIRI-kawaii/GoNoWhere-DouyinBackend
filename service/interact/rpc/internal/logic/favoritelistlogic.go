package logic

import (
	"context"

	"mini-douyin/common/jwtx"
	"mini-douyin/service/interact/rpc/interact"
	"mini-douyin/service/interact/rpc/internal/svc"
	"mini-douyin/service/video/rpc/video"

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

func (l *FavoriteListLogic) FavoriteList(in *interact.DouyinFavoriteListRequest) (*interact.DouyinFavoriteListResponse, error) {
	// todo: add your logic here and delete this line

	// 定义最终返回的结果favorite_list,
	var FavoritesList []*interact.Video = make([]*interact.Video, 0)
	// todo：1. 通过 user_id 查询该用户点赞的videos列表
	claims, err := jwtx.ParseToken(in.Token)
	UserId := claims.UserID
	if err != nil {
		return nil, err
	}
	list, err := l.svcCtx.FavoriteModel.FindVideoListByUserId(l.ctx, UserId)
	if err != nil {
		return nil, err
	}
	println("以上OK")

	// todo: 2. 遍历videos：通过video_id 查Video信息补充Video的除Author的字段    ：FindOneByVideoId(ctx context.Context, videoId int64) (*Videos, error)
	for _, item := range list {

		c, err := l.svcCtx.VideoRpc.SearchVideo(l.ctx, &video.DouyinSearchRequest{
			VideoId: item,
		})
		var videoone = c.Video
		if err != nil {
			return nil, err
		}
		var author = &interact.User{
			Id:              videoone.Author.Id,
			Name:            videoone.Author.Name,
			FollowCount:     videoone.Author.FollowCount,
			FollowerCount:   videoone.Author.FollowerCount,
			IsFollow:        videoone.Author.IsFollow, //待查
			Avatar:          videoone.Author.Avatar,
			BackgroundImage: videoone.Author.BackgroundImage,
			Signature:       videoone.Author.Signature,
			TotalFavorited:  videoone.Author.TotalFavorited,
			WorkCount:       videoone.Author.WorkCount,
			FavoriteCount:   videoone.Author.FavoriteCount,
		}
		var Favorites = &interact.Video{
			Id:            videoone.Id, //被点赞的视频的video_id  ，还是自增id？
			Author:        author,
			PlayUrl:       videoone.PlayUrl,
			CoverUrl:      videoone.CoverUrl,
			FavoriteCount: videoone.FavoriteCount,
			CommentCount:  videoone.CommentCount,
			IsFavorite:    true,
			Title:         videoone.Title,
		}
		FavoritesList = append(FavoritesList, Favorites)

		// v, err := l.svcCtx.VideoModel.FindOneByVideoId(l.ctx, item)
		// // videoone, err := l.svcCtx.VideoModel.FindOneByVideoId(l.ctx, item)
		// if err != nil {
		// 	return nil, err
		// }

		// println(v.AuthorId)
		// // todo: 3. 通过video的author_id字段 查User信息补充videoAuthor字段
		// res, err := l.svcCtx.UserModel.FindOneByUserId(l.ctx, videoone.AuthorId)
		// if err != nil {
		// 	return nil, err
		// }

		// var newComment = interact.User{
		// 	Id:              res.UserId,
		// 	Name:            res.Name,
		// 	FollowCount:     &res.FollowCount,
		// 	FollowerCount:   &res.FollowerCount,
		// 	IsFollow:        false,              //待查
		// 	Avatar:          &res.Avatar.String, //
		// 	BackgroundImage: &res.BackgroundImage.String,
		// 	Signature:       &res.Signature.String,
		// 	TotalFavorited:  &res.TotalFavorited.Int64,
		// 	WorkCount:       &res.WorkCount.Int64,
		// 	FavoriteCount:   &res.FavoriteCount.Int64,
		// }
		// println(newComment.Avatar)
		// FavoritesList = append(FavoritesList, &interact.Video{
		// 	Id:            videoone.Id, //被点赞的视频的video_id  ，还是自增id？
		// 	Author:        &newComment,
		// 	PlayUrl:       videoone.PlayUrl,
		// 	CoverUrl:      videoone.CoverUrl,
		// 	FavoriteCount: videoone.FavoriteCount,
		// 	CommentCount:  videoone.CommentCount,
		// 	IsFavorite:    true,
		// 	Title:         videoone.Title,
		// })
	}

	// FavoriteList []*DouyinFavorite
	var a string = "获取该用户点赞的视频列表success"
	return &interact.DouyinFavoriteListResponse{
		StatusCode: int32(0),
		StatusMsg:  &a,
		VideoList:  FavoritesList,
	}, nil
}
