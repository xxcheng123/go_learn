package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32

type Map struct {
	hash     Hash
	replicas int
	// sort 没有uint32的类型，所以用int
	keys    []int
	hashMap map[int]string
}

func New(replicas int, hash Hash) *Map {
	m := &Map{
		hash:     hash,
		replicas: replicas,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			k := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.hashMap[k] = key
			m.keys = append(m.keys, k)
		}
	}
	sort.Ints(m.keys)
}

func (m *Map) Get(key string) string {
	//计算hash值
	hash := int(m.hash([]byte(key)))
	//顺时针获取第一个
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	//返回
	return m.hashMap[m.keys[(idx)%len(m.keys)]]
}
