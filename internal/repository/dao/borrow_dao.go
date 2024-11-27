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
	var borrow model.Borrow
	err := b.db.WithContext(ctx).First(&borrow, id).Error
	return borrow, err
}

func (b BorrowDao) Update(ctx context.Context, borrow model.Borrow) error {
	err := b.db.WithContext(ctx).Updates(&borrow).Error
	return err
}

func (b BorrowDao) Delete(ctx context.Context, id int) error {
	err := b.db.WithContext(ctx).Delete(&model.Borrow{}, id).Error
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
	err := b.db.WithContext(ctx).Where("created_at > DATE_SUB(CURDATE(), INTERVAL ? DAY)", day).Find(&borrows).Error
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
