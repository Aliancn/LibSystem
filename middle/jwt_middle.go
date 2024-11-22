package middle

import (
	"LibSystem/common"
	"LibSystem/common/utils"
	"LibSystem/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func VerifyJWTAdmin() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		code := common.SUCCESS
// 		token := c.Request.Header.Get(global.Config.Jwt.Name)
// 		// 解析获取用户载荷信息
// 		payLoad, err := utils.ParseToken(token, global.Config.Jwt.Secret)
// 		if err != nil {
// 			code = common.UNKNOW_IDENTITY
// 			c.JSON(http.StatusUnauthorized, common.Result{Code: code,
// 				Msg: common.GetMsg(code)})
// 			c.Abort()
// 			return
// 		}
// 		// 在上下文设置载荷信息
// 		c.Set(common.CurrentID, payLoad.UserId)
// 		c.Set(common.CurrentName, payLoad.GrantScope)
// 		// 这里是否要通知客户端重新保存新的Token
// 		c.Next()
// 	}
// }

// func VerifyJWTUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		code := common.SUCCESS
// 		token := c.Request.Header.Get(global.Config.Jwt.Name)
// 		// 解析获取用户载荷信息
// 		payLoad, err := utils.ParseToken(token, global.Config.Jwt.Secret)
// 		if err != nil {
// 			code = common.UNKNOW_IDENTITY
// 			c.JSON(http.StatusUnauthorized, common.Result{Code: code})
// 			c.Abort()
// 			return
// 		}
// 		// 在上下文设置载荷信息
// 		c.Set(common.CurrentID, payLoad.UserId)
// 		c.Set(common.CurrentName, payLoad.GrantScope)
// 		// 这里是否要通知客户端重新保存新的Token
// 		c.Next()
// 	}
// }

func VerifyJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := common.SUCCESS
		token := c.Request.Header.Get(global.Config.Jwt.Name)
		// 解析获取用户载荷信息
		token = token[7:]
		payLoad, err := utils.ParseJwtToken(global.Config.Jwt.Secret, token)
		if err != nil {
			code = common.UNKNOW_IDENTITY
			global.Log.Error("jwt parse error: ", err)
			c.JSON(http.StatusUnauthorized, common.Result{Code: code, Msg: err.Error()})
			c.Abort()
			return
		}
		// 在上下文设置载荷信息
		c.Set(common.CurrentID, payLoad["uid"])
		c.Set(common.CurrentName, payLoad["role"])
		
		// 这里是否要通知客户端重新保存新的Token
		c.Next()
	}
}
