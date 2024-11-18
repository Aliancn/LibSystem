package logic

import (
	"context"

	"LibSystem/service/user/api/internal/svc"
	"LibSystem/service/user/api/internal/types"
	"LibSystem/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(req *types.UpdateUserInfoRequest) (resp *types.UpdateUserInfoResponse, err error) {
	res, err := l.svcCtx.UserRpc.UpdateUserInfo(l.ctx, &user.UpdateUserInfoRequest{
		UserId: req.UserId,
		Name:   req.Name,
		Phone:  req.Phone,
		Email:  req.Email,
	})

	if err != nil {
		return nil, err
	}

	return &types.UpdateUserInfoResponse{
		Base: types.Base{
			Code: res.Base.Code,
			Msg:  res.Base.Message,
		},
	}, nil
}
