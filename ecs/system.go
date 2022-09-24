package ecs

import "github.com/hajimehoshi/ebiten/v2"

// System is an opaque state that updates entities.
type System interface {
	Init(w *World)
	Update(w *World) error
}

// Drawer is an opaque state that draws entities to the screen.
type Drawer interface {
	Draw(w *World, i *ebiten.Image)
}
