package ecs

import "github.com/hajimehoshi/ebiten/v2"

const maxPools = 64

// World contains all data related to entities, components and systemw.
type World struct {
	pools  []pool
	poolID map[TypeID]uint8

	entities []entity
	// deleted is the index of the most recently deleted entity.
	// if there are no deleted entities, it is set to invalid.
	deleted uint32

	systems   []System
	renderers []Renderer

	init bool
}

// NewWorld creates a new world with (size) capacity for entitiew.
func NewWorld(size int) *World {
	return &World{
		pools:  make([]pool, 0, maxPools),
		poolID: make(map[TypeID]uint8, maxPools),

		entities: make([]entity, 0, size),
		deleted:  invalid,
	}
}

// Spawn creates a new entity in the world.
func (w *World) Spawn() Entity {
	var e Entity

	if w.deleted != invalid {
		// if there is a recently deleted entity, recycle it
		e = w.respawn(w.deleted)
	} else {
		id := uint32(len(w.entities))
		if id == invalid {
			panic("world: too many entities")
		}

		e = Entity{id: id}

		// append new entity
		w.entities = append(w.entities, entity{entity: e})
	}

	return e
}

// Despawn removes the entity from the world.
func (w *World) Despawn(e Entity) {
	// bump the entity version
	entity := &w.entities[e.id].entity
	entity.version++

	id := entity.id

	// make the entity ID a pointer to the previous deleted entity, if any
	if w.deleted != invalid {
		entity.id = w.deleted
	} else {
		entity.id = invalid
	}

	// this is now the most recently deleted entity
	w.deleted = id

	// remove entity from all pools
	for _, p := range w.pools {
		p.Delete(e)
	}
}

// Register adds the systems to the world.
// A System may optionally implement Renderer, whose Render method is called when drawing to the screen.
func (w *World) Register(systems ...System) {
	for _, s := range systems {
		w.systems = append(w.systems, s)
		if r, ok := s.(Renderer); ok {
			w.renderers = append(w.renderers, r)
		}
	}
}

// Update updates the state of the world's systems in the order they were registered.
// This implements the ebiten.Game interface.
func (w *World) Update() error {
	for _, s := range w.systems {
		if !w.init {
			s.Init(w)
		}

		if err := s.Update(w); err != nil {
			return err
		}
	}

	if !w.init {
		// all systems initalised by this point
		w.init = true
	}

	return nil
}

// Draw renders all renderers to the screen in the order they were registered.
// This implements the ebiten.Game interface.
func (w *World) Draw(screen *ebiten.Image) {
	for _, r := range w.renderers {
		r.Render(w, screen)
	}
}

// respawn resurrects a previously deleted entity.
func (w *World) respawn(id uint32) Entity {
	e := &w.entities[id].entity

	// if entity ID is invalid, this is the end of the linked list
	// otherwise, destroyed is set to the next deleted entity.
	if e.id == invalid {
		w.deleted = invalid
	} else {
		w.deleted = e.id
	}

	// entity ID is restored so it can now be reused.
	e.id = id

	return *e
}

// entity is an entry for a entity in a world.
type entity struct {
	entity Entity
	sig    Signature
}
