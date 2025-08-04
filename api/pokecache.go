package api

import (
	"context"
	"sync"
	"time"
)

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

type cache struct {
	mu          sync.RWMutex
	items       map[string]cacheEntry
	cancelCache context.CancelFunc
}

var SharedCache = NewCache(5 * time.Second)

func NewCache(interval time.Duration) *cache {
	ctx, cancelCache := context.WithCancel(context.Background())
	c := &cache{
		mu:          sync.RWMutex{},
		items:       map[string]cacheEntry{},
		cancelCache: cancelCache,
	}
	c.reapLoop(ctx, interval)
	return c
}

func (c *cache) Add(key string, data []byte) {
	c.mu.Lock()
	c.items[key] = cacheEntry{
		val:       data,
		createdAt: time.Now(),
	}
	c.mu.Unlock()
}

func (c *cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.items[key]
	return entry.val, ok
}

func (c *cache) reapLoop(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				c.mu.Lock()
				for k, v := range c.items {
					elapsed := time.Since(v.createdAt)
					if elapsed > interval {
						delete(c.items, k)
					}
				}
				c.mu.Unlock()
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (c *cache) Cancel() {
	c.cancelCache()
}
