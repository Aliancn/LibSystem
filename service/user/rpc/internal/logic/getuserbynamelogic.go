package logic

import (
	"context"

	"LibSystem/service/user/rpc/internal/svc"
	"LibSystem/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByNameLogic {
	return &GetUserByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByNameLogic) GetUserByName(in *user.GetUserByNameRequest) (*user.GetUserByNameResponse, error) {
	_user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		return nil, err
	}
	return &user.GetUserByNameResponse{
		Base: &user.Base{Code: 200, Message: "success"},
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
