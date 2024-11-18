package logic

import (
	"context"

	"LibSystem/service/user/api/internal/svc"
	"LibSystem/service/user/api/internal/types"
	"LibSystem/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllUserLogic {
	return &GetAllUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllUserLogic) GetAllUser(req *types.GetAllUserRequest) (resp *types.GetAllUserResponse, err error) {
	res, err := l.svcCtx.UserRpc.GetAllUser(l.ctx, &user.GetAllUserRequest{})

	if err != nil {
		return nil, err
	}

	var data types.UserList

	for _, userinfo := range res.Data.List {
		data.List = append(data.List, types.UserVO{
			UserId:   userinfo.UserId,
			Username: userinfo.Username,
			Password: userinfo.Password,
			Role:     userinfo.Role,
			Name:     userinfo.Name,
			Phone:    userinfo.Phone,
			Email:    userinfo.Email,
		})
	}

	return &types.GetAllUserResponse{
		Base: types.Base{
			Code: res.Base.Code,
			Msg:  res.Base.Message,
		},
		Data: data,
	}, nil
}
