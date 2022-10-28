package ecs

import "github.com/hajimehoshi/ebiten/v2"

// System is an opaque state that updates entities.
type System interface {
	Init(world *World)
	Update(world *World) error
}

// Drawer is an opaque state that draws entities to the screen.
type Drawer interface {
	Draw(world *World, img *ebiten.Image)
}
