package bento

import "github.com/hajimehoshi/ebiten/v2"

// Entity is a component that can draw to an image.
type Entity interface {
	Component

	Draw(img *ebiten.Image)
}
