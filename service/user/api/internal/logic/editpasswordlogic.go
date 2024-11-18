package logic

import (
	"context"

	"LibSystem/service/user/api/internal/svc"
	"LibSystem/service/user/api/internal/types"
	"LibSystem/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditPasswordLogic {
	return &EditPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditPasswordLogic) EditPassword(req *types.EditPasswordRequest) (resp *types.EditPasswordResponse, err error) {
	res, err := l.svcCtx.UserRpc.EditPassword(l.ctx, &user.EditPasswordRequest{
		UserId:      req.UserId,
		NewPassword: req.NewPassword,
		OldPassword: req.OldPassword,
	})
	if err != nil {
		return nil, err
	}

	return &types.EditPasswordResponse{
		Base: types.Base{
			Code: res.Base.Code,
			Msg:  res.Base.Message,
		},
	}, nil
}
