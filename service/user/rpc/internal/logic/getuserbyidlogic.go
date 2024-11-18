package logic

import (
	"context"

	"LibSystem/service/user/rpc/internal/svc"
	"LibSystem/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *user.GetUserByIdRequest) (*user.GetUserByIdResponse, error) {
	_user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &user.GetUserByIdResponse{
		Base: &user.Base{
			Code:    200,
			Message: "success",
		},
		Data: &user.UserVO{
			UserId:   _user.Id,
			Username: _user.Username,
			Password: _user.Password,
			Role:     _user.Role,
			Name:     _user.Name.String,
			Phone:    _user.Phone.String,
			Email:    _user.Email.String,
		},
	}, nil
}
