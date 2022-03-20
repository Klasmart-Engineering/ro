package ro

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestStringKey_SetAndGet(t *testing.T) {
	ctx := context.Background()

	key := NewStringKey("test1")
	defer key.Del(ctx)

	value := "111"
	expiration := time.Millisecond * 50

	err := key.Set(ctx, value, expiration)
	if err != nil {
		t.Errorf("set string value failed due to %v", err)
	}

	get, err := key.Get(ctx)
	if err != nil {
		t.Errorf("get string value failed due to %v", err)
	}

	if get != value {
		t.Errorf("get string value not equal, want %s, get %s", value, get)
	}

	time.Sleep(expiration * time.Duration(2))

	_, err = key.Get(ctx)
	if err != ErrRecordNotFound {
		t.Errorf("string value expire failed")
	}
}

func TestStringKey_SetAndGetInt(t *testing.T) {
	ctx := context.Background()

	key := NewStringKey("test2")
	defer key.Del(ctx)

	value := 222

	err := key.SetInt(ctx, value, 0)
	if err != nil {
		t.Errorf("set int value failed due to %v", err)
	}

	get, err := key.GetInt(ctx)
	if err != nil {
		t.Errorf("get int value failed due to %v", err)
	}

	if get != value {
		t.Errorf("get int value not equal, want %d, get %d", value, get)
	}
}

func TestStringKey_SetAndGetInt64(t *testing.T) {
	ctx := context.Background()

	key := NewStringKey("test3")
	defer key.Del(ctx)

	value := int64(333)

	err := key.SetInt64(ctx, value, 0)
	if err != nil {
		t.Errorf("set int64 value failed due to %v", err)
	}

	get, err := key.GetInt64(ctx)
	if err != nil {
		t.Errorf("get int64 value failed due to %v", err)
	}

	if get != value {
		t.Errorf("get int64 value not equal, want %d, get %d", value, get)
	}
}

type testStruct struct {
	AAA string
	BBB int
}

func TestStringKey_SetAndGetObject(t *testing.T) {
	ctx := context.Background()

	key := NewStringKey("test4")
	defer key.Del(ctx)

	obj := &testStruct{
		AAA: "aaa",
		BBB: 444,
	}

	err := key.SetObject(ctx, obj, 0)
	if err != nil {
		t.Errorf("set object value failed due to %v", err)
	}

	get := &testStruct{}
	err = key.GetObject(ctx, get)
	if err != nil {
		t.Errorf("get object value failed due to %v", err)
	}

	if !reflect.DeepEqual(get, obj) {
		t.Errorf("get object value not equal, want %v, get %v", obj, get)
	}
}
