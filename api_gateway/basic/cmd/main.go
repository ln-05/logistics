package main

import (
	"api_gateway/basic/appconfig"
	"api_gateway/basic/global"
	"api_gateway/basic/inits"
	"api_gateway/router"

	"github.com/gin-gonic/gin"
)

func main() {

	appconfig.InitConfig()
	global.AppConfig = appconfig.AppCong
	inits.InitRedist()
	r := gin.Default()

	router.LoadRouter(r)
	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
