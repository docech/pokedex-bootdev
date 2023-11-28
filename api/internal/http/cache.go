package http

import (
	"sync"
	"time"
)

type CacheConfig struct {
	MaxAge       time.Duration
}

type cacheEntry[T any] struct {
	createdAt time.Time
	value     T
}

type cache[T any] struct {
	mu sync.Mutex
	entries map[string]cacheEntry[T]
}

func newByteCache(pruneInterval time.Duration) *cache[[]byte] {
	c := &cache[[]byte]{
		entries: map[string]cacheEntry[[]byte]{},
	}

	go schedulePruning(c, pruneInterval)

	return c
}

func (c *cache[T]) Get(key string) (noop T, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	if !ok {
		return noop, false
	}

	return entry.value, true
}

func (c *cache[T]) Set(key string, value T) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry[T]{
		createdAt: time.Now(),
		value:     value,
	}
}

func (c *cache[T]) prune(maxAge time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, entry := range c.entries {
		if time.Since(entry.createdAt) > maxAge {
			delete(c.entries, key)
		}
	}
}

func schedulePruning[T any](c *cache[T], maxAge time.Duration) {
	go func() {
		for range time.Tick(maxAge) {
			c.prune(maxAge)
		}
	}()
}