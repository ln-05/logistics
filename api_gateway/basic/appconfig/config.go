package appconfig

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Minio struct {
		Endpoint        string
		AccessKeyId     string
		AccessKeySecret string
		BucketName      string
		UseSsl          string
		BasePath        string
		BucketUrl       string
	}
	Wechat struct {
		AppID     string
		AppSecret string
	}
	Redis struct {
		Host     string
		Password string
		DB       int
	}
}

var AppCong Config

func InitConfig() {
	viper.SetConfigFile("../api_gateway/basic/appconfig/dev.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("配置文件读取失败")
	}
	log.Println("配置文件读取成功")
	err = viper.Unmarshal(&AppCong)
	if err != nil {
		panic("配置文件解析失败")
	}
	// log.Println("配置文件解析成功", AppCong)
}
