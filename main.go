package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"raft-Kvdb/pkg/kvstore"
)

func main() {
	// 初始化节点
	storageDir := "./data"
	peers := []string{"localhost:8001", "localhost:8002", "localhost:8003"}
	me := 0 // 当前节点在peers中的索引
	kv := kvstore.NewKVStore(peers, me, storageDir)

	// 处理Ctrl+C信号，优雅地关闭KVStore
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Println("Shutting down...")
		os.Exit(0)
	}()

	// 启动HTTP服务器以便其他节点连接（示例，不包含实际实现）
	// go startHTTPServer(kv, me, peers[me])

	// 演示一些基本的KV操作
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		kv.Set(key, value)
	}

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		value, exists := kv.Get(key)
		if exists {
			fmt.Printf("Retrieved %s: %s\n", key, value)
		} else {
			fmt.Printf("Key %s not found\n", key)
		}
	}

	// 保持程序运行
	select {}
}
