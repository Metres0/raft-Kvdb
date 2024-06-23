package cache

import (
	"container/list"
	"sync"
)

// Cache是一个LRU缓存
type Cache struct {
	mu    sync.Mutex
	items map[string]*list.Element
	order *list.List
	max   int
}

type entry struct {
	key   string
	value string
}

// NewCache创建一个新的LRU缓存，最大容量为max
func NewCache(max int) *Cache {
	return &Cache{
		items: make(map[string]*list.Element),
		order: list.New(),
		max:   max,
	}
}

// Get从缓存中获取键对应的值
func (c *Cache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ele, ok := c.items[key]; ok {
		c.order.MoveToFront(ele)
		return ele.Value.(*entry).value, true
	}
	return "", false
}

// Add将键值对添加到缓存中
func (c *Cache) Add(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ele, ok := c.items[key]; ok {
		c.order.MoveToFront(ele)
		ele.Value.(*entry).value = value
		return
	}

	ele := c.order.PushFront(&entry{key, value})
	c.items[key] = ele
	if c.order.Len() > c.max {
		last := c.order.Back()
		if last != nil {
			c.order.Remove(last)
			delete(c.items, last.Value.(*entry).key)
		}
	}
}
