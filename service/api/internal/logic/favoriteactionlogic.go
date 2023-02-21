package logic

import (
	"context"

	"mini-douyin/service/api/internal/svc"
	"mini-douyin/service/api/internal/types"
	"mini-douyin/service/interact/rpc/interact"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteActionLogic) FavoriteAction(req *types.Douyin_favorite_action_request) (resp *types.Douyin_favorite_action_response, err error) {

	res, err := l.svcCtx.InteractRpc.FavoriteAction(l.ctx, &interact.DouyinFavoriteActionRequest{
		Token:      req.Token,
		VideoId:    req.VideoId,
		ActionType: req.ActionType,
	})
	if err != nil {
		return nil, err
	}

	return &types.Douyin_favorite_action_response{
		StatusCode: 0,
		StatusMsg:  *res.StatusMsg,
	}, nil

}
