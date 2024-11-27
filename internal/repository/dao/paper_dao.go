package dao

import (
	"LibSystem/internal/model"
	"LibSystem/internal/repository"
	"context"

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

func (p PaperDao) GetById(ctx context.Context, id uint) (*model.Paper, error) {
	var paper model.Paper
	err := p.db.WithContext(ctx).First(&paper, id).Error
	if err != nil {
		return nil, err
	}
	return &paper, nil
}

func (p PaperDao) Update(ctx context.Context, paper model.Paper) error {
	err := p.db.WithContext(ctx).Updates(&paper).Error
	return err
}

func (p PaperDao) Insert(ctx context.Context, entity model.Paper) error {
	err := p.db.WithContext(ctx).Create(&entity).Error
	return err
}

func (p PaperDao) Delete(ctx context.Context, id uint) error {
	err := p.db.WithContext(ctx).Delete(&model.Paper{}, id).Error
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
