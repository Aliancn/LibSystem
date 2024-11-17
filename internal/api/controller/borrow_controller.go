package controller

import (
	"LibSystem/common"
	"LibSystem/global"
	"LibSystem/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BorrowController struct {
	service service.IBorrowService
}

func NewBorrowController(service service.IBorrowService) *BorrowController {
	return &BorrowController{service: service}
}

// Borrow 借书
func (b *BorrowController) Borrow(ctx *gin.Context) {
	code := common.SUCCESS
	user, _ := strconv.Atoi(ctx.Query("user_id"))
	book, _ := strconv.Atoi(ctx.Query("book_id"))
	err := b.service.BorrowPaper(ctx, uint(user), uint(book))
	if err != nil {
		code = common.ERROR
		global.Log.Warn("borrowController GetBookByTitle Error:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Msg:  "Borrowed successfully",
	})
}

// Return 还书
func (b *BorrowController) Return(ctx *gin.Context) {
	code := common.SUCCESS
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := b.service.ReturnPaper(ctx, uint(id))
	if err != nil {
		code = common.ERROR
		global.Log.Warn("borrowController ReturnPaper Error:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Msg:  "Returned successfully",
	})
}

// Delete 删除借阅记录
func (b *BorrowController) Delete(ctx *gin.Context) {
	code := common.SUCCESS
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := b.service.DeletePaper(ctx, uint(id))
	if err != nil {
		code = common.ERROR
		global.Log.Warn("borrowController DeletePaper Error:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Msg:  "Deleted successfully",
	})
}

// GetAll 获取所有借阅记录
func (b *BorrowController) GetAll(ctx *gin.Context) {
	code := common.SUCCESS
	resp, err := b.service.GetAll(ctx)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("borrowController GetAll Error:", err.Error())
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

func (b *BorrowController) GetByUser(ctx *gin.Context) {
	code := common.SUCCESS
	id, _ := strconv.Atoi(ctx.Query("user_id"))
	resp, err := b.service.GetByUserID(ctx, id)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("borrowController GetByUser Error:", err.Error())
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
