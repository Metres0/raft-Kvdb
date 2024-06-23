// pkg/cache/cache_test.go
package cache_test

import (
	"raft-Kvdb/pkg/cache"
	"testing"
)

func TestCache(t *testing.T) {
	cache := cache.NewCache(2)

	cache.Add("key1", "value1")
	cache.Add("key2", "value2")

	if value, ok := cache.Get("key1"); !ok || value != "value1" {
		t.Fatalf("Expected key1 to be 'value1', got %v", value)
	}

	cache.Add("key3", "value3")

	if _, ok := cache.Get("key2"); ok {
		t.Fatalf("Expected key2 to be evicted")
	}

	if value, ok := cache.Get("key3"); !ok || value != "value3" {
		t.Fatalf("Expected key3 to be 'value3', got %v", value)
	}
}
