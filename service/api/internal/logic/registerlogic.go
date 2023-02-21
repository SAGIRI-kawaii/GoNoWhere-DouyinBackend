package logic

import (
	"context"

	"mini-douyin/service/api/internal/svc"
	"mini-douyin/service/api/internal/types"
	"mini-douyin/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.Douyin_user_register_request) (resp *types.Douyin_user_register_response, err error) {

	res, err := l.svcCtx.UserRpc.Register(l.ctx, &userclient.DouyinUserRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	return &types.Douyin_user_register_response{
		StatusCode: 0,
		StatusMsg:  "success",
		UserID:     res.UserId,
		Token:      res.Token,
	}, nil
}
