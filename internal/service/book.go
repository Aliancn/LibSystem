package service

import (
	"LibSystem/common"
	"LibSystem/internal/api/request"
	"LibSystem/internal/api/response"
	"LibSystem/internal/model"
	"LibSystem/internal/repository"

	"github.com/gin-gonic/gin"
)

type IBookService interface {
	GetBookList(ctx *gin.Context, pageID, pageSize int) (response.BookList, error)
	GetBookById(ctx *gin.Context, id int) (response.BookVO, error)
	GetBookByTitle(ctx *gin.Context, title string) (response.BookList, error)
	AddBook(ctx *gin.Context, add request.BookDTO) error
	UpdateBook(ctx *gin.Context, update request.BookDTO) error
	DeleteBook(ctx *gin.Context, id int) error
}

type BookService struct {
	repo repository.BookRepo
}

func (b BookService) GetBookList(ctx *gin.Context, pageID, pageSize int) (response.BookList, error) {
	var bookList response.BookList
	books, err := b.repo.GetAll(ctx, pageID, pageSize)
	if err != nil {
		return bookList, err
	}
	for _, book := range books {
		bookList.BookList = append(bookList.BookList, response.BookVO{
			BookId:      int64(book.ID),
			Title:       book.Title,
			Author:      book.Author,
			Publisher:   book.Publisher,
			Year:        book.Year,
			Genre:       book.Genre,
			Status:      book.Status,
			Location:    book.Location,
			BorrowTimes: book.BorrowTimes,
		})
	}
	return bookList, nil
}

func (b BookService) GetBookById(ctx *gin.Context, id int) (response.BookVO, error) {
	book, err := b.repo.GetByID(ctx, id)
	if err != nil {
		return response.BookVO{}, err
	}
	return response.BookVO{
		BookId:      int64(book.ID),
		Title:       book.Title,
		Author:      book.Author,
		Publisher:   book.Publisher,
		Year:        book.Year,
		Genre:       book.Genre,
		Status:      book.Status,
		Location:    book.Location,
		BorrowTimes: book.BorrowTimes,
	}, nil
}

func (b BookService) GetBookByTitle(ctx *gin.Context, title string) (response.BookList, error) {
	var bookList response.BookList
	books, err := b.repo.GetByTitle(ctx, title)
	if err != nil {
		return bookList, err
	}
	for _, book := range books {
		bookList.BookList = append(bookList.BookList, response.BookVO{
			BookId:      int64(book.ID),
			Title:       book.Title,
			Author:      book.Author,
			Publisher:   book.Publisher,
			Year:        book.Year,
			Genre:       book.Genre,
			Status:      book.Status,
			Location:    book.Location,
			BorrowTimes: book.BorrowTimes,
		})
	}
	return bookList, nil
}

func (b BookService) AddBook(ctx *gin.Context, add request.BookDTO) error {
	var book = model.Book{
		Title:     add.Title,
		Author:    add.Author,
		Publisher: add.Publisher,
		Year:      add.Year,
		Genre:     add.Genre,
		Status:    add.Status,
		Location:  add.Location,
	}
	err := b.repo.Create(ctx, book)
	if err != nil {
		return err
	}
	return nil

}

func (b BookService) UpdateBook(ctx *gin.Context, update request.BookDTO) error {
	var book = model.Book{
		ID:        update.BookID,
		Title:     update.Title,
		Author:    update.Author,
		Publisher: update.Publisher,
		Year:      update.Year,
		Genre:     update.Genre,
		Status:    update.Status,
		Location:  update.Location,
	}
	err := b.repo.Update(ctx, book)
	if err != nil {
		return err
	}
	return nil
}

func (b BookService) DeleteBook(ctx *gin.Context, id int) error {
	Book, err := b.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if Book.Status == "Borrowed" {
		return common.Error_BOOK_BORROWED
	}
	err = b.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func NewBookService(repo repository.BookRepo) IBookService {
	return BookService{repo: repo}
}
