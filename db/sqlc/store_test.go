package db

import (
	"context"
	"reflect"
	"testing"
)

func TestNewStore(t *testing.T) {
	type args struct {
		ctx     context.Context
		psqlURI string
	}
	tests := []struct {
		name string
		args args
		want Store
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStore(tt.args.ctx, tt.args.psqlURI); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStore() = %v, want %v", got, tt.want)
			}
		})
	}
}
