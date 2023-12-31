package ringhash_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/latavin243/goutils/ringhash"
)

func TestRinghash(t *testing.T) {
	rh := ringhash.New(func(key []byte) uint32 {
		i, err := strconv.Atoi(string(key))
		if err != nil {
			panic(err)
		}
		return uint32(i)
	})

	// add 2, 4, 6, 12, 14, 16, 22, 24, 26
	rh.Add(3, "6", "4", "2")

	testCases := map[string]string{
		"2":  "2",
		"11": "2",
		"23": "4",
		"27": "2",
	}
	for k, expected := range testCases {
		res, ok := rh.Get([]byte(k))
		assert.True(t, ok)
		assert.Equal(t, expected, res)
	}

	// add 8, 18, 28
	rh.Add(3, "8")

	// 27 should now map to 8
	testCases["27"] = "8"
	for k, expected := range testCases {
		res, ok := rh.Get([]byte(k))
		assert.True(t, ok)
		assert.Equal(t, expected, res)
	}
}
