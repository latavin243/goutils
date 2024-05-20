package pool

import (
	"errors"
	"sync"
)

var (
	ErrClosed = errors.New("pool is closed")
	ErrEmpty  = errors.New("pool is empty")
	ErrFull   = errors.New("pool is full")
)

type (
	Pool[T any] interface {
		Get() (T, error)
		Close()
		Len() int
	}

	RecyclePool[T any] interface {
		Pool[T]
		Put(T) error
	}
)

// implementation

type SyncPool[T any] struct {
	*sync.Pool
	cnt uint
}

func NewSyncPool[T any](new func() T) RecyclePool[T] {
	newAny := func() any {
		return new()
	}
	return &SyncPool[T]{
		Pool: &sync.Pool{New: newAny},
		cnt:  0,
	}
}

func (p *SyncPool[T]) Get() (T, error) {
	if p.Pool == nil {
		var res T
		return res, ErrClosed
	}
	return p.Pool.Get().(T), nil
}

func (p *SyncPool[T]) Put(x T) error {
	if p.Pool == nil {
		return ErrClosed
	}
	p.cnt++
	p.Pool.Put(x)
	return nil
}

func (p *SyncPool[T]) Close() {
	p.cnt = 0
	p.Pool = nil
}

func (p *SyncPool[T]) Len() int {
	return int(p.cnt)
}
