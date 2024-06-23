// wal/wal.go
package wal

import (
	"fmt"
	"os"
	"path/filepath"
)

// WAL是预写日志文件
type WAL struct {
	file *os.File
}

// NewWAL创建一个新的WAL实例
func NewWAL(storageDir string) *WAL {
	walPath := filepath.Join(storageDir, "wal.log")
	file, err := os.OpenFile(walPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("Failed to open WAL file: %v", err))
	}
	return &WAL{file: file}
}

// Write将操作记录写入WAL文件
func (wal *WAL) Write(operation, key, value string) {
	entry := fmt.Sprintf("%s %s %s\n", operation, key, value)
	if _, err := wal.file.WriteString(entry); err != nil {
		panic(fmt.Sprintf("Failed to write to WAL: %v", err))
	}
}

// Close关闭WAL文件
func (wal *WAL) Close() {
	wal.file.Close()
}
