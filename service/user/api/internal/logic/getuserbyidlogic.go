package logic

import (
	"context"

	"LibSystem/service/user/api/internal/svc"
	"LibSystem/service/user/api/internal/types"
	"LibSystem/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByIdLogic) GetUserById(req *types.GetUserByIdRequest) (resp *types.GetUserByIdResponse, err error) {
	res, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user.GetUserByIdRequest{
		Id: req.Id,
	})

	if err != nil {
		return nil, err
	}

	return &types.GetUserByIdResponse{
		Base: types.Base{
			Code: res.Base.Code,
			Msg:  res.Base.Message,
		},
		Data: types.UserVO{
			UserId:   res.Data.UserId,
			Username: res.Data.Username,
			Password: res.Data.Password,
			Role:     res.Data.Role,
			Name:     res.Data.Name,
			Phone:    res.Data.Phone,
			Email:    res.Data.Email,
		},
	}, nil
}
