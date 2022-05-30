package cache

import (
	"CanRich/config"
	"CanRich/logger"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
)

var Redis RedisClient

type RedisClient struct {
	client   *redis.Client
	pipeline redis.Pipeliner
}

// InitRedisClient 初始化Redis
func InitRedisClient() error {
	host := config.GlobalConfig.GetString("redis.host")
	port := config.GlobalConfig.GetString("redis.port")
	dbName, _ := strconv.Atoi(config.GlobalConfig.GetString("redis.dbName"))
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
		DB:   dbName,
	})
	_, err := client.Ping().Result()
	if err != nil {
		logger.SugarLogger.Info(err)
		return err
	}
	Redis.client = client
	Redis.pipeline = client.Pipeline()
	return nil
}
