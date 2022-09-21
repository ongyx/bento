package sparse

import "sort"

// Map is a ordered map of K to V backed by a sparse set.
type Map[K comparable, V any] struct {
	keys   *Set[K]
	values []V
}

// NewMap creates a new map of (size) capacity.
func NewMap[K comparable, V any](size int) *Map[K, V] {
	return &Map[K, V]{
		keys:   NewSet[K](size),
		values: make([]V, 0, size),
	}
}

// Get returns a pointer to the value by key.
func (m *Map[K, V]) Get(key K) *V {
	return &m.values[m.keys.Index(key)]
}

// Insert inserts the key and value into the map.
func (m *Map[K, V]) Insert(key K, value V) {
	m.keys.Insert(key)
	m.values = append(m.values, value)
}

// Delete removes the value from the sparse set.
// This does not preserve order.
func (m *Map[K, V]) Delete(key K) {
	idx := m.keys.Index(key)
	size := len(m.values)

	// values are kept in sync with the dense index of the sparse set
	// replace the deleted value with the last entry
	m.values[idx] = m.values[size-1]
	m.values = m.values[:size]

	m.keys.Delete(key)
}

// Contains checks if the map has the key.
func (m *Map[K, V]) Contains(key K) bool {
	return m.keys.Index(key) >= 0
}

// Sort sorts the map by its keys.
func (m *Map[K, V]) Sort(fn func(a, b K) bool) {
	keys := m.keys.dense
	sort.Slice(m.values, func(i, j int) bool {
		return fn(keys[i], keys[j])
	})

	m.keys.Sort(fn)
}
