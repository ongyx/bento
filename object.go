package bento

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Object is a component that draws directly to a screen.
type Object interface {
	Component

	// Draw directly draws images onto the given canvas.
	// Implementations must not retain img.
	Draw(img *ebiten.Image)
}
