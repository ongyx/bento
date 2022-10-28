package ecs

import "github.com/ongyx/bento/internal/sparse"

const viewSize = 10

// View is a cached filter that can be used by systems to search for entities with specific components.
type View struct {
	sig      *Signature
	world    *World
	entities *sparse.Set[Entity]
}

// NewView creates a new view with component types to match the entities against.
func NewView(world *World, types ...TypeID) *View {
	sig := new(Signature)

	for _, t := range types {
		sig.Set(world.poolID[t])
	}

	v := &View{sig, world, sparse.NewSet[Entity](viewSize)}
	v.Cache()

	return v
}

// Cache searches for entities which matches the view's filter.
func (v *View) Cache() {
	for _, e := range v.world.entities {
		// the entity must have the view's signature
		// (and therefore the components of interest),
		// not the other way round.
		if e.sig.Contains(*v.sig) {
			v.entities.Insert(e.entity)
		}
	}
}

// Sort sorts the entities in the view.
func (v *View) Sort(fn func(a, b Entity) bool) {
	v.entities.Sort(fn)
}

// Each iterates over the view's entities with the function fn.
func (v *View) Each(fn func(e Entity)) {
	for _, e := range v.entities.Dense() {
		fn(e)
	}
}
