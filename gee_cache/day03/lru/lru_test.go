package lru

import (
	"fmt"
	"testing"
)

type String string

func (d String) Len() int64 {
	return int64(len(d))
}

func TestNew(t *testing.T) {
	cache := New(20, func(key string, Value any) {
		fmt.Printf("evicted: %s\n", key)
	})
	cache.Add("1", String("111"))
	cache.Add("2", String("222"))
	cache.Add("3", String("333"))
	cache.Add("4", String("1234"))
	cache.Get("2")
	cache.Add("5", String("4444"))
	cache.Add("6", String("3333"))

	fmt.Println(cache.Len())
}
