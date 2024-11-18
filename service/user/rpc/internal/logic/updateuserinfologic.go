package logic

import (
	"context"
	"database/sql"

	"LibSystem/service/user/rpc/internal/svc"
	"LibSystem/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(in *user.UpdateUserInfoRequest) (*user.UpdateUserInfoResponse, error) {
	// 查询是否存在
	newUser, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	newUser.Name = sql.NullString{String: in.Name, Valid: in.Name != ""}
	newUser.Phone = sql.NullString{String: in.Phone, Valid: in.Phone != ""}
	newUser.Email = sql.NullString{String: in.Email, Valid: in.Email != ""}
	err = l.svcCtx.UserModel.Update(l.ctx, newUser)
	if err != nil {
		return nil, err
	}
	return &user.UpdateUserInfoResponse{
		Base: &user.Base{
			Code:    200,
			Message: "success",
		},
	}, nil
}
