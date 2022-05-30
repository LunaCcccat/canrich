package cache

import (
	"CanRich/logger"
	"github.com/go-redis/redis"
	"strings"
	"time"
)

func (c *RedisClient) CommonKeyGen(fields []string) string {
	return strings.Join(fields, ":")
}

func (c *RedisClient) SetHash(key string, fields map[string]interface{}) (string, error) {
	val, err := c.client.HMSet(key, fields).Result()
	if err != nil {
		logger.SugarLogger.Errorf("Redis set hash failed. error: %s", err.Error())
	}
	return val, err
}

func (c *RedisClient) BatchSetHash(keys []string, objs []map[string]interface{}) error {
	for index, key := range keys {
		c.pipeline.HMSet(key, objs[index])
	}
	_, err := c.pipeline.Exec()
	return err
}

func (c *RedisClient) BatchGetHash(keys []string) ([]map[string]string, error) {
	for _, key := range keys {
		c.pipeline.HGetAll(key)
	}
	cmds, err := c.pipeline.Exec()
	if err != nil {
		logger.SugarLogger.Errorf("Redis get hashes failed. error: %s", err.Error())
		return nil, err
	}
	var result []map[string]string
	for _, cmd := range cmds {
		if stringStringMap, ok := cmd.(*redis.StringStringMapCmd); ok {
			res, err := stringStringMap.Result()
			if err != nil {
				logger.SugarLogger.Errorf("Redis get hashes failed. error: %s", err.Error())
				return nil, err
			}
			result = append(result, res)
		} else {
			logger.SugarLogger.Errorf("Redis get hashes failed.")
		}
	}
	return result, nil
}

func (c *RedisClient) BatchDelete(keys []string) error {
	for _, key := range keys {
		c.pipeline.Del(key)
	}
	_, err := c.pipeline.Exec()
	return err
}

func (c *RedisClient) GetKeys(pattern string) ([]string, error) {
	var (
		res, keys []string
		cursor    uint64
		err       error
	)
	for {
		res, cursor, err = c.client.Scan(cursor, pattern, 20).Result()
		if err != nil {
			return nil, err
		}
		keys = append(keys, res...)
		if cursor == 0 {
			break
		}
	}
	return keys, nil
}

func (c *RedisClient) GetString(key string) (string, error) {
	return c.client.Get(key).Result()
}

func (c *RedisClient) SetString(key string, value string, expiration time.Duration) {
	c.client.Set(key, value, expiration)
}

func (c *RedisClient) Delete(key string) {
	c.client.Del(key)
}
