package ecs

// Component represents a component pool of type T in a world.
type Component[T any] struct {
	world *World
	pool  Pool[T]
	id    uint8
}

// Get returns the component data by entity.
func (c *Component[T]) Get(e Entity) *T {
	return c.pool.Get(e)
}

// Insert inserts the entity and component data into the pool.
func (c *Component[T]) Insert(e Entity, data T) {
	c.pool.Insert(e, data)
	c.signature(e).Set(c.id)
}

// Delete deletes the component data by entity from the pool.
func (c *Component[T]) Delete(e Entity) {
	c.pool.Delete(e)
	c.signature(e).Clear(c.id)
}

// Sort sorts the component pool with the function.
func (c *Component[T]) Sort(fn func(a, b T) bool) {
	c.pool.Sort(func(a, b Entity) bool {
		return fn(*c.pool.Get(a), *c.pool.Get(b))
	})
}

func (c *Component[T]) signature(e Entity) *Signature {
	return &c.world.entities[e.id].sig
}

// Query returns a component from the world with type T.
// Register must be called before using Query.
func Query[T any](w *World) *Component[T] {
	id := w.poolID[Type[T]()]
	return &Component[T]{
		world: w,
		pool:  w.pools[id].(Pool[T]),
		id:    id,
	}
}

// Register adds the type T to the world and returns it as a component, allocating a pool of (size) capacity as well.
// This panics if T has already been registered or the world is full (limit is 64 types).
func Register[T any](w *World, size int) *Component[T] {
	t := Type[T]()
	if _, ok := w.poolID[t]; ok {
		panic("component: already registered")
	}

	// create new pool id
	id := uint8(len(w.pools))
	if id >= 64 {
		panic("component: world is full")
	}

	logger.Printf("registering %s with id %d\n", t.t.Name(), id)

	// add pool to world
	p := NewPool[T](size)
	w.pools = append(w.pools, p)
	w.poolID[t] = id

	return &Component[T]{
		world: w,
		pool:  p,
		id:    id,
	}
}
