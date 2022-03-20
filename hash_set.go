package ro

import (
	"context"
	"encoding/json"
	"fmt"

	"gitlab.badanamu.com.cn/calmisland/common-log/log"
)

type HashSetKey struct {
	Key
}

func NewHashSetKey(key string) *HashSetKey {
	return &HashSetKey{Key: Key{key: key}}
}

func (k HashSetKey) HGet(ctx context.Context, field string) (string, error) {
	value, err := MustGetRedis(ctx).HGet(ctx, k.key, field).Result()
	if err != nil {
		log.Warn(ctx, "get field value failed",
			log.Err(err),
			log.String("key", k.key),
			log.String("field", field))
		return "", err
	}

	log.Debug(ctx, "get field value successfully",
		log.String("key", k.key),
		log.String("field", field),
		log.String("value", value))

	return value, nil
}

func (k HashSetKey) HGetObject(ctx context.Context, field string, value interface{}) error {
	result, err := k.HGet(ctx, field)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(result), value)
	if err != nil {
		log.Warn(ctx, "json unmarshal failed", log.Err(err), log.String("result", result))
		return err
	}

	log.Debug(ctx, "get object successfully",
		log.String("key", k.key),
		log.String("field", field),
		log.Any("object", value))

	return nil
}

func (k HashSetKey) HMGet(ctx context.Context, fields []string) (map[string]string, error) {
	values, err := MustGetRedis(ctx).HMGet(ctx, k.key, fields...).Result()
	if err != nil {
		log.Warn(ctx, "get fields value failed",
			log.Err(err),
			log.String("key", k.key))
		return nil, err
	}

	if len(values) != len(fields) {
		log.Warn(ctx, "result count not equal to request count",
			log.Int("requestCount", len(fields)),
			log.Int("resultCount", len(values)))
		return nil, ErrInvalidResultCount
	}

	result := make(map[string]string, len(fields))
	for index, value := range values {
		if value == nil {
			continue
		}

		stringValue, ok := value.(string)
		if !ok {
			continue
		}

		result[fields[index]] = stringValue
	}

	log.Debug(ctx, "get fields value successfully",
		log.String("key", k.key),
		log.Strings("fields", fields),
		log.Any("result", result))

	return result, nil
}

func (k HashSetKey) HGetAll(ctx context.Context) (map[string]string, error) {
	values, err := MustGetRedis(ctx).HGetAll(ctx, k.key).Result()
	if err != nil {
		log.Warn(ctx, "get all fields value failed",
			log.Err(err),
			log.String("key", k.key))
		return nil, err
	}

	log.Debug(ctx, "get all fields value successfully",
		log.String("key", k.key),
		log.Any("result", values))

	return values, nil
}

func (k HashSetKey) HSet(ctx context.Context, field, value string) error {
	err := MustGetRedis(ctx).HSet(ctx, k.key, field, value).Err()
	if err != nil {
		log.Warn(ctx, "set value failed",
			log.Err(err),
			log.String("key", k.key),
			log.String("field", field),
			log.String("value", value))
		return err
	}

	log.Debug(ctx, "set value successfully",
		log.String("key", k.key),
		log.String("field", field),
		log.String("value", value))

	return nil
}

func (k HashSetKey) HSetObject(ctx context.Context, field string, value interface{}) error {
	buffer, err := json.Marshal(value)
	if err != nil {
		log.Warn(ctx, "json marshal failed", log.Err(err), log.Any("value", value))
		return err
	}

	return k.HSet(ctx, field, string(buffer))
}

func (k HashSetKey) HDel(ctx context.Context, field ...string) error {
	if len(field) == 0 {
		return nil
	}

	err := MustGetRedis(ctx).HDel(ctx, k.key, field...).Err()
	if err != nil {
		log.Warn(ctx, "del field failed",
			log.Err(err),
			log.String("key", k.key),
			log.Strings("field", field))
		return err
	}

	log.Debug(ctx, "del field successfully",
		log.String("key", k.key),
		log.Strings("field", field))

	return nil
}

func (k HashSetKey) HLen(ctx context.Context) (int64, error) {
	count, err := MustGetRedis(ctx).HLen(ctx, k.key).Result()
	if err != nil {
		log.Warn(ctx, "get fields count failed",
			log.Err(err),
			log.String("key", k.key))
		return 0, err
	}

	log.Debug(ctx, "get fields count  successfully",
		log.String("key", k.key),
		log.Int64("count", count))

	return count, nil
}

type HashSetParameterKey struct {
	pattern string
}

func NewHashSetParameterKey(pattern string) *HashSetParameterKey {
	return &HashSetParameterKey{pattern: pattern}
}

func (k HashSetParameterKey) Param(parameters ...interface{}) *HashSetKey {
	return NewHashSetKey(fmt.Sprintf(k.pattern, parameters...))
}
