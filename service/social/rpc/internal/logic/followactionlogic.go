package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"mini-douyin/common/jwtx"
	"mini-douyin/service/social/model/follows"
	"mini-douyin/service/social/rpc/follow"
	"mini-douyin/service/social/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
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
		// todo: 改变User里面的关注数
		err = l.svcCtx.UserModel.AddFollowByUserId(l.ctx, int64(userid))
		err = l.svcCtx.UserModel.AddFollowerByUserId(l.ctx, in.ToUserId)
		if err != nil {
			return nil, status.Error(100, "数据库操作出错")
		}
		return &follow.DouyinRelationActionResponse{}, nil

	} else {
		//在follow表中删除一个记录
		userid, err := jwtx.ParseToken2Uid("a", in.Token)
		if err != nil {
			return nil, err
		}
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
