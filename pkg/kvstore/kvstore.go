// pkg/kvstore/kvstore.go
package kvstore

import (
	"sync"

	"raft-Kvdb/pkg/cache"
	"raft-Kvdb/pkg/raft"
	"raft-Kvdb/wal"
)

type KVStore struct {
	data      map[string]string
	dataMutex sync.RWMutex
	cache     *cache.Cache
	wal       *wal.WAL
	raft      *raft.Raft
}

func NewKVStore(peers []string, me int, storageDir string) *KVStore {
	r := raft.NewRaft(peers, me)
	return &KVStore{
		data:  make(map[string]string),
		cache: cache.NewCache(100),
		wal:   wal.NewWAL(storageDir),
		raft:  r,
	}
}

func (kv *KVStore) Set(key, value string) {
	kv.dataMutex.Lock()
	defer kv.dataMutex.Unlock()

	kv.data[key] = value
	kv.cache.Add(key, value)
	kv.wal.Write("set", key, value)

	if kv.raft.State() == raft.Leader {
		kv.raft.AppendEntries([]raft.LogEntry{
			{Operation: "set", Key: key, Value: value},
		})
	}
}

func (kv *KVStore) Get(key string) (string, bool) {
	if value, ok := kv.cache.Get(key); ok {
		return value, true
	}

	kv.dataMutex.RLock()
	defer kv.dataMutex.RUnlock()

	value, exists := kv.data[key]
	if exists {
		kv.cache.Add(key, value)
	}
	return value, exists
}
