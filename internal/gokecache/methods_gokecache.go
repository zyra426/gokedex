package gokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		CacheEntries: make(map[string]cacheEntry),
		mu:           &sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	ce := cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}

	c.CacheEntries[key] = ce
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	ce, exists := c.CacheEntries[key]
	return ce.val, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.Tick(interval)

	for range ticker {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, entry := range c.CacheEntries {
		if entry.createdAt.Before(now.Add(-last)) {
			delete(c.CacheEntries, key)
		}
	}
}
