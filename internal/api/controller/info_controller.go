package controller

import (
	"LibSystem/common"
	"LibSystem/global"
	"LibSystem/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InfoController struct {
	service service.IInfoService
}

func NewInfoController(service service.IInfoService) InfoController {
	return InfoController{service: service}
}
func (ic *InfoController) GetInfo(ctx *gin.Context) {
	code := common.SUCCESS
	resp, err := ic.service.GetInfo(ctx)
	if err != nil {
		code = common.ERROR
		global.Log.Warn("InfoController GetServerInfo Error:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Msg:  "获取成功",
		Data: resp,
	})
}
