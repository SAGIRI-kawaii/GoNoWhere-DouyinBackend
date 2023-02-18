package logic

import (
	"context"
	"mini-douyin/common/errno"
	"mini-douyin/common/jwtx"
	"mini-douyin/service/video/rpc/internal/svc"
	"mini-douyin/service/video/rpc/video"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeedLogic) Feed(in *video.DouyinFeedRequest) (*video.DouyinFeedResponse, error) {
	// todo: add your logic here and delete this line
	var vs []*video.Video
	token, err := strconv.ParseInt(*in.Token, 10, 64)
	userid, err := jwtx.ParseToken2Uid("a", uint64(token))
	if err != nil {
		return nil, err
	}
	videos, err := l.svcCtx.VideoModel.GetFeedVideos(l.ctx, 30, in.LatestTime)
	if err != nil {
		return nil, err
	}
	nextTime := time.Now().UnixMilli()
	if len(videos) == 0 {
		return &video.DouyinFeedResponse{
			StatusCode: int32(errno.OK.Code),
			StatusMsg:  &errno.OK.Message,
			VideoList:  vs,
			NextTime:   &nextTime,
		}, nil
	} else {
		nextTime = videos[len(videos)-1].UpdateTime.UnixMilli()
	}
	for _, v := range videos {
		author, err := l.svcCtx.UserModel.FindOneByUserId(l.ctx, v.AuthorId)
		if err != nil {
			return nil, err
		}
		flag, err := l.svcCtx.FollowModel.JudgeFollow(l.ctx, int64(userid), v.AuthorId)
		if err != nil {
			return nil, err
		}
		video_t := &video.Video{
			Id: v.Id,
			Author: &video.User{
				Id:              author.Id,
				Name:            author.Name,
				FollowCount:     &author.FollowCount,
				FollowerCount:   &author.FollowerCount,
				IsFollow:        flag,
				Avatar:          nil,
				BackgroundImage: nil,
				Signature:       nil,
				TotalFavorited:  nil,
				WorkCount:       nil,
				FavoriteCount:   nil,
			},
			PlayUrl:       "",
			CoverUrl:      "",
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    false,
			Title:         "",
		}
	}

	return &video.DouyinFeedResponse{}, nil
}
