package dao

import (
	"LibSystem/global"
	"LibSystem/internal/model"
	"LibSystem/internal/repository"
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type BorrowDao struct {
	db *gorm.DB
}

func (b BorrowDao) Create(ctx context.Context, borrow model.Borrow) error {
	err := b.db.WithContext(ctx).Create(&borrow).Error
	return err
}

func (b BorrowDao) GetAll(ctx context.Context, pageID, pageSize int) ([]model.Borrow, error) {
	var borrows []model.Borrow
	// err := b.db.WithContext(ctx).Find(&borrows).Error
	err := b.db.WithContext(ctx).Offset((pageID - 1) * pageSize).Limit(pageSize).Find(&borrows).Error
	if err != nil {
		return nil, err
	}
	if len(borrows) == 0 {
		return nil, nil
	}
	return borrows, nil
}

func (b BorrowDao) GetByID(ctx context.Context, id int) (model.Borrow, error) {
	// TODO 返回值的类型
	var borrow model.Borrow
	// 缓存
	cacheKey := fmt.Sprintf("borrow:%d", id)
	val, err := global.Redis.WithContext(ctx).Get(cacheKey).Result()
	if err == redis.Nil {
		fmt.Println("Cache miss, fetching from database...")
	} else if err != nil {
		return borrow, err
	} else {
		fmt.Println("Cache hit!")
		err = json.Unmarshal([]byte(val), &borrow)
		if err != nil {
			return borrow, err
		}
	}
	// 缓存中没有，从数据库中获取
	err = b.db.WithContext(ctx).First(&borrow, id).Error
	if err != nil {
		return borrow, err
	}
	// 存入缓存
	data, err := json.Marshal(borrow)
	if err != nil {
		return borrow, err
	}
	err = global.Redis.WithContext(ctx).Set(cacheKey, data, 0).Err()
	if err != nil {
		return borrow, err
	}
	return borrow, err
}

func (b BorrowDao) Update(ctx context.Context, borrow model.Borrow) error {
	// 更新缓存
	cacheKey := fmt.Sprintf("borrow:%d", borrow.ID)
	val, err := json.Marshal(borrow)
	if err != nil {
		return err
	}
	err = global.Redis.WithContext(ctx).Set(cacheKey, val, 0).Err()
	if err != nil {
		return err
	}
	// 更新数据库
	err = b.db.WithContext(ctx).Model(&borrow).Updates(borrow).Error
	return err
}

func (b BorrowDao) Delete(ctx context.Context, id int) error {
	// 删除缓存
	cacheKey := fmt.Sprintf("borrow:%d", id)
	err := global.Redis.WithContext(ctx).Del(cacheKey).Err()
	if err != nil {
		return err
	}
	// 删除数据库
	err = b.db.WithContext(ctx).Delete(&model.Borrow{}, id).Error
	return err
}

func (b BorrowDao) GetByUserID(ctx context.Context, userID int) ([]model.Borrow, error) {
	var borrows []model.Borrow
	err := b.db.WithContext(ctx).Where("user_id = ?", userID).Find(&borrows).Error
	if err != nil {
		return nil, err
	}
	if len(borrows) == 0 {
		return nil, nil
	}
	return borrows, nil
}

func (b BorrowDao) GetBorrowInfo(ctx context.Context, day int) ([]model.Borrow, error) {
	var borrows []model.Borrow
	err := b.db.WithContext(ctx).Where("borrow_date > CURRENT_DATE - INTERVAL ? DAY", day).Find(&borrows).Error
	if err != nil {
		return nil, err
	}
	if len(borrows) == 0 {
		return nil, nil
	}
	return borrows, nil
}

func NewBorrowDao(db *gorm.DB) repository.BorrowRepo {
	return &BorrowDao{db: db}
}
