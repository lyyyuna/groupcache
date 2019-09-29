package consistenhash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// Hash function type
type Hash func(data []byte) uint32

// Map structure
type Map struct {
	replicas int            // 每个 key 的副本数量
	hash     Hash           // 哈希函数
	keys     []int          // 每一个哈希点, keep it sorted
	hashMap  map[int]string // 哈希环上的一个点到服务器名的映射
}

// New a map consistentmap
func New(replicas int) *Map {
	m := &Map{
		replicas: replicas,
		hash:     crc32.ChecksumIEEE,
		hashMap:  make(map[int]string),
	}

	return m
}

// Add some keys to the hash
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}

	sort.Ints(m.keys)
}

// Get one value from hash
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))

	// 二分查找，即顺时针在哈希环上离这个key 最近的一个 服务器或服务器副本
	i := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	if i == len(m.keys) {
		i = 0
	}

	return m.hashMap[m.keys[i]]
}
