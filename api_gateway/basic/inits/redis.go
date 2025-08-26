package inits

import (
	"api_gateway/basic/appconfig"
	"api_gateway/basic/global"
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func InitRedist() {
	global.Rdb = redis.NewClient(&redis.Options{
		Addr:     appconfig.AppCong.Redis.Host,
		Password: appconfig.AppCong.Redis.Password,
		DB:       appconfig.AppCong.Redis.DB,
	})
	err := global.Rdb.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
		return
	}
	log.Println("redis连接成功")

}
