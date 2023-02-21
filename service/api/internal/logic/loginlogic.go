package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"mini-douyin/service/api/internal/svc"
	"mini-douyin/service/api/internal/types"
	"mini-douyin/service/user/rpc/userclient"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.Douyin_user_login_request) (resp *types.Douyin_user_login_response, err error) {

	res, err := l.svcCtx.UserRpc.Login(l.ctx, &userclient.DouyinUserLoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	return &types.Douyin_user_login_response{
		StatusCode: 0,
		StatusMsg:  "success",
		UserID:     res.UserId,
		Token:      res.Token,
	}, nil

}
