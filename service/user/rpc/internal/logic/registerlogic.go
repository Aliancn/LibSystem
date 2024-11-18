package logic

import (
	"context"

	"LibSystem/common"
	"LibSystem/common/utils"
	"LibSystem/service/user/model"
	"LibSystem/service/user/rpc/internal/svc"
	"LibSystem/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	_, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err == nil {
		return nil, status.Errorf(100, "用户已存在")
	}

	if err == model.ErrNotFound {
		userPasswordEncrypted := utils.MD5V(in.Password, "", 0)
		newUser := model.User{
			Username: in.Username,
			Password: userPasswordEncrypted,
			Role:     common.RoleUser,
		}

		_, err := l.svcCtx.UserModel.Insert(l.ctx, &newUser)
		if err != nil {
			return nil, status.Errorf(500, "注册失败"+err.Error())
		}

		return &user.RegisterResponse{
			Base: &user.Base{
				Code:    200,
				Message: "注册成功",
			}}, nil
	}
	return nil, err
}
