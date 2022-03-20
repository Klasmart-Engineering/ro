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

	log.Debug(ctx, "expire value successfully", log.String("key", k.key), log.Duration("expiration", expiration))

	return nil
}

func (k Key) ExpireAt(ctx context.Context, t time.Time) error {
	err := MustGetRedis(ctx).ExpireAt(ctx, k.key, t).Err()
	if err != nil {
		log.Warn(ctx, "expire key failed", log.Err(err), log.String("key", k.key))
		return err
	}

	log.Debug(ctx, "expire value successfully", log.String("key", k.key), log.Time("time", t))

	return nil
}

func (k Key) TTL(ctx context.Context) (time.Duration, error) {
	ttl, err := MustGetRedis(ctx).TTL(ctx, k.key).Result()
	if err != nil {
		log.Warn(ctx, "get key ttl failed", log.Err(err), log.String("key", k.key))
		return 0, err
	}

	log.Debug(ctx, "get key ttl successfully", log.String("key", k.key), log.Duration("ttl", ttl))

	return ttl, nil
}

func (k Key) Exists(ctx context.Context) (bool, error) {
	count, err := MustGetRedis(ctx).Exists(ctx, k.key).Result()
	if err != nil {
		log.Warn(ctx, "check key exists failed", log.Err(err), log.String("key", k.key))
		return false, err
	}

	exists := count > 0
	log.Debug(ctx, "check key exists successfully", log.String("key", k.key), log.Bool("exists", exists))

	return exists, nil
}
