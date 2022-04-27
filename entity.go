package bento

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Closure is a function that is called after a sprite renders but before any transition draws on it.
type Closure func(img *ebiten.Image, entity *Entity, size image.Point)

// Entity is a sprite with rendering state.
// While a sprite renders to an image, an entity handles drawing the rendered image to the screen.
// Op is the options used to draw the entity to the screen.
type Entity struct {
	Sprite
	*Transition

	Op *ebiten.DrawImageOptions
}

// NewEntity constructs an entity from a sprite.
// Entities are hidden by default.
func NewEntity(sprite Sprite) *Entity {
	return &Entity{Sprite: sprite, Transition: NewTransition(), Op: &ebiten.DrawImageOptions{}}
}

// Update updates the sprite's state.
func (e *Entity) Update() error {
	if err := e.Sprite.Update(); err != nil {
		return err
	}

	if err := e.Transition.Update(); err != nil {
		return err
	}

	return nil
}

// Draw draws the sprite's render onto the screen.
func (e *Entity) Draw(screen *ebiten.Image) {
	size := screen.Bounds().Size()

	render := e.Sprite.Render(e, size)

	e.Transition.Draw(render)

	if e.Transition.RenderState() != Hidden {
		screen.DrawImage(render, e.Op)
	}
}
