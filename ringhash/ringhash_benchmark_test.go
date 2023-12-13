package ringhash_test

import (
	"fmt"
	"testing"

	"github.com/latavin243/goutils/ringhash"
)

func BenchmarkGet8(b *testing.B)   { benchmarkGet(b, 8) }
func BenchmarkGet32(b *testing.B)  { benchmarkGet(b, 32) }
func BenchmarkGet128(b *testing.B) { benchmarkGet(b, 128) }
func BenchmarkGet512(b *testing.B) { benchmarkGet(b, 512) }

func benchmarkGet(b *testing.B, shards int) {
	rh := ringhash.New(nil)
	var buckets []string
	for i := 0; i < shards; i++ {
		buckets = append(buckets, fmt.Sprintf("shard-%d", i))
	}

	rh.Add(50, buckets...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rh.Get([]byte(buckets[i&(shards-1)]))
	}
}
