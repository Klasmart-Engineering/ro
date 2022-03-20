package ro

import (
	"context"
	"testing"
)

func TestKey_Exists(t *testing.T) {
	ctx := context.Background()

	key := NewStringKey("test:abc")
	err := key.Set(ctx, "a", 0)
	if err != nil {
		t.Errorf("set string value failed due to %v", err)
	}

	exists, err := key.Exists(ctx)
	if err != nil {
		t.Errorf("check key exists failed due to %v", err)
	}

	if !exists {
		t.Errorf("check key exists failed, want true, get %v", exists)
	}

	err = key.Del(ctx)
	if err != nil {
		t.Errorf("delete key failed due to %v", err)
	}

	exists, err = key.Exists(ctx)
	if err != nil {
		t.Errorf("check key exists failed due to %v", err)
	}

	if exists {
		t.Errorf("check key exists failed, want false, get %v", exists)
	}
}
