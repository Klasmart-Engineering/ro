package ro

import (
	"context"
	"testing"
)

func TestSetKey_SAdd(t *testing.T) {
	ctx := context.Background()

	key := NewSetKey("test2")
	defer key.Del(ctx)

	type fields struct {
		Key Key
	}
	type args struct {
		ctx     context.Context
		members []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := SetKey{
				Key: tt.fields.Key,
			}
			if err := k.SAdd(tt.args.ctx, tt.args.members...); (err != nil) != tt.wantErr {
				t.Errorf("SetKey.SAdd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
