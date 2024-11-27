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

type BookDao struct {
	db *gorm.DB
}

func (b BookDao) GetAll(ctx context.Context, pageID, pageSize int) ([]model.Book, error) {
	var books []model.Book
	// err := b.db.Find(&books).Error
	err := b.db.WithContext(ctx).Offset((pageID - 1) * pageSize).Limit(pageSize).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (b BookDao) GetByID(ctx context.Context, id int) (model.Book, error) {
	var book model.Book
	// 缓存
	cacheKey := fmt.Sprintf("book:%d", id)
	val, err := global.Redis.WithContext(ctx).Get(cacheKey).Result()
	if err == redis.Nil {
		fmt.Println("Cache miss, fetching from database...")
	} else if err != nil {
		return book, err
	} else {
		fmt.Println("Cache hit!")
		err = json.Unmarshal([]byte(val), &book)
		if err != nil {
			return book, err

		}
		return book, nil
	}
	// 缓存中没有，从数据库中获取
	err = b.db.WithContext(ctx).First(&book, id).Error
	if err != nil {
		return model.Book{}, err
	}
	// 存入缓存
	data, err := json.Marshal(book)
	if err != nil {
		return model.Book{}, err
	}
	err = global.Redis.WithContext(ctx).Set(cacheKey, data, time.Minute*10).Err()
	if err != nil {
		return model.Book{}, err
	}
	return book, nil
}

func (b BookDao) GetByTitle(ctx context.Context, title string) ([]model.Book, error) {
	var books []model.Book
	err := b.db.WithContext(ctx).Where("title like ?", title).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (b BookDao) Create(ctx context.Context, book model.Book) error {
	err := b.db.WithContext(ctx).Create(&book).Error
	if err != nil {
		return err
	}
	return nil
}

func (b BookDao) Update(ctx context.Context, book model.Book) error {
	// 更新缓存
	cacheKey := fmt.Sprintf("book:%d", book.ID)
	val, err := json.Marshal(book)
	if err != nil {
		return err
	}
	err = global.Redis.WithContext(ctx).Set(cacheKey, val, 0).Err()
	if err != nil {
		return err
	}
	// 更新数据库
	err = b.db.WithContext(ctx).Model(&book).Updates(book).Error
	if err != nil {
		return err
	}
	return nil
}

func (b BookDao) Delete(ctx context.Context, id int) error {
	// 删除缓存
	cacheKey := fmt.Sprintf("book:%d", id)
	err := global.Redis.WithContext(ctx).Del(cacheKey).Err()
	if err != nil {
		return err
	}
	// 删除数据库
	err = b.db.WithContext(ctx).Delete(&model.Book{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (b BookDao) GetNum(ctx context.Context) (int, error) {
	var count int64
	err := b.db.WithContext(ctx).Model(&model.Book{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func NewBookDao(db *gorm.DB) repository.BookRepo {
	return &BookDao{db: db}
}
