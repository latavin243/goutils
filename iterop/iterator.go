package iterop

import "context"

type Iterator[T any] interface {
	HasNext() bool
	Next() (elem T, hasNext bool)
}

func IterToSlice[T any](iter Iterator[T]) []T {
	var s []T
	for elem, hasNext := iter.Next(); hasNext; elem, hasNext = iter.Next() {
		s = append(s, elem)
	}
	return s
}

func IterToChan[T any](ctx context.Context, iter Iterator[T], buffer int) <-chan T {
	ch := make(chan T, buffer)
	go func() {
		defer close(ch)
		for elem, hasNext := iter.Next(); hasNext; elem, hasNext = iter.Next() {
			select {
			case ch <- elem:
			case <-ctx.Done():
				return
			}
		}
	}()
	return ch
}

// sliceIter

type sliceIter[T any] struct {
	s   []T
	cur int
}

func (iter *sliceIter[T]) HasNext() bool {
	return iter.cur < len(iter.s)-1
}

func (iter *sliceIter[T]) Next() (elem T, hasNext bool) {
	return iter.s[iter.cur], iter.cur < len(iter.s)-1
}

func SliceToIter[T any](s []T) Iterator[T] {
	return &sliceIter[T]{s, 0}
}

// chanIter

type chanIter[T any] struct {
	ch <-chan T
}

func (iter *chanIter[T]) HasNext() bool {
	return len(iter.ch) == 0
}

func (iter *chanIter[T]) Next() (elem T, hasNext bool) {
	elem, hasNext = <-iter.ch
	return elem, hasNext
}

func ChanToIter[T any](ch <-chan T) Iterator[T] {
	return &chanIter[T]{ch}
}
