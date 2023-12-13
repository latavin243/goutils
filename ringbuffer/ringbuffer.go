package ringbuffer

import "sync"

type Buffer[T any] interface {
	Length() uint32
	Capacity() uint32
	IsFull() bool
	IsEmpty() bool
	Reset()

	Write(elements []T) (successCnt uint32)
	Peek(n uint32) (elements []T, reachEnd bool)
	PeekAll() (elements []T)
	Read(n uint32) (elements []T, reachEnd bool)
	Remove(n uint32) (reachEnd bool)
}

type RingBuffer[T any] struct {
	mu     sync.RWMutex
	buf    []T
	size   uint32
	r      uint32 // read pointer
	w      uint32 // write pointer
	isFull bool
}

func New[T any](size uint32) Buffer[T] {
	return &RingBuffer[T]{
		size: size,
		buf:  make([]T, size),
	}
}

// Length returns the number of elements in the buffer
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

// Capacity returns the capacity of the buffer
func (r *RingBuffer[T]) Capacity() uint32 {
	return r.size
}

// IsFull returns true if the buffer is full
func (r *RingBuffer[T]) IsFull() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.isFull
}

// IsEmpty returns true if the buffer is empty
func (r *RingBuffer[T]) IsEmpty() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.isEmpty()
}

// Reset resets the buffer, but not clearing the elements (to be overwritten)
func (r *RingBuffer[T]) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.r = 0
	r.w = 0
	r.isFull = false
}

// Write writes elements to the buffer
// if successCnt < len(elements), input elements are not fully inserted
func (r *RingBuffer[T]) Write(elements []T) (successCnt uint32) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.IsFull() || len(elements) == 0 {
		return 0
	}
	inputLen := uint32(len(elements))

	if r.r < r.w {
		if inputLen <= r.size-r.w {
			copy(r.buf[r.w:], elements)
			r.w += inputLen
			if r.w == r.size {
				r.w = 0
				r.isFull = true
			}
			return inputLen
		}

		copy(r.buf[r.w:], elements[:r.size-r.w])
		copy(r.buf, elements[r.size-r.w:min(inputLen, r.size-r.w+r.r)])
		r.w = min(inputLen-(r.size-r.w), r.r)
		r.isFull = r.w == r.r
		return inputLen - r.w
	}

	if r.w+inputLen <= r.r {
		copy(r.buf[r.w:], elements)
		r.w += inputLen
		r.isFull = r.w == r.r
		return inputLen
	}

	copy(r.buf[r.w:], elements[:r.r-r.w])
	r.w = r.r
	r.isFull = true
	return r.r - r.w
}

// Peek returns the next n elements without advancing the read pointer
func (r *RingBuffer[T]) Peek(n uint32) (elements []T, reachEnd bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.peek(n)
}

// PeekAll returns all elements without advancing the read pointer
func (r *RingBuffer[T]) PeekAll() (elements []T) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.IsEmpty() {
		return []T{}
	}

	if r.w > r.r {
		return r.buf[r.r:r.w]
	}
	return append(r.buf[r.r:], r.buf[:r.w]...)
}

func (r *RingBuffer[T]) Read(n uint32) (elements []T, reachEnd bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	elements, reachEnd = r.peek(n)
	r.remove(n)
	return elements, reachEnd
}

func (r *RingBuffer[T]) Remove(n uint32) (reachEnd bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.remove(n)
	return r.isEmpty()
}

func (r *RingBuffer[T]) isEmpty() bool {
	return r.w == r.r && !r.isFull
}

func (r *RingBuffer[T]) peek(n uint32) (elements []T, reachEnd bool) {
	if r.IsEmpty() || n <= 0 {
		return []T{}, true
	}

	// if not across the conjunction of the ring, take straight forward
	if r.r+n <= r.size || r.w > r.r {
		return r.buf[r.r : r.r+min(n, r.w-r.r)], n > r.w-r.r
	}

	// take r.r -> end, then start -> tail
	tailLen := n - (r.size - r.r)
	return append(r.buf[r.r:], r.buf[:min(tailLen, r.w)]...), tailLen > r.w
}

func (r *RingBuffer[T]) remove(n uint32) {
	if r.IsEmpty() || n <= 0 {
		return
	}

	r.isFull = false
	if r.w > r.r || r.r+n <= r.size {
		r.r = min(r.r+n, r.w)
		if r.r == r.size {
			r.r = 0
		}
		return
	}

	tailLen := n - (r.size - r.r)
	r.r = min(tailLen, r.w)
}
