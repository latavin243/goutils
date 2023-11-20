package strcase_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/latavin243/goutils/strcase"
)

func TestCamelToSnake(t *testing.T) {
	testCases := []struct {
		input, expected string
	}{
		{"UserTotalCnt", "user_total_cnt"},
	}
	for _, tc := range testCases {
		got := strcase.CamelToSnake(tc.input)
		fmt.Printf("got: %s\n", got)
		assert.Equal(t, tc.expected, got)
	}
}

func TestSnakeToTitle(t *testing.T) {
	testCases := []struct {
		input, expected string
	}{
		{"user_total_cnt", "User Total Cnt"},
	}
	for _, tc := range testCases {
		got := strcase.SnakeToTitle(tc.input)
		fmt.Printf("got: %s\n", got)
		assert.Equal(t, tc.expected, got)
	}
}
