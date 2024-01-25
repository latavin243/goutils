package fileutil_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/latavin243/goutils/fileutil"
)

func TestRealPath(t *testing.T) {
	type testCase struct {
		input, expected string
	}

	curDir, _ := os.Getwd()

	for _, tc := range []testCase{
		{"a/b/c.txt", curDir + "/a/b/c.txt"},
		{"/a/b/c.txt", "/a/b/c.txt"},
	} {
		res, err := RealPath(tc.input)
		assert.NoError(t, err)
		assert.Equal(t, tc.expected, res)
	}
}

func TestBasename(t *testing.T) {
	assert.Equal(t, Basename("/path/to/file.txt"), "file.txt")
}

func TestDir(t *testing.T) {
	assert.Equal(t, Dir("/path/to/file.txt"), "/path/to")
}

func TestExt(t *testing.T) {
	assert.Equal(t, Ext("/path/to/file.txt"), ".txt")
}
