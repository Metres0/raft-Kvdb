// pkg/kvstore/kvstore_test.go
package kvstore

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestKVStore(t *testing.T) {
	dir, err := ioutil.TempDir("", "kvstoretest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	peers := []string{"0", "1", "2"}
	kv := NewKVStore(peers, 0, dir)

	kv.Set("key1", "value1")
	value, exists := kv.Get("key1")
	if !exists || value != "value1" {
		t.Fatalf("Expected key1 to be 'value1', got %v", value)
	}
}
