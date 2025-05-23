package cache

import "sync"

type Cache[K comparable, V any] struct {
	mu sync.RWMutex
	m  map[K]V
}

func New[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{m: make(map[K]V)}
}

func (c *Cache[K, V]) Get(key K) (v V, ok bool) {
	c.mu.RLock()
	v, ok = c.m[key]
	c.mu.RUnlock()
	return
}

func (c *Cache[K, V]) Set(key K, value V) {
	c.mu.Lock()
	c.m[key] = value
	c.mu.Unlock()
}
