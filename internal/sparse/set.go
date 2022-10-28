package sparse

import "sort"

// Set is a sparse set of type T.
// A sparse set is composed of a sparse array which maps indices to a dense array, and elements in the dense array are packed with no gaps.
// Iteration of the dense array is not guaranteed to be in order of insertion.
type Set[T comparable] struct {
	sparse map[T]int
	dense  []T
}

// NewSet creates a new sparse set of (size) capacity.
func NewSet[T comparable](size int) *Set[T] {
	return &Set[T]{
		sparse: make(map[T]int, size),
		dense:  make([]T, 0, size),
	}
}

// Index returns the dense index of the value.
// If the value is not in the sparse set, idx is -1 and ok is false.
func (s *Set[T]) Index(value T) (idx int, ok bool) {
	idx, ok = s.sparse[value]
	if !ok {
		// returning a negative index will always be out of bounds on purpose.
		idx = -1
	}

	return
}

// Insert inserts the value into the sparse set.
func (s *Set[T]) Insert(value T) {
	s.sparse[value] = len(s.dense)
	s.dense = append(s.dense, value)
}

// Delete removes the value from the sparse set.
// This does not preserve order.
func (s *Set[T]) Delete(value T) {
	size := len(s.dense)
	last := s.dense[size-1]

	// swap the indices of the value to be deleted and the last value in dense
	s.sparse[last] = s.sparse[value]

	// remove entry for value in sparse and reslice dense
	delete(s.sparse, value)
	s.dense = s.dense[:size]
}

// Dense returns the dense array of the set.
// The dense array should be treated as read-only (for iteration).
func (s *Set[T]) Dense() []T {
	return s.dense
}

// Sort sorts the dense array by its values.
func (s *Set[T]) Sort(fn func(a, b T) bool) {
	sort.Slice(s.dense, func(i, j int) bool {
		return fn(s.dense[i], s.dense[j])
	})

	for i, v := range s.dense {
		s.sparse[v] = i
	}
}
