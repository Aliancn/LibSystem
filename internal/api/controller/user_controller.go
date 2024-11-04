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

type UserController struct {
	service service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{service: userService}
}

func (uc *UserController) Login(ctx *gin.Context) {
	code := common.SUCCESS
	userLogin := request.UserLogin{}
	err := ctx.Bind(&userLogin)
	if err != nil {
		code = common.ERROR
		global.Log.Debug("UserController login 解析失败")
		return
	}
	resp, err := uc.service.Login(ctx, userLogin)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("UserController login Error:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Msg:  "登录成功",
		Data: resp,
	})
}

func (uc *UserController) Logout(ctx *gin.Context) {
	code := common.SUCCESS
	var err error
	err = uc.service.Logout(ctx)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("UserController login Error:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Msg:  "退出成功",
	})
}

func (uc *UserController) Register(ctx *gin.Context) {
	code := common.SUCCESS
	userRegister := request.UserRegister{}
	err := ctx.Bind(&userRegister)
	if err != nil {
		code = common.ERROR
		global.Log.Debug("UserController Register 解析失败")
		return
	}
	err = uc.service.RegisterRole(ctx, userRegister)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("UserController Register Error:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Msg:  "注册成功",
	})
}

func (uc *UserController) GetById(ctx *gin.Context) {
	code := common.SUCCESS
	var userId request.UserID
	uid, err := strconv.Atoi(ctx.Param("id"))
	userId.UserId = int64(uid)

	resp, err := uc.service.GetById(ctx, userId)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("UserController GetById Error:", err.Error())
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

func (uc *UserController) GetByUsername(ctx *gin.Context) {
	code := common.SUCCESS
	var username request.Username
	username.Username = ctx.Query("username")
	global.Log.Info("GetByUsername:", username)
	resp, err := uc.service.GetByUsername(ctx, username)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("UserController GetByUsername Error:", err.Error())
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

func (uc *UserController) AddUser(ctx *gin.Context) {
	code := common.SUCCESS
	userAdd := request.UserDTO{}
	err := ctx.Bind(&userAdd)
	if err != nil {
		code = common.ERROR
		global.Log.Debug("UserController AddUser 解析失败")
		return
	}
	err = uc.service.AddUser(ctx, userAdd)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("UserController AddUser Error:", err.Error())
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

func (uc *UserController) EditPassword(ctx *gin.Context) {
	code := common.SUCCESS
	var reqs request.UserEditPassword
	var err error
	err = ctx.Bind(&reqs)
	if err != nil {
		global.Log.Debug("EditPassword Error:", err.Error())
		return
	}
	err = uc.service.EditPassword(ctx, reqs)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("EditPassword Error:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Msg:  "修改成功",
	})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	code := common.SUCCESS
	var userId request.UserID
	err := ctx.Bind(&userId)
	if err != nil {
		code = common.ERROR
		global.Log.Debug("UserController DeleteUser 解析失败")
		return
	}
	err = uc.service.DeleteUser(ctx, userId)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("UserController DeleteUser Error:", err.Error())
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

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	code := common.SUCCESS
	userUpdate := request.UserDTO{}
	err := ctx.Bind(&userUpdate)
	if err != nil {
		code = common.ERROR
		global.Log.Debug("UserController UpdateUser 解析失败")
		return
	}
	err = uc.service.UpdateUser(ctx, userUpdate)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("UserController UpdateUser Error:", err.Error())
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

func (uc *UserController) GetList(ctx *gin.Context) {
	code := common.SUCCESS
	resp, err := uc.service.GetList(ctx)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("UserController GetList Error:", err.Error())
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
