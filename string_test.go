package ro

import (
	"context"
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

	get, err = key.Get(ctx)
	if err != ErrRecordNotFound {
		t.Errorf("string value expire failed")
	}
}
