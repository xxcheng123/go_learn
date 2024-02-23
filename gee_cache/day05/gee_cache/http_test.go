package geecache

import (
	"fmt"
	"log"
	"testing"
)

var dbs = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func TestHttp(t *testing.T) {
	httpPool := NewHttpPool(":8080")
	NewGroup("tom", 1024, GetterFunc(func(key string) ([]byte, error) {
		log.Println("[SlowDB] search key", key)
		if v, ok := dbs[key]; ok {
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s not exist", key)
	}))
	if err := httpPool.Run(); err != nil {
		t.Error(err)
	}
}
