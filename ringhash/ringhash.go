package ringhash

import (
	"sort"
	"strconv"
	"sync"

	"github.com/spaolacci/murmur3"
)

type hashFn func([]byte) uint32

type Ringhash struct {
	mu       sync.RWMutex
	hash     hashFn
	hashMap  map[int]string
	hashKeys []int
}

func New(hash hashFn) *Ringhash {
	if hash == nil {
		hash = murmur3.Sum32
	}
	return &Ringhash{
		hash:    hash,
		hashMap: make(map[int]string),
	}
}

func (r *Ringhash) IsEmpty() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.hashMap) == 0
}

// Add replicates virtual nodes for the input keys and put on the hash ring.
func (r *Ringhash) Add(replicas uint32, keys ...string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, key := range keys {
		for i := uint32(0); i < replicas; i++ {
			hashInt := int(r.hash([]byte(strconv.Itoa(int(i)) + key)))
			r.hashKeys = append(r.hashKeys, hashInt)
			sort.Slice(r.hashKeys, func(i, j int) bool { return r.hashKeys[i] < r.hashKeys[j] })
			r.hashMap[hashInt] = key
		}
	}
}

// Remove removes virtual nodes for the input key
func (r *Ringhash) Remove(key string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for k, v := range r.hashMap {
		if v == key {
			delete(r.hashMap, k)
		}
	}
}

// Get returns the key associated with the refKey specified
func (r *Ringhash) Get(refKey []byte) (string, bool) {
	if r.IsEmpty() {
		return "", false
	}

	hash := int(r.hash(refKey))
	r.mu.RLock()
	defer r.mu.RUnlock()
	idx := sort.Search(len(r.hashMap), func(i int) bool { return r.hashKeys[i] >= hash })
	// cycled back to the first replica if reached the end
	if idx == len(r.hashMap) {
		idx = 0
	}
	return r.hashMap[r.hashKeys[idx]], true
}
