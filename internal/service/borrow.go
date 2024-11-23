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
	bookRepo   repository.BookRepo
}

func (b BorrowService) BorrowPaper(ctx *gin.Context, userID, bookID uint) error {

	bo := model.Borrow{
		BorrowDate: time.Now(),
		ReturnDate: time.Now(),
		BookID:     int(bookID),
		UserID:     int(userID),
	}
	book, err := b.bookRepo.GetByID(ctx, int(bookID))
	if err != nil {
		return err
	}
	if book.Status == common.StatusBorrowed {
		return common.Error_BOOK_BORROWED
	}
	// book -> borrowed
	book.Status = common.StatusBorrowed
	err = b.bookRepo.Update(ctx, book)
	if err != nil {
		return err
	}
	err = b.borrowRepo.Create(ctx, bo)
	if err != nil {
		return err
	}
	return nil
}

func (b BorrowService) ReturnPaper(ctx *gin.Context, id uint) error {
	bo, err := b.borrowRepo.GetByID(ctx, int(id))
	if err != nil {
		return err
	}
	bo.Status = common.StatusReturned
	bo.ReturnDate = time.Now()
	book, err := b.bookRepo.GetByID(ctx, bo.BookID)
	if err != nil {
		return err
	}
	// book -> available
	book.Status = common.StatusAvailable
	err = b.bookRepo.Update(ctx, book)
	if err != nil {
		return err
	}
	err = b.borrowRepo.Update(ctx, bo)
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

func NewBorrowService(borrowRepo repository.BorrowRepo, bookRepo repository.BookRepo) IBorrowService {
	return &BorrowService{borrowRepo: borrowRepo, bookRepo: bookRepo}
}
