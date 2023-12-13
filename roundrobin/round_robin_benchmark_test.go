package roundrobin_test

import (
	"fmt"
	"testing"

	"github.com/latavin243/goutils/roundrobin"
)

func BenchmarkRoundRobinSync(b *testing.B) {
	resources := []*struct {
		id   int
		name string
	}{
		{1, "resource-1"},
		{2, "resource-2"},
		{3, "resource-3"},
		{4, "resource-4"},
		{5, "resource-5"},
		{6, "resource-6"},
		{7, "resource-7"},
	}

	for i := 1; i < len(resources)+1; i++ {
		b.Run(fmt.Sprintf("RoundRobinSize(%d)", i), func(b *testing.B) {
			rr, err := roundrobin.New(resources[:i]...)
			if err != nil {
				b.Fatal(err)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				rr.Next()
			}
		})
	}
}

func BenchmarkRoundRobinASync(b *testing.B) {
	resources := []*struct {
		id   int
		name string
	}{
		{1, "resource-1"},
		{2, "resource-2"},
		{3, "resource-3"},
		{4, "resource-4"},
		{5, "resource-5"},
		{6, "resource-6"},
		{7, "resource-7"},
	}

	for i := 1; i < len(resources)+1; i++ {
		b.Run(fmt.Sprintf("RoundRobinSize(%d)", i), func(b *testing.B) {
			rr, err := roundrobin.New(resources[:i]...)
			if err != nil {
				b.Fatal(err)
			}
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					rr.Next()
				}
			})
		})
	}
}
