package global

import (
	"api_gateway/basic/appconfig"
	"github.com/go-redis/redis/v8"
)

var Minio appconfig.Config
var AppConfig appconfig.Config
var Rdb *redis.Client
