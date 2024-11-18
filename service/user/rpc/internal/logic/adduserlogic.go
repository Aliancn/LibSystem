package logic

import (
	"context"
	"database/sql"

	"LibSystem/service/user/model"
	"LibSystem/service/user/rpc/internal/svc"
	"LibSystem/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserLogic) AddUser(in *user.AddUserRequest) (*user.AddUserResponse, error) {
	_, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err == nil {
		return nil, status.Errorf(100, "用户已存在")
	}

	if err == model.ErrNotFound {
		newUser := model.User{
			Username: in.Username,
			Password: in.Password,
			Role:     in.Role,
			Name:     sql.NullString{String: in.Name, Valid: in.Name != ""},
			Phone:    sql.NullString{String: in.Phone, Valid: in.Phone != ""},
			Email:    sql.NullString{String: in.Email, Valid: in.Email != ""},
		}

		_, err := l.svcCtx.UserModel.Insert(l.ctx, &newUser)
		if err != nil {
			return nil, status.Errorf(500, "添加失败"+err.Error())
		}

		return &user.AddUserResponse{
			Base: &user.Base{
				Code:    200,
				Message: "添加成功",
			}}, nil

	}

	return &user.AddUserResponse{}, nil
}
