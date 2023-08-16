package middleware

import "github.com/go-redis/redis/v8"

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.31.191:6379", // Redis 服务器地址
		Password: "123456",              // Redis 访问密码
		DB:       0,                     // 使用的数据库编号
	})
	return rdb
}
