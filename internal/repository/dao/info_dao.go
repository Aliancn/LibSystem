package dao

import (
	"LibSystem/internal/model"
	"LibSystem/internal/repository"
	"context"

	"gorm.io/gorm"
)

type InfoDao struct {
	db *gorm.DB
}

func (i InfoDao) GetInfo(ctx context.Context, day int) ([]model.Info, error) {
	var infos []model.Info
	err := i.db.WithContext(ctx).Where("download_time > CURRENT_DATE - INTERVAL ? DAY", day).Find(&infos).Error
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, nil
	}
	return infos, nil
}

func (i InfoDao) AddInfo(ctx context.Context, info model.Info) error {
	err := i.db.WithContext(ctx).Create(&info).Error
	return err
}

func NewInfoDao(db *gorm.DB) repository.InfoRepo {
	return &InfoDao{db: db}
}
