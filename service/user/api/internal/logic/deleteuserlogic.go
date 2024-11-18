package logic

import (
	"context"

	"LibSystem/service/user/api/internal/svc"
	"LibSystem/service/user/api/internal/types"
	"LibSystem/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.DeleteUserRequest) (resp *types.DeleteUserResponse, err error) {
	res, err := l.svcCtx.UserRpc.DeleteUser(l.ctx, &user.DeleteUserRequest{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	return &types.DeleteUserResponse{
		Base: types.Base{
			Code: res.Base.Code,
			Msg:  res.Base.Message,
		},
	}, nil
}
