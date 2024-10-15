package reflectutil_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/latavin243/goutils/reflectutil"
)

type MyStruct struct {
	Name  string
	Value int
}

func TestIsSameType(t *testing.T) {
	type MyStructA struct {
		Name string
		Rate float64
	}
	type MyStructB struct {
		Name  string
		Value int
	}

	s := MyStruct{}
	a := MyStructA{}
	assert.False(t, reflectutil.SameType(s, a))

	s1 := &MyStruct{Name: "s1", Value: 1}
	s2 := &MyStruct{Name: "s2", Value: 2}
	assert.True(t, reflectutil.SameType(s1, s2))

	sp := &MyStruct{}
	ap := &MyStructA{}
	bp := &MyStructB{}
	assert.False(t, reflectutil.SameType(sp, ap))
	assert.False(t, reflectutil.SameType(sp, bp))

	p1 := &MyStruct{Name: "p1", Value: 1}
	p2 := &MyStruct{Name: "p2", Value: 2}
	assert.True(t, reflectutil.SameType(p1, p2))
}

func TestIsStructPtr(t *testing.T) {
	a := 1
	assert.False(t, reflectutil.IsStructPtr(a))

	aPtr := &a
	assert.False(t, reflectutil.IsStructPtr(aPtr))

	s := MyStruct{}
	assert.False(t, reflectutil.IsStructPtr(s))

	sPtr := &MyStruct{}
	assert.True(t, reflectutil.IsStructPtr(sPtr))
}
