package ro

import (
	"context"
	"testing"
	"time"
)

func TestConnectRedisCluster(t *testing.T){
	SetConfig(&Config{
		Addrs:        []string{"test-redis.bxypks.ng.0001.apn2.cache.amazonaws.com:6379"},
		Password:     "",
		PoolSize:     10,
		MinIdleConns: 4,
	})
	client := MustGetRedis(context.Background())
	client.Set(":mytest", "1", time.Minute * 30)

	t.Log(client.Get(":mytest").String())
}
