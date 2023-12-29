package iterop_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/latavin243/goutils/iterop"
)

func TestMapOp(t *testing.T) {
	m := map[string]int{"a": 2, "b": 1, "c": 3}

	expectedKeys := []string{"a", "b", "c"}
	keys := iterop.MapKeys(m, func(keys []string) {
		sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	})
	fmt.Printf("keys: %v\n", keys)
	assert.Equal(t, expectedKeys, keys)

	expectedSortedVals := []int{1, 2, 3}
	vals := iterop.MapValues(m, func(vals []int) {
		sort.Slice(vals, func(i, j int) bool { return vals[i] < vals[j] })
	})
	fmt.Printf("vals: %v\n", vals)
	assert.Equal(t, expectedSortedVals, vals)
}
