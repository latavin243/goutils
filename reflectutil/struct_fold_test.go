package reflectutil_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/latavin243/goutils/reflectutil"
)

func TestStructFoldAdd(t *testing.T) {
	type MyStruct struct {
		A int
		B float64
		C int `fold:"skip"`
		D int `fold:"zero"`
		E int `fold:"zero"`
		F int `fold:"first"`
	}

	res := &MyStruct{B: 1.1, C: 10, D: 20, E: 0, F: 10}
	item1 := &MyStruct{A: 1, B: 2.2, C: 3, D: 10, F: 20}
	item2 := &MyStruct{A: 3, B: 4.3, C: 5, E: 10}
	expected := &MyStruct{A: 4, B: 7.6, C: 10, D: 0, E: 0, F: 10}

	err := reflectutil.StructFoldSum(res, item1, item2)
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
	fmt.Printf("res=%+v\n", res)
}
