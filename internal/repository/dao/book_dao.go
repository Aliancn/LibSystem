package dao

import (
	"LibSystem/internal/model"
	"LibSystem/internal/repository"
	"context"
	"gorm.io/gorm"
)

type BookDao struct {
	db *gorm.DB
}

func (b BookDao) GetAll(ctx context.Context) ([]model.Book, error) {
	var books []model.Book
	err := b.db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (b BookDao) GetByID(ctx context.Context, id int) (model.Book, error) {
	var book model.Book
	err := b.db.First(&book, id).Error
	if err != nil {
		return model.Book{}, err
	}
	return book, nil
}

func (b BookDao) GetByTitle(ctx context.Context, title string) ([]model.Book, error) {
	var books []model.Book
	err := b.db.Where("title like ?", title).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (b BookDao) Create(ctx context.Context, book model.Book) error {
	err := b.db.Create(&book).Error
	if err != nil {
		return err
	}
	return nil
}

func (b BookDao) Update(ctx context.Context, book model.Book) error {
	err := b.db.Updates(&book).Error
	if err != nil {
		return err
	}
	return nil
}

func (b BookDao) Delete(ctx context.Context, id int) error {
	err := b.db.Delete(&model.Book{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func NewBookDao(db *gorm.DB) repository.BookRepo {
	return &BookDao{db: db}
}
