package bento

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Geometry provides convenience for operations on a matrix.
type Geometry struct {
	M ebiten.GeoM
}

// Align translates the matrix by alignment with size.
func (g *Geometry) Align(a Align, size image.Point) {
	p := a.Align(image.Point{}, size)
	g.M.Translate(float64(p.X), float64(p.Y))
}

// Scale scales the matrix by a scalar factor.
func (g *Geometry) Scale(f float64) {
	g.M.Scale(f, f)
}

// Translate translates the matrix by a point.
func (g *Geometry) Translate(p image.Point) {
	g.M.Translate(float64(p.X), float64(p.Y))
}
