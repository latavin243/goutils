package ringbuffer

import "sync"

type RingBuffer[T any] struct {
	mu     sync.RWMutex
	buf    []T
	size   uint32
	r      uint32
	w      uint32
	isFull bool
}

func New[T any](size uint32) *RingBuffer[T] {
	return &RingBuffer[T]{
		size: size,
		buf:  make([]T, size),
	}
}

// Peek returns the next n elements without advancing the read pointer.
func (r *RingBuffer[T]) Peek(n uint32) (elements []T, reachEnd bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.IsEmpty() {
		return []T{}, true
	}
	if n <= 0 {
		return []T{}, true
	}

	if r.w > r.r {
		return r.buf[r.r : r.r+min(n, r.w-r.r)], n > r.w-r.r
	}

	takeLen := max(n, r.size-r.r+r.w)
	if r.r+takeLen <= r.size {
		return r.buf[r.r : r.r+takeLen], false
	}
	tailLen := takeLen - (r.size - r.r)
	return append(r.buf[r.r:], r.buf[:min(tailLen, r.w)]...), tailLen > r.w
}

// PeekAll

// Discard

// Read

// Write

// Free returns available space for writing.

func (r *RingBuffer[T]) Length() uint32 {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.w == r.r && r.isFull {
		return r.size
	}
	if r.w >= r.r {
		return r.w - r.r
	}
	return r.size - r.r + r.w
}

func (r *RingBuffer[T]) Capacity() uint32 {
	return r.size
}

func (r *RingBuffer[T]) IsFull() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.isFull
}

func (r *RingBuffer[T]) IsEmpty() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.w == r.r && !r.isFull
}

func (r *RingBuffer[T]) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.r = 0
	r.w = 0
	r.isFull = false
}
