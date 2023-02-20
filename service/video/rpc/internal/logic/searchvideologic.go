package logic

import (
	"context"
	"mini-douyin/common/errno"
	"mini-douyin/service/video/rpc/internal/svc"
	"mini-douyin/service/video/rpc/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchVideoLogic {
	return &SearchVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchVideoLogic) SearchVideo(in *video.DouyinSearchRequest) (*video.DouyinSearchResponse, error) {
	// todo: add your logic here and delete this line
	v, err := l.svcCtx.VideoModel.FindOneByVideoId(l.ctx, in.VideoId)
	if err != nil {
		return nil, err
	}
	au, err := l.svcCtx.UserModel.FindOneByUserId(l.ctx, v.AuthorId)
	if err != nil {
		return nil, err
	}

	author := &video.User{
		Id:              au.UserId,
		Name:            au.Name,
		FollowCount:     &au.FollowCount,
		FollowerCount:   &au.FollowerCount,
		IsFollow:        false,
		Avatar:          &au.Avatar.String,
		BackgroundImage: &au.BackgroundImage.String,
		Signature:       &au.Signature.String,
		TotalFavorited:  &au.TotalFavorited.Int64, //获赞数量
		WorkCount:       &au.WorkCount.Int64,
		FavoriteCount:   &au.FavoriteCount.Int64, //点赞数量
	}
	video_t := &video.Video{
		Id:            v.VideoId,
		Author:        author,
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    false,
		Title:         v.Title,
	}
	return &video.DouyinSearchResponse{
		StatusCode: int32(errno.OK.Code),
		StatusMsg:  &errno.OK.Message,
		Video:      video_t,
	}, nil
}
