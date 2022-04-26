package bento

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Character is a character sprite with one or more frames.
type Character struct {
	index, limit int
	frames       []*ebiten.Image
}

// NewCharacter creates a new character from a slice of frames.
func NewCharacter(frames []*ebiten.Image) *Character {
	c := &Character{}
	c.SetFrames(frames)

	return c
}

// SetFrames sets the character's frames.
func (c *Character) SetFrames(frames []*ebiten.Image) {
	c.index = 0
	c.limit = len(frames)
	c.frames = frames
}

// Update updates the character's logical state.
func (c *Character) Update() error {
	c.index++
	if c.index >= c.limit {
		c.index = 0
	}

	return nil
}

// Render renders the character's current frame to an image.
func (c *Character) Render(entity *Entity, size image.Point) *ebiten.Image {
	entity.Show(nil)

	return c.frames[c.index]
}
