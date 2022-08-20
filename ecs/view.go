package ecs

// View is a cached filter that can be used by systems to search for entities with specific components.
type View struct {
	world *World
	sig   Signature
}

// NewView creates a new view with the component types.
func NewView(w *World, types ...TypeID) *View {
	v := &View{world: w}
	for _, t := range types {
		v.sig.Set(w.tables[t].id)
	}
	return v
}

// Each calls the function for each entity that matches the view's signature.
// Entities are not guaranteed to be in order; if you want to sort them before iteration use Filter.
func (v *View) Each(f func(e Entity)) {
	for e, sig := range v.world.entities {
		if sig.Contains(v.sig) {
			f(e)
		}
	}
}

// Filter appends all entities that match the view's signature to a buffer and returns it.
// The buffer may be nil, and can be reused across calls to Filter to reduce allocation.
func (v *View) Filter(buf []Entity) []Entity {
	// NOTE(ongyx): this avoids alloc if the buffer already has an underlying array
	buf = buf[:0]

	for e, sig := range v.world.entities {
		if sig.Contains(v.sig) {
			buf = append(buf, e)
		}
	}

	return buf
}
