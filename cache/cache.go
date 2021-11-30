package cache

import (
	"container/list"
)

type Callback func(string, Value)

type Value interface {
	Len() int
}

type Entry struct {
	key string
	value Value
}

func (entry *Entry) len() int {
	return len(entry.key) + entry.value.Len()
}

type Cache struct {
	maxBytes int
	nBytes int
	ll *list.List
	cache map[string]*list.Element

	OnEvicted func(key string, value Value)
}

func New(maxBytes int, onEvicted Callback) *Cache {
	return &Cache{
		maxBytes: maxBytes,
		ll: list.New(),
		cache: make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c *Cache) Get(key string) (Value, bool) {
	if elem, ok := c.cache[key]; ok {
		c.ll.MoveToFront(elem)
		kv := elem.Value.(*Entry)

		return kv.value, true
	}
	return nil, false
}

func (c *Cache) RemoveOldest() {
	elem := c.ll.Back()
	if elem == nil {
		return
	}
	c.ll.Remove(elem)
	kv := elem.Value.(*Entry)
	delete(c.cache, kv.key)
	c.nBytes -= kv.len()
	if c.OnEvicted != nil {
		c.OnEvicted(kv.key, kv.value)
	}
}

func (c *Cache) Add(key string, value Value) {
	if elem, ok := c.cache[key]; ok {
		c.ll.MoveToFront(elem)
		kv := elem.Value.(*Entry)
		c.nBytes += value.Len() - kv.value.Len()
		kv.value = value
	}else {
		elem := c.ll.PushFront(&Entry{key: key, value: value})
		c.cache[key] = elem
		c.nBytes += len(key) + value.Len()
	}
	for c.maxBytes !=0 && c.nBytes > c.maxBytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}