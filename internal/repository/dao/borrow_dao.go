package dao

import (
	"LibSystem/internal/model"
	"LibSystem/internal/repository"
	"context"
	"gorm.io/gorm"
)

type BorrowDao struct {
	db *gorm.DB
}

func (b BorrowDao) Create(ctx context.Context, borrow model.Borrow) error {
	err := b.db.WithContext(ctx).Create(&borrow).Error
	return err
}

func (b BorrowDao) GetAll(ctx context.Context) ([]model.Borrow, error) {
	var borrows []model.Borrow
	err := b.db.WithContext(ctx).Find(&borrows).Error
	if err != nil {
		return nil, err
	}
	if len(borrows) == 0 {
		return nil, nil
	}
	return borrows, nil
}

func (b BorrowDao) GetByID(ctx context.Context, id int) (model.Borrow, error) {
	//TODO implement me
	panic("implement me")
}

func (b BorrowDao) Update(ctx context.Context, borrow model.Borrow) error {
	err := b.db.WithContext(ctx).Updates(&borrow).Error
	return err
}

func (b BorrowDao) Delete(ctx context.Context, id int) error {
	err := b.db.WithContext(ctx).Delete(&model.Borrow{}, id).Error
	return err
}

func NewBorrowDao(db *gorm.DB) repository.BorrowRepo {
	return &BorrowDao{db: db}
}
