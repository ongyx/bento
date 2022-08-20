package ecs

// Component represents a component of type T in a world.
// The methods Get, Set, Insert, and Delete are wrappers; refer to the Table type for their documentation.
type Component[T any] struct {
	world *World
	table *Table[T]
	id    uint8
}

// Query returns a component from the world with type T.
// Register must be called before using Query.
func Query[T any](w *World) *Component[T] {
	e := w.tables[Type[T]()]
	return &Component[T]{
		world: w,
		table: e.table.(*Table[T]),
		id:    e.id,
	}
}

// Register adds the type T to the world and returns it as a component, allocating a table of (size) capacity as well.
// This panics if T has already been registered or the world is full (limit is 64 types).
func Register[T any](w *World, size int) *Component[T] {
	t := Type[T]()
	if _, ok := w.tables[t]; ok {
		panic("component: already registered")
	}

	// attempt to get new table id
	id := &w.tableID
	if *id >= 64 {
		panic("component: world is full")
	}
	*id++

	// add table to world
	tbl := NewTable[T](size)
	w.tables[t] = entry{tbl, *id}

	return &Component[T]{
		world: w,
		table: tbl,
		id:    *id,
	}
}

func (c *Component[T]) Get(e Entity) *T {
	return c.table.Get(e)
}

func (c *Component[T]) Set(e Entity, com T) {
	c.table.Set(e, com)
	c.world.entities[e].Set(c.id)
}

func (c *Component[T]) Insert(e Entity) *T {
	p := c.table.Insert(e)
	c.world.entities[e].Set(c.id)
	return p
}

func (c *Component[T]) Delete(e Entity) {
	c.table.Delete(e)
	c.world.entities[e].Clear(c.id)
}
