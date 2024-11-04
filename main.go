package main

import (
	"LibSystem/global"
	"LibSystem/initialize"
	"github.com/gin-gonic/gin"
)

func main() {

	router := initialize.GlobalInit()

	gin.SetMode(global.Config.Server.Level)

	router.Run(":8080")
}
