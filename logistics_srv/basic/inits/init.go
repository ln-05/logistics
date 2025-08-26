package inits

import "logistics_srv/basic/appconfig"

func Init() {
	appconfig.InitConfig()
	Zap()
	InitNacos()
	InitMysql()
	InitRedis()
}
