package ro

import (
	"context"
	"errors"
	"sync"

	"gitlab.badanamu.com.cn/calmisland/common-log/log"

	"github.com/go-redis/redis"
)

var (
	_globalRedis *redis.Client
	_globalMutex sync.RWMutex
	_conf        *redis.Options
)

var (
	ErrClientIsNil = errors.New("global redis is nil")
	ErrConfIsNil   = errors.New("conf is nil")
	ErrKeyNotExist = redis.Nil
)

type Config struct {
	Addrs        []string
	Password     string
	PoolSize     int
	MinIdleConns int
}

func MustGetRedis(ctx context.Context) *redis.Client {
	client, err := GetRedis(ctx)
	if err != nil {
		log.Error(ctx, "Get redis failed", log.Err(err))
		panic("get redis failed")
	}

	return client

}

func GetRedis(ctx context.Context) (*redis.Client, error) {
	_globalMutex.Lock()
	defer _globalMutex.Unlock()

	if _globalRedis != nil {
		return _globalRedis, nil
	}

	if _conf == nil {
		log.Error(ctx, "Config is nil", log.Err(ErrConfIsNil))
		return nil, ErrConfIsNil
	}
	//Create new redis cluster client
	client := redis.NewClient(_conf)

	//Try to ping redis server
	_, err := client.Ping().Result()
	if err != nil {
		log.Error(ctx, "Connect to redis failed", log.Err(err))
		return nil, err
	}
	_globalRedis = client
	return _globalRedis, nil
}

func loadConfig() {
	if _conf != nil {
		return
	}
	//TODO: Load config from kr or env
}

func SetConfig(c *redis.Options) {
	_conf = c
}
