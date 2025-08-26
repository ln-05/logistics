package inits

import (
	"log"

	"go.uber.org/zap"
)

func Zap() {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"../logistics_srv/zap.log"}
	build, err := config.Build()
	if err != nil {
		panic("日志初始化失败")
	}
	zap.ReplaceGlobals(build)
	log.Println("日志初始化成功")
}
