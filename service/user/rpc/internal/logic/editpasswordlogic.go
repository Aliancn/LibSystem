package logic

import (
	"context"

	"LibSystem/common/utils"
	"LibSystem/service/user/rpc/internal/svc"
	"LibSystem/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type EditPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditPasswordLogic {
	return &EditPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditPasswordLogic) EditPassword(in *user.EditPasswordRequest) (*user.EditPasswordResponse, error) {
	// 查询是否存在
	loginUser, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	encryptedPassword := utils.MD5V(in.OldPassword, "", 0)
	if encryptedPassword != loginUser.Password {
		return nil, status.Errorf(100, "密码错误")
	}

	newEncryptedPassword := utils.MD5V(in.NewPassword, "", 0)
	loginUser.Password = newEncryptedPassword

	err = l.svcCtx.UserModel.Update(l.ctx, loginUser)
	if err != nil {
		return nil, err
	}

	return &user.EditPasswordResponse{
		Base: &user.Base{
			Code:    0,
			Message: "修改成功",
		},
	}, nil
}
