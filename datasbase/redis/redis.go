package redis

import (
	"blog-server/configs"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	// 获取config
	ip := configs.GetValue("redis", "ip")
	port := configs.GetValue("redis", "port")
	password := configs.GetValue("redis", "password")
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", ip, port),
		Password: password, // 没有密码，默认值
		DB:       0,        // 默认DB 0
	})
	return rdb
}
