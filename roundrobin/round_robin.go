package roundrobin

import (
	"errors"
	"sync/atomic"
)

var ErrEmptyCandidates = errors.New("empty candidates")

type RoundRobin[T any] interface {
	Next() T
}

type roundRobinImpl[T any] struct {
	candidates []T
	capacity   int32
	current    int32
}

func New[T any](candidates ...T) (RoundRobin[T], error) {
	if len(candidates) == 0 {
		return nil, ErrEmptyCandidates
	}
	return &roundRobinImpl[T]{
		candidates: candidates,
		capacity:   int32(len(candidates)),
		current:    -1,
	}, nil
}

func (r *roundRobinImpl[T]) Next() T {
	var defaultVal T
	if r == nil {
		return defaultVal
	}

	n := atomic.AddInt32(&r.current, 1)
	if n > r.capacity {
		atomic.AddInt32(&r.current, r.capacity)
	}
	return r.candidates[n%r.capacity]
}
