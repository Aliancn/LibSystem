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

type PaperDao struct {
	db *gorm.DB
}

func NewPaperDao(db *gorm.DB) repository.PaperRepo {
	return &PaperDao{db: db}
}

func (p PaperDao) GetByPaperName(ctx context.Context, paperName string) ([]model.Paper, error) {
	// 模糊查询
	var paper []model.Paper
	err := p.db.WithContext(ctx).Where("title like ?", "%"+paperName+"%").Find(&paper).Error
	if err != nil {
		return nil, err
	}
	if len(paper) == 0 {
		return nil, nil
	}
	return paper, nil
}
func (p PaperDao) GetNum(ctx context.Context) (int, error) {
	var count int64
	err := p.db.WithContext(ctx).Model(&model.Paper{}).Count(&count).Error
	return int(count), err
}

func (p PaperDao) GetById(ctx context.Context, id uint) (*model.Paper, error) {
	var paper model.Paper
	// 先从缓存中获取
	cacheKey := fmt.Sprintf("paper:%d", id)
	val, err := global.Redis.WithContext(ctx).Get(cacheKey).Result()
	if err == redis.Nil {
		fmt.Println("Cache miss, fetching from database...")
	} else if err != nil {
		return nil, err
	} else {
		fmt.Println("Cache hit!")
		err = json.Unmarshal([]byte(val), &paper)
		if err != nil {
			return nil, err
		}
		return &paper, nil
	}
	// 缓存中没有，从数据库中获取
	err = p.db.WithContext(ctx).First(&paper, id).Error
	if err != nil {
		return nil, err
	}
	// 将数据存入缓存
	data, err := json.Marshal(paper)
	if err != nil {
		return nil, err
	}
	err = global.Redis.WithContext(ctx).Set(cacheKey, data, 10*time.Minute).Err()
	if err != nil {
		return nil, err
	}
	return &paper, nil
}

func (p PaperDao) Update(ctx context.Context, paper model.Paper) error {
	// 更新缓存
	cacheKey := fmt.Sprintf("paper:%d", paper.ID)
	data, err := json.Marshal(paper)
	if err != nil {
		return err
	}
	err = global.Redis.WithContext(ctx).Set(cacheKey, data, 10*time.Minute).Err()
	if err != nil {
		return err
	}
	// 更新数据库
	err = p.db.WithContext(ctx).Model(&paper).Updates(paper).Error
	return err
}

func (p PaperDao) Insert(ctx context.Context, entity model.Paper) error {
	err := p.db.WithContext(ctx).Create(&entity).Error
	return err
}

func (p PaperDao) Delete(ctx context.Context, id uint) error {
	// 删除缓存
	cacheKey := fmt.Sprintf("paper:%d", id)
	err := global.Redis.WithContext(ctx).Del(cacheKey).Err()
	if err != nil {
		return err
	}
	// 删除数据库
	err = p.db.WithContext(ctx).Delete(&model.Paper{}, id).Error
	return err
}

func (p PaperDao) GetAll(ctx context.Context, pageID, pageSize int) ([]model.Paper, error) {
	var papers []model.Paper
	// err := p.db.WithContext(ctx).Find(&papers).Error
	err := p.db.WithContext(ctx).Limit(pageSize).Offset((pageID - 1) * pageSize).Find(&papers).Error
	if err != nil {
		return nil, err
	}
	return papers, nil
}

func (p PaperDao) GetFilePath(ctx context.Context, id uint) (string, error) {
	var paper model.Paper
	err := p.db.WithContext(ctx).First(&paper, id).Error
	if err != nil {
		return "", err
	}
	return paper.FilePath, nil
}
