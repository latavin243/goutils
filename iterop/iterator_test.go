package iterop_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/latavin243/goutils/iterop"
)

func TestSliceIter(t *testing.T) {
	last := 5
	s := make([]int, 0, last)
	for i := 0; i <= last; i++ {
		s = append(s, i)
	}

	iter := iterop.SliceToIter(s)
	for i := 0; i <= last; i++ {
		expectedHasNext := i < last
		assert.Equal(t, expectedHasNext, iter.HasNext())

		elem, hasHext := iter.Next()
		fmt.Printf("elem: %v, hasNext: %v\n", elem, hasHext)
		assert.Equal(t, i, elem)

		assert.Equal(t, expectedHasNext, hasHext)
	}
}
