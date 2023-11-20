package roundrobin

import (
	"errors"
	"sync/atomic"
)

var (
	ErrEmptyCandidates = errors.New("empty candidates")
	ErrNotInitialized  = errors.New("not initialized")
)

type RoundRobin[T any] struct {
	candidates []T
	capacity   int32
	current    int32
}

func New[T any](candidates ...T) (*RoundRobin[T], error) {
	if len(candidates) == 0 {
		return nil, ErrEmptyCandidates
	}
	return &RoundRobin[T]{
		candidates: candidates,
		capacity:   int32(len(candidates)),
		current:    -1,
	}, nil
}

func (r *RoundRobin[T]) Next() (T, error) {
	var defaultVal T
	if r == nil {
		return defaultVal, ErrNotInitialized
	}
	if len(r.candidates) == 0 || r.capacity == 0 {
		return defaultVal, ErrEmptyCandidates
	}

	n := atomic.AddInt32(&r.current, 1)
	if n > r.capacity {
		atomic.AddInt32(&r.current, r.capacity)
	}
	return r.candidates[n%r.capacity], nil
}
