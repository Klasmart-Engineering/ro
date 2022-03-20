package ro

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gitlab.badanamu.com.cn/calmisland/common-log/log"
)

type StringKey struct {
	Key
}

func NewStringKey(key string) *StringKey {
	return &StringKey{Key: Key{key: key}}
}

func (k StringKey) Get(ctx context.Context) (string, error) {
	value, err := MustGetRedis(ctx).Get(ctx, k.key).Result()
	if err == redis.Nil {
		return "", ErrRecordNotFound
	}

	if err != nil {
		log.Warn(ctx, "get key value failed", log.Err(err), log.String("key", k.key))
		return "", err
	}

	log.Debug(ctx, "get value successfully",
		log.String("key", k.key),
		log.String("value", value))

	return value, nil
}

func (k StringKey) GetDefault(ctx context.Context, defaultValue string) string {
	value, err := k.Get(ctx)
	if err != nil {
		log.Debug(ctx, "get value failed, use default value instead",
			log.Err(err), log.String("key", k.key),
			log.String("defaultValue", defaultValue))
		return defaultValue
	}

	return value
}

func (k StringKey) Set(ctx context.Context, value string, expiration time.Duration) error {
	err := MustGetRedis(ctx).Set(ctx, k.key, value, expiration).Err()
	if err != nil {
		log.Warn(ctx, "set key value failed",
			log.Err(err),
			log.String("key", k.key),
			log.String("value", value))
		return err
	}

	log.Debug(ctx, "set value successfully",
		log.String("key", k.key),
		log.String("value", value))

	return nil
}

func (k StringKey) SetNX(ctx context.Context, value string, expiration time.Duration) (bool, error) {
	result, err := MustGetRedis(ctx).SetNX(ctx, k.key, value, expiration).Result()
	if err != nil {
		log.Warn(ctx, "setnx key value failed",
			log.Err(err),
			log.String("key", k.key),
			log.String("value", value))
		return false, err
	}

	log.Debug(ctx, "setnx value successfully",
		log.String("key", k.key),
		log.String("value", value),
		log.Bool("result", result))

	return result, nil
}

type StringParameterKey struct {
	pattern string
}

func NewStringParameterKey(pattern string) *StringParameterKey {
	return &StringParameterKey{pattern: pattern}
}

func (k StringParameterKey) Param(parameters ...interface{}) *StringKey {
	return NewStringKey(fmt.Sprintf(k.pattern, parameters...))
}
