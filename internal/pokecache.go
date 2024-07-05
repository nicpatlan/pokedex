package pokedexAPI

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	storage map[string]cacheEntry
	mu      sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	var cache *Cache = &Cache{storage: make(map[string]cacheEntry)}
	ticker := time.NewTicker(interval)
	go cache.reapLoop(ticker, interval)
	return cache
}

func (cache *Cache) Add(key string, value []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.storage[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	entry, ok := cache.storage[key]
	if !ok {
		return []byte(""), false
	}
	cache.storage[key] = cacheEntry{
		createdAt: time.Now(),
		val:       entry.val,
	}
	return entry.val, true
}

func (cache *Cache) reapLoop(ticker *time.Ticker, interval time.Duration) {
	defer ticker.Stop()
	for {
		t := <-ticker.C
		cache.mu.Lock()
		for key := range cache.storage {
			if t.Sub(cache.storage[key].createdAt) > interval {
				delete(cache.storage, key)
			}
		}
		cache.mu.Unlock()
	}
}
