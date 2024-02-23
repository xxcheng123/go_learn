package lru

import "container/list"

type Cache struct {
	maxBytes  int64
	nBytes    int64
	ll        *list.List
	cache     map[string]*list.Element
	OnEvicted func(key string, Value any)
}

type Value interface {
	Len() int64
}

type entry struct {
	key   string
	value Value
}

func (c *Cache) Get(key string) (value Value, ok bool) {
	if v, ok := c.cache[key]; ok {
		c.ll.MoveToFront(v)
		return v.Value.(*entry).value, true
	}
	return
}

func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nBytes -= kv.value.Len() + int64(len(kv.key))
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

func (c *Cache) Add(key string, value Value) {
	//判断是否存在
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nBytes += value.Len() - kv.value.Len()
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key: key, value: value})
		c.cache[key] = ele
		c.nBytes += value.Len() + int64(len(key))
	}
	//判断是否溢出
	if c.maxBytes != 0 && c.nBytes > c.maxBytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Len() int64 {
	return int64(c.ll.Len())
}

func New(maxBytes int64, OnEvicted func(key string, Value any)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: OnEvicted,
	}
}
