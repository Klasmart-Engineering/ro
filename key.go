package ro

import (
	"context"
	"time"

	"gitlab.badanamu.com.cn/calmisland/common-log/log"
)

type Key struct {
	key string
}

func (k Key) Del(ctx context.Context) error {
	err := MustGetRedis(ctx).Del(ctx, k.key).Err()
	if err != nil {
		log.Warn(ctx, "delete key failed", log.Err(err), log.String("key", k.key))
		return err
	}

	log.Debug(ctx, "delete key successfully", log.String("key", k.key))

	return nil
}

func (k Key) Expire(ctx context.Context, expiration time.Duration) error {
	err := MustGetRedis(ctx).Expire(ctx, k.key, expiration).Err()
	if err != nil {
		log.Warn(ctx, "expire key failed", log.Err(err), log.String("key", k.key))
		return err
	}

	log.Debug(ctx, "expire value successfully", log.String("key", k.key))

	return nil
}
