package g

import (
	redis "github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var DB *gorm.DB
var CacheDB *redis.ClusterClient
