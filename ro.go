package ro

import (
	"context"
	"errors"
	"sync"

	"github.com/go-redis/redis"
)

var (
	_globalRedis *redis.ClusterClient
	_globalOnce  sync.Once
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
		panic("get db context failed")
	}

	return client

}

func GetRedis(ctx context.Context) (*redis.ClusterClient, error) {
	var err error
	_globalOnce.Do(func() {
		var client *redis.ClusterClient
		//Create new redis cluster client
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:        _conf.Addrs,
			Password:     _conf.Password,
			PoolSize:     _conf.PoolSize,
			MinIdleConns: _conf.MinIdleConns,
		})

		//Try to ping redis server
		_, err := client.Ping().Result()
		if err != nil {
			return
		}

		_globalRedis = client
	})
	if err != nil {
		return nil, err
	}
	if _globalRedis == nil {
		return nil, ErrClientIsNil
	}

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
