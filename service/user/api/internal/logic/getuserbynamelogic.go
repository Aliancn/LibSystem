package logic

import (
	"context"

	"LibSystem/service/user/api/internal/svc"
	"LibSystem/service/user/api/internal/types"
	"LibSystem/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByNameLogic {
	return &GetUserByNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByNameLogic) GetUserByName(req *types.GetUserByNameRequest) (resp *types.GetUserByNameResponse, err error) {
	res, err := l.svcCtx.UserRpc.GetUserByName(l.ctx, &user.GetUserByNameRequest{
		Username: req.Username,
	})

	if err != nil {
		return nil, err
	}

	return &types.GetUserByNameResponse{
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
