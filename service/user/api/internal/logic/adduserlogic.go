package logic

import (
	"context"

	"LibSystem/service/user/api/internal/svc"
	"LibSystem/service/user/api/internal/types"
	"LibSystem/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.AddUserRequest) (resp *types.AddUserResponse, err error) {
	res, err := l.svcCtx.UserRpc.AddUser(l.ctx, &user.AddUserRequest{
		Username: req.Username,
		Password: req.Password,
		Role:     req.Role,
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
	})
	if err != nil {
		return nil, err
	}
	return &types.AddUserResponse{
		Base: types.Base{
			Code: res.Base.Code,
			Msg:  res.Base.Message,
		},
	}, nil

}
