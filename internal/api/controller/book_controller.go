package controller

import (
	"LibSystem/common"
	"LibSystem/global"
	"LibSystem/internal/api/request"
	"LibSystem/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BookController struct {
	service service.IBookService
}

func NewBookController(bookService service.IBookService) *BookController {
	return &BookController{service: bookService}
}

func (bc *BookController) GetBookList(ctx *gin.Context) {
	code := common.SUCCESS
	resp, err := bc.service.GetBookList(ctx)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("bookController GetList Error:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Data: resp,
	})
}

func (bc *BookController) GetBookById(ctx *gin.Context) {
	code := common.SUCCESS
	id, _ := strconv.Atoi(ctx.Param("id"))
	resp, err := bc.service.GetBookById(ctx, id)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("bookController GetBookById Error:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Data: resp,
	})
}

func (bc *BookController) GetBookByTitle(ctx *gin.Context) {
	code := common.SUCCESS
	title := ctx.Query("title")
	resp, err := bc.service.GetBookByTitle(ctx, title)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("bookController GetBookByTitle Error:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Data: resp,
	})
}

func (bc *BookController) AddBook(ctx *gin.Context) {
	code := common.SUCCESS
	var book request.BookDTO
	err := ctx.ShouldBindJSON(&book)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("bookController AddBook Error:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	err = bc.service.AddBook(ctx, book)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("bookController AddBook Error:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Msg:  "添加成功",
	})
}

func (bc *BookController) UpdateBook(ctx *gin.Context) {
	code := common.SUCCESS
	var book request.BookDTO
	err := ctx.ShouldBindJSON(&book)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("bookController UpdateBook Error:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	err = bc.service.UpdateBook(ctx, book)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("bookController UpdateBook Error:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Msg:  "更新成功",
	})
}

func (bc *BookController) DeleteBook(ctx *gin.Context) {
	code := common.SUCCESS
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := bc.service.DeleteBook(ctx, id)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("bookController DeleteBook Error:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Msg:  "删除成功",
	})
}
