package service

import (
	"LibSystem/common"
	"LibSystem/internal/api/response"
	"LibSystem/internal/model"
	"LibSystem/internal/repository"
	"time"

	"github.com/gin-gonic/gin"
)

type IBorrowService interface {
	BorrowPaper(ctx *gin.Context, userID, paperID uint) error
	ReturnPaper(ctx *gin.Context, id uint) error
	DeletePaper(ctx *gin.Context, id uint) error
	GetAll(ctx *gin.Context) ([]response.BorrowVO, error)
	GetByUserID(ctx *gin.Context, userID int) ([]response.BorrowVO, error)
}

type BorrowService struct {
	borrowRepo repository.BorrowRepo
}

func (b BorrowService) BorrowPaper(ctx *gin.Context, userID, bookID uint) error {

	bo := model.Borrow{
		BorrowDate: time.Time{},
		ReturnDate: time.Time{},
		BookID:     int(bookID),
		UserID:     int(userID),
	}

	err := b.borrowRepo.Create(ctx, bo)
	if err != nil {
		return err
	}
	return nil
}

func (b BorrowService) ReturnPaper(ctx *gin.Context, id uint) error {
	bo := model.Borrow{
		ID:         id,
		ReturnDate: time.Now(),
		Status:     common.StatusReturned,
	}
	err := b.borrowRepo.Update(ctx, bo)
	if err != nil {
		return err
	}
	return nil
}

func (b BorrowService) DeletePaper(ctx *gin.Context, id uint) error {
	err := b.borrowRepo.Delete(ctx, int(id))
	if err != nil {
		return err
	}
	return nil
}

func (b BorrowService) GetAll(ctx *gin.Context) ([]response.BorrowVO, error) {
	borrows, err := b.borrowRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	var res []response.BorrowVO
	for _, borrow := range borrows {
		res = append(res, response.BorrowVO{
			ID:         borrow.ID,
			BorrowDate: borrow.BorrowDate,
			ReturnDate: borrow.ReturnDate,
			Status:     borrow.Status,
			Book:       borrow.BookID,
			User:       borrow.UserID,
		})
	}
	return res, nil
}

func (b BorrowService) GetByUserID(ctx *gin.Context, userID int) ([]response.BorrowVO, error) {
	borrows, err := b.borrowRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	var res []response.BorrowVO
	for _, borrow := range borrows {
		res = append(res, response.BorrowVO{
			ID:         borrow.ID,
			BorrowDate: borrow.BorrowDate,
			ReturnDate: borrow.ReturnDate,
			Status:     borrow.Status,
			Book:       borrow.BookID,
			User:       borrow.UserID,
		})
	}
	return res, nil
}

func NewBorrowService(borrowRepo repository.BorrowRepo) IBorrowService {
	return &BorrowService{borrowRepo: borrowRepo}
}
