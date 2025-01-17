package ro

import (
	"context"
	"os"
	"testing"

	"github.com/go-redis/redis/v8"
)

func TestMain(m *testing.M) {
	SetConfig(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Username: "",
		Password: "",
		DB:       0,
	})

	os.Exit(m.Run())
}

func TestMustGetRedis(t *testing.T) {
	ctx := context.Background()
	client := MustGetRedis(ctx)
	if client == nil {
		t.Error("MustGetRedis() get nil")
	}
}
