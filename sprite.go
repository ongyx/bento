package bento

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Sprite is a component that renders to an image.
type Sprite interface {
	Component

	// Render renders and returns an image, given the size of the canvas the image will be drawn to.
	Render() *ebiten.Image
}
