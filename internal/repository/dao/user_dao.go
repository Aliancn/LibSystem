package dao

import (
	"LibSystem/internal/model"
	"LibSystem/internal/repository"
	"context"

	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

// GetByUserName 根据username获取用户信息
func (u *UserDao) GetByUserName(ctx context.Context, userName string) (*model.User, error) {
	var user model.User
	err := u.db.WithContext(ctx).Where("username like ?", userName).First(&user).Error
	return &user, err
}

// GetById 根据id获取用户信息
func (u *UserDao) GetById(ctx context.Context, id uint) (*model.User, error) {
	var user model.User
	err := u.db.WithContext(ctx).First(&user, id).Error
	return &user, err
}

// Update 动态修改
func (u *UserDao) Update(ctx context.Context, user model.User) error {
	return u.db.WithContext(ctx).Model(&user).Updates(user).Error
}

// Insert 插入数据
func (u *UserDao) Insert(ctx context.Context, entity model.User) error {
	return u.db.WithContext(ctx).Create(&entity).Error
}

func (u *UserDao) Delete(ctx context.Context, id uint) error {
	return u.db.WithContext(ctx).Delete(&model.User{}, id).Error
}
func (u *UserDao) GetAll(ctx context.Context, pageID, pageSize int) ([]model.User, error) {
	var users []model.User
	// 实现分页
	err := u.db.WithContext(ctx).Offset((pageID - 1) * pageSize).Limit(pageSize).Find(&users).Error
	// err := u.db.WithContext(ctx).Find(&users).Error
	return users, err
}
func (u *UserDao) GetNum(ctx context.Context) (int, error) {
	var count int64
	err := u.db.WithContext(ctx).Model(&model.User{}).Count(&count).Error
	return int(count), err
}

func NewUserDao(db *gorm.DB) repository.UserRepo {
	return &UserDao{db: db}
}
