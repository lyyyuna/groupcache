package lru

import "container/list"

// Cache is an LRU cache
type Cache struct {
	MaxEntries int
	OnEvicted  func(key Key, value interface{})
	cache      map[interface{}]*list.Element
	ll         *list.List
}

// Key is a value that can be comarable
type Key interface{}

type entry struct {
	key   Key
	value interface{}
}

// New creates a new Cache
// if maxEntries is 0, the cache has no limit and it's assumed
// that eviction is done by the caller,
func New(maxEntries int) *Cache {
	return &Cache{
		MaxEntries: maxEntries,
		ll:         list.New(),
		cache:      make(map[interface{}]*list.Element),
	}
}
