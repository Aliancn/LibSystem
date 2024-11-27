package repository

import (
	"LibSystem/internal/model"
	"context"
)

type BorrowRepo interface {
	GetAll(ctx context.Context, pageID, pageSize int) ([]model.Borrow, error)
	GetByID(ctx context.Context, id int) (model.Borrow, error)
	Update(ctx context.Context, borrow model.Borrow) error
	Delete(ctx context.Context, id int) error
	Create(ctx context.Context, borrow model.Borrow) error
	GetByUserID(ctx context.Context, userID int) ([]model.Borrow, error)
	GetBorrowInfo(ctx context.Context, day int) ([]model.Borrow, error)
}
