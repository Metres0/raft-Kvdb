// wal/wal_test.go
package wal

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestWAL(t *testing.T) {
	dir, err := ioutil.TempDir("", "waltest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	wal := NewWAL(dir)
	defer wal.Close()

	wal.Write("set", "key1", "value1")

	data, err := ioutil.ReadFile(filepath.Join(dir, "wal.log"))
	if err != nil {
		t.Fatal(err)
	}

	expected := "set key1 value1\n"
	if string(data) != expected {
		t.Fatalf("Expected '%s', got '%s'", expected, string(data))
	}
}
