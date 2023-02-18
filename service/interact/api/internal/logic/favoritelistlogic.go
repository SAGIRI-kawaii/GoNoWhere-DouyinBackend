package logic

import (
	"context"

	"mini-douyin/service/interact/api/internal/svc"
	"mini-douyin/service/interact/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteListLogic) FavoriteList(req *types.Douyin_favorite_list_request) (resp *types.Douyin_favorite_list_response, err error) {
	// todo: add your logic here and delete this line

	return
}
