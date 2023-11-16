package routinegroup_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/panjf2000/ants/v2"
	"github.com/stretchr/testify/assert"

	"github.com/latavin243/goutils/routinegroup"
)

type parallelableFunc func() error

func newAntsRoutineGroup(poolSize int, maxBlockingSize int) (routinegroup.Group, error) {
	pool, err := ants.NewPool(
		poolSize,
		ants.WithMaxBlockingTasks(maxBlockingSize),
	)
	if err != nil {
		return nil, err
	}
	return routinegroup.NewAntsGroup(pool), nil
}

func TestRoutineGroup(t *testing.T) {
	var err error
	a, b, c := 0, 0, 0
	group, err := newAntsRoutineGroup(1, 0)
	if err != nil {
		panic(err)
	}
	group.Submit(func() error {
		a = 1
		return nil
	})
	group.Submit(func() error {
		b = 2
		time.Sleep(time.Second)
		return nil
	})
	group.Submit(func() error {
		c = 3
		return nil
	})
	err = group.Wait()
	if err != nil {
		panic(err)
	}
	fmt.Printf("a=%d, b=%d, c=%d\t", a, b, c)
	assert.Equal(t, a, 1)
	assert.Equal(t, b, 2)
	assert.Equal(t, c, 3)
}

func TestRoutineGroupLargeAmount(t *testing.T) {
	var err error
	size := 100
	s := make([]int, size)
	var mu sync.Mutex
	group, err := newAntsRoutineGroup(1, 100)
	if err != nil {
		panic(err)
	}
	for i := 0; i < size; i++ {
		i := i
		group.Submit(func() error {
			mu.Lock()
			s[i] = i
			mu.Unlock()
			return nil
		})
	}
	err = group.Wait()
	if err != nil {
		panic(err)
	}
	for i := 0; i < size; i++ {
		assert.Equal(t, s[i], i)
	}
}
