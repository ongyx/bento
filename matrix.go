package bento

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Matrix is a wrapper around a geometry matrix that supports point-based operations.
type Matrix struct {
	ebiten.GeoM
}

// TranslateP translates the matrix by a point.
func (m *Matrix) TranslateP(p image.Point) {
	m.Translate(float64(p.X), float64(p.Y))
}

// ScaleP scales the matrix by a point.
func (m *Matrix) ScaleP(p image.Point) {
	m.Scale(float64(p.X), float64(p.Y))
}

// Op returns DrawImageOptions with the matrix for convenience.
func (m *Matrix) Op() *ebiten.DrawImageOptions {
	return &ebiten.DrawImageOptions{GeoM: m.GeoM}
}
