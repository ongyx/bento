package ecs

type table interface {
	Delete(e Entity)
}

// Table is a map of entities to indices into a dense slice of components of type T.
// It allows fast insertion and removal, but the order of components are not guaranteed.
type Table[T any] struct {
	entities   map[Entity]int
	components []T
}

// NewTable creates a new map with a capacity of (size) components.
func NewTable[T any](size int) *Table[T] {
	return &Table[T]{
		entities:   make(map[Entity]int, size),
		components: make([]T, size),
	}
}

// Get returns a reference to the component indexed by the entity.
// This panics if the entity is not in the table.
func (t *Table[T]) Get(e Entity) *T {
	return &t.components[t.entities[e]]
}

// Set inserts an entity-component pair into the table.
func (t *Table[T]) Set(e Entity, c T) {
	*t.Insert(e) = c
}

// Insert inserts the entity if it does not exist, and returns a reference to the component.
// The value of the component is guaranteed to be the zero value of T.
func (t *Table[T]) Insert(e Entity) *T {
	idx, ok := t.entities[e]

	if !ok {
		var zv T

		// after appending, the old length is now the last index.
		idx = len(t.components)

		t.entities[e] = idx
		t.components = append(t.components, zv)
	}

	return &t.components[idx]
}

// Delete removes the component by entity, swapping the index of the last element in place.
func (t *Table[T]) Delete(e Entity) {
	idx := t.entities[e]
	last := len(t.components) - 1

	t.components[idx] = t.components[last]
	t.components = t.components[:len(t.components)-1]
	delete(t.entities, e)
}
