package logic

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"google.golang.org/grpc/status"
	"mini-douyin/common/jwtx"
	"mini-douyin/service/message/model/friends"
	"mini-douyin/service/social/model/follows"
	"mini-douyin/service/social/rpc/follow"
	"mini-douyin/service/social/rpc/internal/svc"
)

type FollowActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowActionLogic {
	return &FollowActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowActionLogic) FollowAction(in *follow.DouyinRelationActionRequest) (*follow.DouyinRelationActionResponse, error) {
	// todo: add your logic here and delete this line
	//根据ActionType判断操作类型
	if in.ActionType == 1 {
		// 在follow表中插入一条记录
		userid, err := jwtx.ParseToken2Uid("a", in.Token)
		if err != nil {
			return nil, err
		}

		newFollow := follows.Follows{
			UserId:   int64(userid),
			ToUserId: in.ToUserId,
		}
		_, err = l.svcCtx.FollowModel.Insert(l.ctx, &newFollow)
		if err != nil {
			return nil, status.Error(100, "关注失败")
		}
		// 如果两人互关，就增加好友关系
		res, err := l.svcCtx.FollowModel.FindAllByUserId(l.ctx, in.ToUserId)
		for _, f := range res {
			if f.ToUserId == userid {
				newFriend := friends.Friends{
					UserId:   sql.NullInt64{userid, true},
					ToUserId: sql.NullInt64{in.ToUserId, true},
					CreateAt: sql.NullTime{},
				}
				_, err := l.svcCtx.FriendModel.Insert(l.ctx, &newFriend)
				if err != nil {
					return nil, status.Error(100, "生成好友失败")
				}
				break
			}
		}
		// todo: 改变User里面的关注数
		err = l.svcCtx.UserModel.AddFollowByUserId(l.ctx, int64(userid))
		err = l.svcCtx.UserModel.AddFollowerByUserId(l.ctx, in.ToUserId)
		if err != nil {
			return nil, status.Error(100, err.Error())
		}

		return &follow.DouyinRelationActionResponse{}, nil

	} else {

		userid, err := jwtx.ParseToken2Uid("a", in.Token)
		if err != nil {
			return nil, err
		}
		//若两人互关，就删除好友关系
		res, err := l.svcCtx.FollowModel.FindAllByUserId(l.ctx, in.ToUserId)
		for _, f := range res {
			if f.ToUserId == userid {

				if _, err := l.svcCtx.FriendModel.FindOneByBothway(l.ctx, userid, in.ToUserId); err == sqlc.ErrNotFound {
					err := l.svcCtx.FriendModel.DeleteById(l.ctx, in.ToUserId, userid)
					if err != nil {
						return nil, status.Error(100, "删除好友失败")
					}
				} else {
					err := l.svcCtx.FriendModel.DeleteById(l.ctx, userid, in.ToUserId)
					if err != nil {
						return nil, status.Error(100, "删除好友失败")
					}

				}

				break
			}
		}
		//在follow表中删除一个记录

		err = l.svcCtx.FollowModel.DeleteByuid(l.ctx, int64(userid), in.ToUserId)
		if err != nil {
			return nil, status.Error(100, "删除失败")
		}

		//todo : 改变User里面的关注数
		err = l.svcCtx.UserModel.ReduceFollowByUserId(l.ctx, int64(userid))
		err = l.svcCtx.UserModel.ReduceFollowerByUserId(l.ctx, in.ToUserId)
		if err != nil {
			return nil, status.Error(100, "数据库操作出错")
		}
		return &follow.DouyinRelationActionResponse{}, nil

	}

}
