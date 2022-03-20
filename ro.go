package ro

import (
	"context"
	"sync"

	"github.com/go-redis/redis/v8"
	"gitlab.badanamu.com.cn/calmisland/common-log/log"
)

var (
	globalRedisClient *redis.Client
	_globalMutex      sync.RWMutex
	globalOption      *redis.Options
)

func MustGetRedis(ctx context.Context) *redis.Client {
	client, err := GetRedis(ctx)
	if err != nil {
		log.Panic(ctx, "Get redis failed", log.Err(err))
	}

	return client
}

func GetRedis(ctx context.Context) (*redis.Client, error) {
	_globalMutex.Lock()
	defer _globalMutex.Unlock()

	if globalRedisClient != nil {
		return globalRedisClient, nil
	}

	if globalOption == nil {
		log.Error(ctx, "global config undefined")
		return nil, ErrConfigUndefined
	}

	client := redis.NewClient(globalOption)

	err := client.Ping(ctx).Err()
	if err != nil {
		log.Error(ctx, "Connect to redis failed", log.Err(err))
		return nil, err
	}

	log.Debug(ctx, "connect to redis successfully", log.Any("option", globalOption))

	globalRedisClient = client
	return globalRedisClient, nil
}

func SetConfig(c *redis.Options) {
	globalOption = c
}
