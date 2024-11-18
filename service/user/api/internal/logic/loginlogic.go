package logic

import (
	"context"
	"time"

	"LibSystem/common/utils"
	"LibSystem/service/user/api/internal/svc"
	"LibSystem/service/user/api/internal/types"
	"LibSystem/service/user/rpc/pb/user"

	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	l.ctx.Value("userId")
	res, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	// jwt 认证
	uid, err := strconv.ParseInt(res.Data.UserId, 10, 64)
	if err != nil {
		return nil, err
	}
	now := time.Now().Unix()
	accessToken, err := utils.GetJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, uid, res.Data.Role)
	if err != nil {
		return nil, err
	}
	return &types.LoginResponse{
		Base: types.Base{
			Code: res.Base.Code,
			Msg:  res.Base.Message,
		},
		Data: types.LoginResponseData{
			Token:    accessToken,
			UserID:   res.Data.UserId,
			Username: res.Data.Username,
			Role:     res.Data.Role,
		},
	}, nil
}
