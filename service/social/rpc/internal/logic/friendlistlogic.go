package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"google.golang.org/grpc/status"
	"mini-douyin/common/jwtx"
	"mini-douyin/service/social/rpc/follow"
	"mini-douyin/service/social/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendListLogic) FriendList(in *follow.DouyinRelationFriendListRequest) (*follow.DouyinRelationFriendListResponse, error) {
	// todo: add your logic here and delete this line
	res, err := jwtx.ParseToken(in.Token)
	if err != nil {
		return nil, err
	}
	if in.UserId != res.UserID {
		return nil, status.Error(100, "非法token")
	}
	userid := res.UserID

	var u []*follow.FriendUser
	friends, err := l.svcCtx.FriendModel.FindAllByUserId(l.ctx, userid)

	if friends != nil {
		for _, f := range friends {
			res, err := l.svcCtx.UserModel.FindOneByUserId(l.ctx, f.ToUserId.Int64)
			if err != nil {
				return nil, status.Error(100, "查询用户失败")
			}
			msg, err := l.svcCtx.MessageModel.FindOneLatestMsgByUid(l.ctx, userid, f.ToUserId.Int64)
			var msgcontent string
			switch err {
			case nil:
				msgcontent = msg.Content.String
			case sqlc.ErrNotFound:
				msgcontent = ""
			default:
				return nil, err
			}

			us := &follow.FriendUser{
				Id:              res.Id,
				Name:            res.Name,
				FollowCount:     &res.FollowerCount,
				FollowerCount:   &res.FollowerCount,
				IsFollow:        true,
				Avatar:          &res.Avatar.String,
				BackgroundImage: &res.BackgroundImage.String,
				Signature:       &res.Signature.String,
				TotalFavorited:  &res.TotalFavorited.Int64,
				WorkCount:       &res.WorkCount.Int64,
				FavoriteCount:   &res.FavoriteCount.Int64,
				Message:         msgcontent,
				MsgType:         1,
			}
			u = append(u, us)

		}
	}
	friends, err = l.svcCtx.FriendModel.FindAllByToUserId(l.ctx, userid)

	if friends != nil {
		for _, f := range friends {
			res, err := l.svcCtx.UserModel.FindOneByUserId(l.ctx, f.UserId.Int64)
			if err != nil {
				return nil, status.Error(100, "查询用户失败")
			}
			msg, err := l.svcCtx.MessageModel.FindOneLatestMsgByUid(l.ctx, f.UserId.Int64, userid)
			var msgcontent string
			switch err {
			case nil:
				msgcontent = msg.Content.String
			case sqlc.ErrNotFound:
				msgcontent = ""
			default:
				return nil, err
			}

			us := &follow.FriendUser{
				Id:              res.Id,
				Name:            res.Name,
				FollowCount:     &res.FollowerCount,
				FollowerCount:   &res.FollowerCount,
				IsFollow:        true,
				Avatar:          &res.Avatar.String,
				BackgroundImage: &res.BackgroundImage.String,
				Signature:       &res.Signature.String,
				TotalFavorited:  &res.TotalFavorited.Int64,
				WorkCount:       &res.WorkCount.Int64,
				FavoriteCount:   &res.FavoriteCount.Int64,
				Message:         msgcontent,
				MsgType:         0,
			}
			u = append(u, us)

		}
	}

	return &follow.DouyinRelationFriendListResponse{
		UserList: u,
	}, nil
}
