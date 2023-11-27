package set_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/latavin243/goutils/set"
)

func TestSet(t *testing.T) {
	// test New(), Empty(), Len()
	s := set.New[int]()
	assert.True(t, s.Empty())
	assert.Equal(t, 0, s.Len())

	// test Contains(), Len()
	s.Add(1) // {1}
	assert.False(t, s.Empty())
	assert.True(t, s.Contains(1))
	assert.False(t, s.Contains(2))
	assert.Equal(t, 1, s.Len())

	// test Remove()
	s.Add(2)
	s.Remove(1) // {2}
	assert.True(t, s.Contains(2))
	assert.False(t, s.Contains(1))

	// test Clear()
	s.Clear() // {}
	assert.True(t, s.Empty())

	// test ToSlice(), ToUnsortedSlice()
	{
		s := set.New(2, 3, 1) // {2,3,1}
		assert.Equal(t, []int{1, 2, 3}, s.ToSlice(func(ss []int) {
			sort.Slice(ss, func(i, j int) bool { return ss[i] < ss[j] })
		}))
		fmt.Printf("unsorted slice: %+v\n", s.ToUnsortedSlice())
	}

	// test Equal()
	{
		s1 := set.New(1, 2, 3)
		s2 := set.New(3, 1, 2, 2)
		s3 := set.New(2, 3, 4)
		assert.True(t, s1.Equal(s2))
		assert.False(t, s1.Equal(s3))
	}

	// test Clone()
	{
		s1 := set.New(1, 2, 3)
		s2 := s1.Clone()
		assert.True(t, s1.Equal(s2))
	}

	// test Union(), Intersect(), Difference()
	{
		s1 := set.New(1, 2, 3)
		s2 := set.New(2, 3, 4)
		assert.True(t, s1.Union(s2).Equal(set.New(1, 2, 3, 4)))
		assert.True(t, s1.Intersect(s2).Equal(set.New(2, 3)))
		assert.True(t, s1.Difference(s2).Equal(set.New(1)))
	}
}
