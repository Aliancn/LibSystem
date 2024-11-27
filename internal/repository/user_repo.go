package repository

import (
	"LibSystem/internal/model"
	"context"
)

type UserRepo interface {
	// GetByUserName 根据username获取用户信息
	GetByUserName(ctx context.Context, userName string) (*model.User, error)
	// GetById 根据id获取用户信息
	GetById(ctx context.Context, id uint) (*model.User, error)
	// Update  修改
	Update(ctx context.Context, user model.User) error
	// Insert 插入数据
	Insert(ctx context.Context, entity model.User) error
	// Delete 删除数据
	Delete(ctx context.Context, id uint) error
	// GetAll 获取所有数据
	GetAll(ctx context.Context, pageID, pageSize int) ([]model.User, error)
	// PageQuery 分页查询
	//PageQuery(ctx context.Context, dto request.UserPageQueryDTO) (*common.PageResult, error)
	// UpdateStatus 更新状态
	//UpdateStatus(ctx context.Context, employee model.Employee) error
}
