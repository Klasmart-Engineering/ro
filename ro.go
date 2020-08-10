package ro

import (
	"context"
	"errors"
	"gitlab.badanamu.com.cn/calmisland/common-log/log"
	"sync"

	"github.com/go-redis/redis"
)

var (
	_globalRedis *redis.ClusterClient
	_globalMutex sync.RWMutex
	_conf        *Config
)

var (
	ErrClientIsNil = errors.New("global redis is nil")
)

type Config struct {
	Addrs        []string
	Password     string
	PoolSize     int
	MinIdleConns int
}

func MustGetRedis(ctx context.Context) *redis.ClusterClient {
	client, err := GetRedis(ctx)
	if err != nil {
		log.Error(ctx, "Get redis failed", log.Err(err))
		panic("get redis failed")
	}

	return client

}

func GetRedis(ctx context.Context) (*redis.ClusterClient, error) {
	_globalMutex.Lock()
	defer _globalMutex.Unlock()

	if _globalRedis != nil {
		return _globalRedis, nil
	}

	//Create new redis cluster client
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        _conf.Addrs,
		Password:     _conf.Password,
		PoolSize:     _conf.PoolSize,
		MinIdleConns: _conf.MinIdleConns,
	})

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

func SetConfig(c *Config) {
	_conf = c
}
