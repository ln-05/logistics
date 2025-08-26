package appconfig

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Mysql struct {
		Username string
		Password string
		Host     string
		Port     int
		Database string
	}
	Redis struct {
		Addr     string
		Password string
		DB       int
	}
}

type Nacos struct {
	NamespaceID string
	DataId      string
	Group       string
	IdAddr      string
	Port        int
}

var AppConf Nacos

func InitConfig() {
	viper.SetConfigFile("../logistics_srv/basic/appconfig/dev.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("配置文件读取失败")
	}
	log.Println("配置文件读取成功")
	err = viper.Unmarshal(&AppConf)
	if err != nil {
		panic("配置文件解析失败")
	}
	log.Println("配置文件解析成功")
}
