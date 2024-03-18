package iterop_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/latavin243/goutils/iterop"
)

func TestSliceFilter(t *testing.T) {
	raw := []int{1, 2, 3, 4, 5}
	res := iterop.SliceFilter(raw, func(i int) bool { return i%2 == 0 })
	assert.Equal(t, []int{2, 4}, res)
}

func TestSliceMap(t *testing.T) {
	raw := []int{1, 2, 3, 4, 5}
	res := iterop.SliceMap(raw, func(i int) int { return i * 2 })
	assert.Equal(t, []int{2, 4, 6, 8, 10}, res)
}

func TestSliceFlatMap(t *testing.T) {
	raw := []int{1, 2, 3, 4, 5}
	res := iterop.SliceFlatMap(raw, func(i int, collector func(int)) {
		collector(i * 2)
		collector(i * 3)
	})
	assert.Equal(t, []int{2, 3, 4, 6, 6, 9, 8, 12, 10, 15}, res)
}

func TestSliceReduce(t *testing.T) {
	raw := []int{1, 2, 3, 4, 5}
	res := iterop.SliceReduce(raw, 0, func(acc, i int) int { return acc + i })
	assert.Equal(t, 15, res)
}

func TestSliceChunk(t *testing.T) {
	raw := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	chunks := iterop.SliceChunk(raw, 3)
	assert.Equal(t, 4, len(chunks))
	assert.Equal(t, [][]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}, {"j"}}, chunks)
}

func TestSliceContains(t *testing.T) {
	raw := []int{1, 2, 3, 4, 5}
	assert.True(t, iterop.SliceContains(raw, []int{1, 2}))
	assert.False(t, iterop.SliceContains(raw, []int{1, 6}))
}

func TestSliceUnique(t *testing.T) {
	raw := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	res := iterop.SliceUnique(raw)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, res)
}

func TestSliceSub(t *testing.T) {
	raw := []int{1, 2, 3, 4, 5}
	res := iterop.SliceSub(raw, []int{1, 5})
	assert.Equal(t, []int{2, 3, 4}, res)
}
