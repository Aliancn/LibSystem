package logic

import (
	"context"
	"strconv"

	"LibSystem/common/utils"
	"LibSystem/service/user/rpc/internal/svc"
	"LibSystem/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {

	loginUser, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		return nil, status.Errorf(100, "用户不存在")
	}

	userPasswordEncrypted := utils.MD5V(in.Password, "", 0)
	if loginUser.Password != userPasswordEncrypted {
		return nil, status.Errorf(100, "密码错误")
	}
	return &user.LoginResponse{
		Base: &user.Base{
			Code:    0,
			Message: "登录成功",
		},
		Data: &user.LoginResponseData{
			UserId:   strconv.FormatInt(loginUser.Id, 10),
			Username: loginUser.Username,
			Role:     loginUser.Role,
		},
	}, nil
}
