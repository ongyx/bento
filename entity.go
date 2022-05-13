package bento

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Entity is a sprite with rendering state.
// Op is the options used to draw the entity to the screen.
type Entity struct {
	Sprite
	*Transition
	*ebiten.DrawImageOptions
}

// NewEntity constructs an entity from a sprite.
func NewEntity(s Sprite) *Entity {
	return &Entity{Sprite: s, Transition: NewTransition()}
}

// Update updates the entity's state.
func (e *Entity) Update() error {
	if e.Transition.RenderState() != Hidden {
		if err := e.Sprite.Update(); err != nil {
			return err
		}
	}

	if err := e.Transition.Update(); err != nil {
		return err
	}

	return nil
}

// Draw draws the entity to an image.
func (e *Entity) Draw(img *ebiten.Image) {
	if e.Transition.RenderState() != Hidden {
		r := e.Sprite.Render()
		e.Transition.Draw(r)

		img.DrawImage(r, e.DrawImageOptions)
	}
}
