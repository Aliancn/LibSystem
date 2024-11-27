package dao

import (
	"LibSystem/global"
	"LibSystem/internal/model"
	"LibSystem/internal/repository"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

// GetByUserName 根据username获取用户信息
func (u *UserDao) GetByUserName(ctx context.Context, userName string) (*model.User, error) {
	var user model.User
	cache := fmt.Sprintf("user:%s", userName)
	// 先从缓存中获取
	val, err := global.Redis.WithContext(ctx).Get(cache).Result()
	if err == redis.Nil {
		fmt.Println("Cache miss, fetching from database...")
	} else if err != nil {
		return nil, err
	} else {
		fmt.Println("Cache hit!")
		err = json.Unmarshal([]byte(val), &user)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}
	// 缓存中没有，从数据库中获取
	err = u.db.WithContext(ctx).Where("username like ?", userName).First(&user).Error
	if err != nil {
		return nil, err
	}
	// 将数据存入缓存
	data, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	err = global.Redis.WithContext(ctx).Set(cache, data, 10*time.Minute).Err()
	if err != nil {
		return nil, err
	}
	return &user, err
}

// GetById 根据id获取用户信息
func (u *UserDao) GetById(ctx context.Context, id uint) (*model.User, error) {
	var user model.User
	cacheKey := fmt.Sprintf("user:%d", id)
	// 先从缓存中获取
	val, err := global.Redis.WithContext(ctx).Get(cacheKey).Result()
	if err == redis.Nil {
		fmt.Println("Cache miss, fetching from database...")
	} else if err != nil {
		return nil, err
	} else {
		fmt.Println("Cache hit!")
		err = json.Unmarshal([]byte(val), &user)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}
	// 缓存中没有，从数据库中获取
	err = u.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return nil, err
	}
	// 将数据存入缓存
	data, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	err = global.Redis.WithContext(ctx).Set(cacheKey, data, 10*time.Minute).Err() // 0表示永不过期
	if err != nil {
		return nil, err
	}
	return &user, err
}

// Update 动态修改
func (u *UserDao) Update(ctx context.Context, user model.User) error {
	// 更新缓存
	cacheKey := fmt.Sprintf("user:%d", user.ID)
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = global.Redis.WithContext(ctx).Set(cacheKey, data, 10*time.Minute).Err()
	if err != nil {
		return err
	}
	// 更新数据库
	return u.db.WithContext(ctx).Model(&user).Updates(user).Error
}

// Insert 插入数据
func (u *UserDao) Insert(ctx context.Context, entity model.User) error {
	return u.db.WithContext(ctx).Create(&entity).Error
}

func (u *UserDao) Delete(ctx context.Context, id uint) error {
	// 删除缓存
	cacheKey := fmt.Sprintf("user:%d", id)
	err := global.Redis.WithContext(ctx).Del(cacheKey).Err()
	if err != nil {
		return err
	}
	// 删除数据库
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
