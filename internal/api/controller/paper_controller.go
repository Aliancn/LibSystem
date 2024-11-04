package controller

import (
	"LibSystem/common"
	"LibSystem/global"
	"LibSystem/internal/api/request"
	"LibSystem/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type PaperController struct {
	service service.IPaperService
}

func NewPaperController(service service.IPaperService) *PaperController {
	return &PaperController{service: service}
}

// GetPaperList 获取所有论文
func (pc *PaperController) GetPaperList(ctx *gin.Context) {
	code := common.SUCCESS
	resp, err := pc.service.GetPaperList(ctx)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("PaperController GetList Error:", err.Error())
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

// GetPaperById 根据id获取论文
func (pc *PaperController) GetPaperById(ctx *gin.Context) {
	code := common.SUCCESS
	id, _ := strconv.Atoi(ctx.Param("id"))
	resp, err := pc.service.GetPaperById(ctx, id)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("PaperController GetPaperById Error:", err.Error())
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

// GetPaperByTitle 根据title获取论文
func (pc *PaperController) GetPaperByTitle(ctx *gin.Context) {
	code := common.SUCCESS
	title := ctx.Query("title")
	resp, err := pc.service.GetPaperByTitle(ctx, title)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("PaperController GetPaperByTitle Error:", err.Error())
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

// AddPaper 添加论文
func (pc *PaperController) AddPaper(ctx *gin.Context) {
	code := common.SUCCESS
	year, err := strconv.Atoi(ctx.PostForm("year"))
	var paperDTO = request.PaperDTO{
		Title:      ctx.PostForm("title"),
		Author:     ctx.PostForm("author"),
		Department: ctx.PostForm("department"),
		Year:       year,
	}
	if err != nil {
		code = common.ERROR
		global.Log.Warn("PaperController AddPaper BindJSON Error:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.Result{
			Code: common.ERROR,
			Msg:  "文件上传失败" + err.Error(),
		})
		return
	}

	// 设置文件保存路径
	uploadDir := "uploads/papers"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err := os.MkdirAll(uploadDir, os.ModePerm)
		if err != nil {
			return
		}
	}

	file.Filename = uuid.New().String() + filepath.Ext(file.Filename)
	filePath := filepath.Join(uploadDir, file.Filename)
	global.Log.Debug("file path:", filePath)
	paperDTO.FilePath = filePath
	// 保存文件到服务器
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, common.Result{
			Code: common.ERROR,
			Msg:  "文件保存失败" + err.Error(),
		})
		return
	}
	// 设置保存表单数据
	err = pc.service.AddPaper(ctx, paperDTO)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("PaperController AddPaper Error:", err.Error())
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

// UpdatePaper 更新论文
func (pc *PaperController) UpdatePaper(ctx *gin.Context) {
	code := common.SUCCESS
	var paperDTO request.PaperDTO
	err := ctx.BindJSON(&paperDTO)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("PaperController UpdatePaper BindJSON Error:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	err = pc.service.UpdatePaper(ctx, paperDTO)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("PaperController UpdatePaper Error:", err.Error())
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

// DeletePaper 删除论文
func (pc *PaperController) DeletePaper(ctx *gin.Context) {
	code := common.SUCCESS
	id, _ := strconv.Atoi(ctx.Param("id"))
	filePath, err := pc.service.GetPaperFilePath(ctx, id)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("PaperController DeletePaper Error:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	err = os.Remove(filePath)
	err = pc.service.DeletePaper(ctx, id)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("PaperController DeletePaper Error:", err.Error())
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

func (pc *PaperController) DownloadPaper(ctx *gin.Context) {
	code := common.SUCCESS
	id, _ := strconv.Atoi(ctx.Param("id"))
	filePath, err := pc.service.GetPaperFilePath(ctx, id)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("PaperController DownloadPaper Error:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	ctx.File(filePath)
}
