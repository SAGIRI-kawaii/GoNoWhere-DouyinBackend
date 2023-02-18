package logic

import (
	"context"
	"mini-douyin/service/video/rpc/internal/svc"
	"mini-douyin/service/video/rpc/video"

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
	//videos, err := l.svcCtx
	//db.MGetVideos(s.ctx, LIMIT, req.LatestTime)
	//if err != nil {
	//	return vis, nextTime, err
	//}
	//
	//if len(videos) == 0 {
	//	nextTime = time.Now().UnixMilli()
	//	return vis, nextTime, nil
	//} else {
	//	nextTime = videos[len(videos)-1].UpdatedAt.UnixMilli()
	//}
	//
	//if vis, err = pack.Videos(s.ctx, videos, &fromID); err != nil {
	//	nextTime = time.Now().UnixMilli()
	//	return vis, nextTime, err
	//}
	//
	//return vis, nextTime, nil
	return &video.DouyinFeedResponse{}, nil
}
