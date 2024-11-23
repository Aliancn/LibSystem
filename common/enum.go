package common

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS         = 1   // ok
	ERROR           = 2   // 内部错误
	UNKNOW_IDENTITY = 403 // 未知身份
)

var ErrMsg = map[int]string{
	SUCCESS:         "ok",
	ERROR:           "内部错误",
	UNKNOW_IDENTITY: "未知身份",
}

func GetMsg(code int) string {
	return ErrMsg[code]
}

func Send(ctx *gin.Context, code int) {
	ctx.JSON(http.StatusOK, Result{Code: code, Msg: GetMsg(code)})
}

const (
	// jwt的字段名
	CurrentID   = "current_id"
	CurrentRole = "current_role"
	CurrentName = "current_name"

	// 用户角色
	RoleAdmin = "admin"
	RoleUser  = "user"

	// book - available borrowed 
	// borrow - borrowed returned
	StatusAvailable = "available"
	StatusBorrowed  = "borrowed"
	StatusArchived  = "archived"
	StatusReturned  = "returned"
)

var (
	Error_PASSWORD_ERROR                 = errors.New("密码错误")
	Error_ACCOUNT_NOT_FOUND              = errors.New("账号不存在")
	Error_ACCOUNT_LOCKED                 = errors.New("账号被锁定")
	Error_ALREADY_EXISTS                 = errors.New("已存在")
	Error_UNKNOWN_ERROR                  = errors.New("未知错误")
	Error_USER_NOT_LOGIN                 = errors.New("用户未登录")
	Error_LOGIN_FAILED                   = errors.New("登录失败")
	Error_UPLOAD_FAILED                  = errors.New("文件上传失败")
	Error_PASSWORD_EDIT_FAILED           = errors.New("密码修改失败")
	// book
	Error_BOOK_NOT_FOUND                 = errors.New("书籍不存在")
	Error_BOOK_BORROWED                  = errors.New("书籍已被借出")
	Error_BOOK_RETURN_FAILED             = errors.New("书籍归还失败")
	Error_BOOK_DELETE_FAILED             = errors.New("书籍删除失败")
	Error_BOOK_CREATE_FAILED             = errors.New("书籍创建失败")
	Error_BOOK_UPDATE_FAILED             = errors.New("书籍更新失败")
	Error_BOOK_BORROW_FAILED             = errors.New("书籍借阅失败")
	Error_BOOK_RETURNED                  = errors.New("书籍已归还")
	Error_BOOK_BORROW_FAILED_NOT_FOUND   = errors.New("借阅记录不存在")
	Error_BOOK_BORROW_FAILED_NOT_ALLOWED = errors.New("不允许借阅")
	// borrow
	Error_BORROW_NOT_FOUND               = errors.New("借阅记录不存在")
	Error_BORROW_CREATE_FAILED           = errors.New("借阅记录创建失败")
	Error_BORROW_UPDATE_FAILED           = errors.New("借阅记录更新失败")
	Error_BORROW_DELETE_FAILED           = errors.New("借阅记录删除失败")
	Error_BORROW_GET_FAILED              = errors.New("借阅记录获取失败")
	Error_BORROW_GET_BY_USER_FAILED      = errors.New("借阅记录获取失败")
	// user
	Error_USER_NOT_FOUND                 = errors.New("用户不存在")
	Error_USER_CREATE_FAILED             = errors.New("用户创建失败")
	Error_USER_UPDATE_FAILED             = errors.New("用户更新失败")
	Error_USER_DELETE_FAILED             = errors.New("用户删除失败")
	Error_USER_GET_FAILED                = errors.New("用户获取失败")
	Error_USER_GET_BY_NAME_FAILED        = errors.New("用户获取失败")
	Error_USER_GET_BY_ID_FAILED          = errors.New("用户获取失败")
	Error_USER_GET_BY_ROLE_FAILED        = errors.New("用户获取失败")
	Error_USER_GET_BY_ROLE_NAME_FAILED   = errors.New("用户获取失败")
	Error_USER_GET_BY_ROLE_ID_FAILED     = errors.New("用户获取失败")
	Error_USER_GET_BY_ROLE_NAME_ID_FAILED = errors.New("用户获取失败")

)

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}
