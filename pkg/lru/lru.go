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

// Add adds a value to the cache
func (c *Cache) Add(key Key, value interface{}) {
	if c.cache == nil {
		c.cache = make(map[interface{}]*list.Element)
		c.ll = list.New()
	}

	if ee, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ee)
		ee.Value.(*entry).value = value
		return
	}

	ele := c.ll.PushFront(&entry{key, value})
	c.cache[key] = ele
	if c.MaxEntries != 0 && c.ll.Len() > c.MaxEntries {
		c.RemoveOldest()
	}
}

// RemoveOldest removes the oldest item from the cache
func (c *Cache) RemoveOldest() {
	if c.cache == nil {
		return
	}

	ele := c.ll.Back()
	if ele != nil {
		c.removeElement(ele)
	}
}

func (c *Cache) removeElement(e *list.Element) {
	c.ll.Remove(e)
	kv := e.Value.(*entry)
	delete(c.cache, kv.key)

	if c.OnEvicted != nil {
		c.OnEvicted(kv.key, kv.value)
	}
}
