package set

type Set[T comparable] struct {
	m map[T]struct{}
}

func New[T comparable](elems ...T) *Set[T] {
	s := &Set[T]{m: make(map[T]struct{})}
	s.Add(elems...)
	return s
}

func (s *Set[T]) Empty() bool {
	return len(s.m) == 0
}

func (s *Set[T]) Add(elems ...T) {
	for _, elem := range elems {
		s.m[elem] = struct{}{}
	}
}

func (s *Set[T]) Remove(elems ...T) {
	if s == nil {
		return
	}
	for _, elem := range elems {
		delete(s.m, elem)
	}
}

func (s *Set[T]) Contains(elem T) bool {
	if s == nil {
		return false
	}
	_, ok := s.m[elem]
	return ok
}

func (s *Set[T]) Len() int {
	if s == nil {
		return 0
	}
	return len(s.m)
}

func (s *Set[T]) Clear() {
	if s == nil {
		return
	}
	s.m = make(map[T]struct{})
}

func (s *Set[T]) ToSlice(sortBy func(s []T)) []T {
	elems := make([]T, 0, len(s.m))
	for elem := range s.m {
		elems = append(elems, elem)
	}
	if sortBy != nil {
		sortBy(elems)
	}
	return elems
}

func (s *Set[T]) ToUnsortedSlice() []T {
	if s == nil {
		return nil
	}
	elems := make([]T, 0, len(s.m))
	for elem := range s.m {
		elems = append(elems, elem)
	}
	return elems
}

func (s *Set[T]) Clone() *Set[T] {
	if s == nil {
		return &Set[T]{}
	}
	clone := &Set[T]{m: make(map[T]struct{}, len(s.m))}
	for elem := range s.m {
		clone.m[elem] = struct{}{}
	}
	return clone
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	if s == nil {
		return other.Clone()
	}
	union := s.Clone()
	for elem := range other.m {
		union.m[elem] = struct{}{}
	}
	return union
}

func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	if s == nil {
		return &Set[T]{}
	}
	intersect := &Set[T]{m: make(map[T]struct{})}
	for elem := range s.m {
		if _, ok := other.m[elem]; ok {
			intersect.m[elem] = struct{}{}
		}
	}
	return intersect
}

func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	if s == nil {
		return other.Clone()
	}
	difference := &Set[T]{m: make(map[T]struct{})}
	for elem := range s.m {
		if _, ok := other.m[elem]; !ok {
			difference.m[elem] = struct{}{}
		}
	}
	return difference
}

func (s *Set[T]) Equal(other *Set[T]) bool {
	if s == nil && other == nil {
		return true
	}
	if s == nil || other == nil {
		return false
	}
	if len(s.m) != len(other.m) {
		return false
	}
	for elem := range s.m {
		if _, ok := other.m[elem]; !ok {
			return false
		}
	}
	return true
}
