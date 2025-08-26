package global

import (
	"logistics_srv/basic/appconfig"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	Rdb   *redis.Client
	Nacos appconfig.Config
)
