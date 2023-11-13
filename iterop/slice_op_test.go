package iterop_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/latavin243/goutils/iterop"
)

func TestSliceChunk(t *testing.T) {
	chunks := iterop.SliceChunk([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}, 3)
	assert.Equal(t, 4, len(chunks))
	for _, chunk := range chunks {
		fmt.Printf("%+v\n", chunk)
	}
}
