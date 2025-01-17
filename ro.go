package ro

import (
	"context"
	"fmt"
	"sync"

	"github.com/KL-Engineering/common-log/log"
	"github.com/go-redis/redis/v8"
	"github.com/newrelic/go-agent/v3/integrations/nrredis-v8"
)

var (
	globalRedisClient *redis.Client
	globalMutex       sync.RWMutex
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
	globalMutex.Lock()
	defer globalMutex.Unlock()

	if globalRedisClient != nil {
		return globalRedisClient, nil
	}

	if globalOption == nil {
		log.Error(ctx, "global config undefined")
		return nil, ErrConfigUndefined
	}

	client := redis.NewClient(globalOption)
	client.AddHook(nrredis.NewHook(globalOption))

	err := client.Ping(ctx).Err()
	if err != nil {
		log.Error(ctx, "Connect to redis failed", log.Err(err))
		return nil, err
	}

	log.Debug(ctx, "connect to redis successfully", log.Any("option", fmt.Sprintf("%+v", globalOption)))

	globalRedisClient = client
	return globalRedisClient, nil
}

func SetConfig(c *redis.Options) {
	globalOption = c
}
