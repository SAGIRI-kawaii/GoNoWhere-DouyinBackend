package logic

import (
	"context"
	"mini-douyin/common/errno"
	"mini-douyin/common/jwtx"
	"mini-douyin/service/video/rpc/internal/svc"
	"mini-douyin/service/video/rpc/video"
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
	var vs []*video.Video
	claims, err := jwtx.ParseToken(*in.Token)
	userid := claims.UserID
	if err != nil {
		return nil, err
	}
	videos, err := l.svcCtx.VideoModel.GetFeedVideos(l.ctx, 30, in.LatestTime)
	if err != nil {
		return nil, err
	}
	nextTime := time.Now().Unix()
	if len(videos) == 0 {
		return &video.DouyinFeedResponse{
			StatusCode: int32(errno.ErrQueryVideosFail.Code),
			StatusMsg:  &errno.ErrQueryVideosFail.Message,
			VideoList:  nil,
			NextTime:   &nextTime,
		}, nil
	} else {
		nextTime = videos[len(videos)-1].UpdateTime.Unix()
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
		//logx.Info("videoID:" + strconv.FormatInt(v.VideoId, 10) + "time:" + strconv.FormatInt(v.UpdateTime.Unix(), 10))
		video_t := &video.Video{
			Id: v.Id,
			Author: &video.User{
				Id:              author.Id,
				Name:            author.Name,
				FollowCount:     &author.FollowCount,
				FollowerCount:   &author.FollowerCount,
				IsFollow:        flag,
				Avatar:          &author.Avatar.String,
				BackgroundImage: &author.BackgroundImage.String,
				Signature:       &author.Signature.String,
				TotalFavorited:  &author.TotalFavorited.Int64, //获赞数量
				WorkCount:       &author.WorkCount.Int64,
				FavoriteCount:   &author.FavoriteCount.Int64, //点赞数量
			},
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    false,
			Title:         v.Title,
		}
		vs = append(vs, video_t)
	}
	nextTime = videos[len(videos)-1].UpdateTime.Unix()
	//logx.Info("Next_Time:" + time.Unix(nextTime, 0).String())
	return &video.DouyinFeedResponse{
		StatusCode: int32(errno.OK.Code),
		StatusMsg:  &errno.OK.Message,
		VideoList:  vs,
		NextTime:   &nextTime,
	}, nil
}
