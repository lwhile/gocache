package lru

import (
	"reflect"
	"testing"
)

func TestNewCache(t *testing.T) {
	tests := []struct {
		name string
		want *Cache
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCache(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
