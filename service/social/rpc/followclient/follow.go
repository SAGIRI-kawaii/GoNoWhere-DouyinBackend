// Code generated by goctl. DO NOT EDIT.
// Source: follow.proto

package followclient

import (
	"context"

	"mini-douyin/service/social/rpc/follow"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	DouyinRelationActionRequest        = follow.DouyinRelationActionRequest
	DouyinRelationActionResponse       = follow.DouyinRelationActionResponse
	DouyinRelationFollowListRequest    = follow.DouyinRelationFollowListRequest
	DouyinRelationFollowListResponse   = follow.DouyinRelationFollowListResponse
	DouyinRelationFollowerListRequest  = follow.DouyinRelationFollowerListRequest
	DouyinRelationFollowerListResponse = follow.DouyinRelationFollowerListResponse
	DouyinRelationFriendListRequest    = follow.DouyinRelationFriendListRequest
	DouyinRelationFriendListResponse   = follow.DouyinRelationFriendListResponse
	FriendUser                         = follow.FriendUser
	User                               = follow.User

	Follow interface {
		FollowAction(ctx context.Context, in *DouyinRelationActionRequest, opts ...grpc.CallOption) (*DouyinRelationActionResponse, error)
		FollowList(ctx context.Context, in *DouyinRelationFollowListRequest, opts ...grpc.CallOption) (*DouyinRelationFollowListResponse, error)
		FollowerList(ctx context.Context, in *DouyinRelationFollowerListRequest, opts ...grpc.CallOption) (*DouyinRelationFollowerListResponse, error)
		FriendList(ctx context.Context, in *DouyinRelationFriendListRequest, opts ...grpc.CallOption) (*DouyinRelationFriendListResponse, error)
	}

	defaultFollow struct {
		cli zrpc.Client
	}
)

func NewFollow(cli zrpc.Client) Follow {
	return &defaultFollow{
		cli: cli,
	}
}

func (m *defaultFollow) FollowAction(ctx context.Context, in *DouyinRelationActionRequest, opts ...grpc.CallOption) (*DouyinRelationActionResponse, error) {
	client := follow.NewFollowClient(m.cli.Conn())
	return client.FollowAction(ctx, in, opts...)
}

func (m *defaultFollow) FollowList(ctx context.Context, in *DouyinRelationFollowListRequest, opts ...grpc.CallOption) (*DouyinRelationFollowListResponse, error) {
	client := follow.NewFollowClient(m.cli.Conn())
	return client.FollowList(ctx, in, opts...)
}

func (m *defaultFollow) FollowerList(ctx context.Context, in *DouyinRelationFollowerListRequest, opts ...grpc.CallOption) (*DouyinRelationFollowerListResponse, error) {
	client := follow.NewFollowClient(m.cli.Conn())
	return client.FollowerList(ctx, in, opts...)
}

func (m *defaultFollow) FriendList(ctx context.Context, in *DouyinRelationFriendListRequest, opts ...grpc.CallOption) (*DouyinRelationFriendListResponse, error) {
	client := follow.NewFollowClient(m.cli.Conn())
	return client.FriendList(ctx, in, opts...)
}
