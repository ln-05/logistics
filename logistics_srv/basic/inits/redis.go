package inits

import (
	"context"
	"logistics_srv/basic/global"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func InitRedis() {
	global.Rdb = redis.NewClient(&redis.Options{
		Addr:     global.Nacos.Redis.Addr,
		Password: global.Nacos.Redis.Password,
		DB:       global.Nacos.Redis.DB,
	})
	err := global.Rdb.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
		return
	}
	zap.L().Info("redis连接成功")

}
