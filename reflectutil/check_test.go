package reflectutil_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/latavin243/goutils/reflectutil"
)

func TestIsSameType(t *testing.T) {
	type MyStructA struct {
		Name  string
		Value int
	}
	type MyStructB struct {
		Name string
		Rate float64
	}
	type MyStructC struct {
		Name  string
		Value int
	}

	a := MyStructA{}
	b := MyStructB{}
	assert.False(t, reflectutil.SameType(a, b))

	pa := &MyStructA{}
	pb := &MyStructB{}
	pc := &MyStructC{}
	assert.False(t, reflectutil.SameType(pa, pb))
	assert.False(t, reflectutil.SameType(pa, pc))

	pa1 := &MyStructA{Name: "a1", Value: 1}
	pa2 := &MyStructA{Name: "a2", Value: 2}
	assert.True(t, reflectutil.SameType(pa1, pa2))
}
