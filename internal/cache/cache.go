package cache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mu      *sync.RWMutex
}

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{
		entries: make(map[string]cacheEntry),
		mu:      &sync.RWMutex{},
	}

	go newCache.reapLoop(interval)
	go newCache.reapRandomKey()
	return &newCache
}

func (cache *Cache) Add(key string, value []byte) {
	cache.mu.Lock()

	defer cache.mu.Unlock()

	cache.entries[key] = cacheEntry{
		CreatedAt: time.Now(),
		Val:       value,
	}
}

func (cache *Cache) Get(key string) (cacheEntry, bool) {

	cache.mu.Lock()

	defer cache.mu.Unlock()

	v, ok := cache.entries[key]

	if !ok {
		return cacheEntry{}, false
	}

	return v, ok

}

func (cache *Cache) reapLoop(interval time.Duration) {

	//time.Since(cache.entries[])

	ticker := time.NewTicker(interval)

	for range ticker.C {
		for k, v := range cache.entries {
			cache.mu.Lock()
			startTime := time.Since(v.CreatedAt)
			if startTime > interval {
				delete(cache.entries, k)
			}
		}
		cache.mu.Unlock()

	}
}

// is this even a good idea?
/*
Use with caution. If  Usr.Puuid is deleted while the corresponding entries struct
has not, behavior may be strange?

Ideally it should be fine, because we should always check if it is in cache first
otherwise we send the get request
*/
func (cache *Cache) reapRandomKey() {
	cache.mu.Lock()

	defer cache.mu.Unlock()
	randomKey := ""
	if len(cache.entries) > 15 {
		for k := range cache.entries {
			randomKey = k
		}
		v, ok := cache.entries[randomKey]
		if !ok {
			fmt.Println("Key does not exist in reapRandomKey")
		}

		timeStamp := v.CreatedAt

		timeSince := time.Since(timeStamp)

		if timeSince > 20000000000 {
			delete(cache.entries, randomKey)
		}

	}
}
