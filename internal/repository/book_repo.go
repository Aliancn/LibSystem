package repository

import (
	"LibSystem/internal/model"
	"context"
)

type BookRepo interface {
	GetAll(ctx context.Context, pageID , pageSize int) ([]model.Book, error)
	GetByID(ctx context.Context, id int) (model.Book, error)
	GetByTitle(ctx context.Context, title string) ([]model.Book, error)
	Create(ctx context.Context, book model.Book) error
	Update(ctx context.Context, book model.Book) error
	Delete(ctx context.Context, id int) error
}
