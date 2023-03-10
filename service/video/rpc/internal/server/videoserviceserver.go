// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package server

import (
	"context"

	"mini-douyin/service/video/rpc/internal/logic"
	"mini-douyin/service/video/rpc/internal/svc"
	"mini-douyin/service/video/rpc/video"
)

type VideoServiceServer struct {
	svcCtx *svc.ServiceContext
	video.UnimplementedVideoServiceServer
}

func NewVideoServiceServer(svcCtx *svc.ServiceContext) *VideoServiceServer {
	return &VideoServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *VideoServiceServer) PublishAction(ctx context.Context, in *video.DouyinPublishActionRequest) (*video.DouyinPublishActionResponse, error) {
	l := logic.NewPublishActionLogic(ctx, s.svcCtx)
	return l.PublishAction(in)
}

func (s *VideoServiceServer) PublishList(ctx context.Context, in *video.DouyinPublishListRequest) (*video.DouyinPublishListResponse, error) {
	l := logic.NewPublishListLogic(ctx, s.svcCtx)
	return l.PublishList(in)
}

func (s *VideoServiceServer) Feed(ctx context.Context, in *video.DouyinFeedRequest) (*video.DouyinFeedResponse, error) {
	l := logic.NewFeedLogic(ctx, s.svcCtx)
	return l.Feed(in)
}

func (s *VideoServiceServer) SearchVideo(ctx context.Context, in *video.DouyinSearchRequest) (*video.DouyinSearchResponse, error) {
	l := logic.NewSearchVideoLogic(ctx, s.svcCtx)
	return l.SearchVideo(in)
}
