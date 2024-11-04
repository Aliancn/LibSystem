package middle

import (
	"LibSystem/common"
	"LibSystem/common/utils"
	"LibSystem/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func VerifyJWTAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := common.SUCCESS
		token := c.Request.Header.Get(global.Config.Jwt.Admin.Name)
		// 解析获取用户载荷信息
		payLoad, err := utils.ParseToken(token, global.Config.Jwt.Admin.Secret)
		if err != nil {
			code = common.UNKNOW_IDENTITY
			c.JSON(http.StatusUnauthorized, common.Result{Code: code,
				Msg: common.GetMsg(code)})
			c.Abort()
			return
		}
		// 在上下文设置载荷信息
		c.Set(common.CurrentID, payLoad.UserId)
		c.Set(common.CurrentName, payLoad.GrantScope)
		// 这里是否要通知客户端重新保存新的Token
		c.Next()
	}
}

func VerifyJWTUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := common.SUCCESS
		token := c.Request.Header.Get(global.Config.Jwt.User.Name)
		// 解析获取用户载荷信息
		payLoad, err := utils.ParseToken(token, global.Config.Jwt.User.Secret)
		if err != nil {
			code = common.UNKNOW_IDENTITY
			c.JSON(http.StatusUnauthorized, common.Result{Code: code})
			c.Abort()
			return
		}
		// 在上下文设置载荷信息
		c.Set(common.CurrentID, payLoad.UserId)
		c.Set(common.CurrentName, payLoad.GrantScope)
		// 这里是否要通知客户端重新保存新的Token
		c.Next()
	}
}
