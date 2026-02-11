package gokecache

import (
	"sync"
	"time"
)

type Cache struct {
	CacheEntries map[string]cacheEntry
	mu           *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
