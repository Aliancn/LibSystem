package logic

import (
	"context"

	"LibSystem/service/user/rpc/internal/svc"
	"LibSystem/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllUserLogic {
	return &GetAllUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllUserLogic) GetAllUser(in *user.GetAllUserRequest) (*user.GetAllUserResponse, error) {
	users, err := l.svcCtx.UserModel.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}
	var userRespList user.UserList
	for _, _user := range users {
		userRespList.List = append(userRespList.List, &user.UserVO{
			UserId:   _user.Id,
			Username: _user.Username,
			Password: _user.Password,
			Phone:    _user.Phone.String,
			Email:    _user.Email.String,
		})
	}

	return &user.GetAllUserResponse{
		Base: &user.Base{
			Code:    200,
			Message: "success",
		},
		Data: &userRespList,
	}, nil
}
