package reflectutil_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/latavin243/goutils/reflectutil"
)

func TestStructFoldAdd(t *testing.T) {
	type MyStruct struct {
		Name string `fold:"first"`

		FloatSum float64 `fold:"num/sum"`
		IntSum   int     `fold:"num/sum"`
		IntSkip  int     `fold:"skip"`
		IntZero1 int     `fold:"num/zero"`
		IntZero2 int     `fold:"num/zero"`
		IntFirst int     `fold:"first"`
	}

	res := &MyStruct{
		Name: "name", FloatSum: 1.1,
		IntSkip: 10, IntZero1: 20, IntZero2: 0, IntFirst: 10,
	}
	item1 := &MyStruct{
		FloatSum: 2.2,
		IntSum:   1, IntSkip: 3, IntZero1: 10, IntFirst: 20,
	}
	item2 := &MyStruct{
		FloatSum: 4.3,
		IntSum:   3, IntSkip: 5, IntZero2: 10,
	}
	expected := &MyStruct{
		Name: "name", FloatSum: 7.6,
		IntSum: 4, IntSkip: 10, IntZero1: 0, IntZero2: 0, IntFirst: 10,
	}

	err := reflectutil.StructFold(res, item1, item2)
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
	fmt.Printf("res=%+v\n", res)
}
