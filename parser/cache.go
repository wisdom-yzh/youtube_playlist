package parser

import (
	"sync"
	"time"
)

type CachedData struct {
	data   interface{}
	key    string
	expire time.Time
}

type Cache struct {
	mu    sync.RWMutex
	ttl   time.Duration
	cache map[string]CachedData
}

func NewCache(ttl time.Duration) *Cache {
	return &Cache{
		mu:    sync.RWMutex{},
		cache: map[string]CachedData{},
		ttl:   ttl,
	}
}

func (c *Cache) Get(key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()

	val, exist := c.cache[key]
	if !exist || time.Now().After(val.expire) {
		return nil
	}

	return val.data
}

func (c *Cache) Set(key string, data interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = CachedData{
		key:    key,
		data:   data,
		expire: time.Now().Add(c.ttl),
	}
}
