package logic

import (
	"context"
	"mini-douyin/common/errno"
	_ "mini-douyin/common/jwtx"
	_ "strconv"
	_ "time"

	"mini-douyin/service/video/rpc/internal/svc"
	"mini-douyin/service/video/rpc/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishListLogic) PublishList(in *video.DouyinPublishListRequest) (*video.DouyinPublishListResponse, error) {
	// todo: add your logic here and delete this line
	//token, err := strconv.ParseInt(in.Token, 10, 64)
	//if err != nil {
	//	return nil, err
	//}
	//userid, err := jwtx.ParseToken2Uid("a", uint64(token))
	//if err != nil {
	//
	var vs []*video.Video
	videos, err := l.svcCtx.VideoModel.GetVideosByAuthorID(l.ctx, &in.UserId)
	if err != nil {
		return nil, err
	}
	au, err := l.svcCtx.UserModel.FindOneByUserId(l.ctx, in.UserId)
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
	if len(videos) == 0 {
		return &video.DouyinPublishListResponse{
			StatusCode: int32(errno.OK.Code),
			StatusMsg:  &errno.OK.Message,
			VideoList:  nil,
		}, nil
	}
	for _, v := range videos {
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
		vs = append(vs, video_t)
	}
	return &video.DouyinPublishListResponse{
		StatusCode: int32(errno.OK.Code),
		StatusMsg:  &errno.OK.Message,
		VideoList:  vs,
	}, nil
}
