package bento

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Character struct {
	Position image.Point
	Align    Align
	Scale    float64

	index  int
	limit  int
	frames []*ebiten.Image
}

func NewCharacter(frames []*ebiten.Image) *Character {
	c := &Character{Align: Default, Scale: 1}
	c.SetFrames(frames)

	return c
}

func (c *Character) Frame() *ebiten.Image {
	return c.frames[c.index]
}

func (c *Character) SetFrames(frames []*ebiten.Image) {
	c.index = 0
	c.limit = len(frames)
	c.frames = frames
}

func (c *Character) Update() error {
	c.index++

	if c.index >= c.limit {
		c.index = 0
	}

	return nil
}

func (c *Character) Render(entity *Entity, size image.Point) *ebiten.Image {
	f := c.Frame()

	var g Geometry
	g.Align(c.Align, f.Bounds().Size())
	g.Scale(c.Scale)
	g.Translate(c.Position)

	entity.Op.GeoM = g.M

	return f
}
