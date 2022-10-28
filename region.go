package bento

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Region is a wrapper around a rect that represents an area in an image.
//
// Region is considered a value type:
// all methods that mutate the region return a new one
// (to avoid accidentally modifying the original).
type Region struct {
	image.Rectangle
}

// NewRegion creates a new region with an origin point and size.
// The origin is considered the top-left of the rectangle.
func NewRegion(origin, size image.Point) Region {
	return Region{
		image.Rectangle{
			Min: origin,
			Max: origin.Add(size),
		},
	}
}

// Expand increases the region by size on all sides.
func (r Region) Expand(size image.Point) Region {
	r.Min = r.Min.Sub(size)
	r.Max = r.Max.Add(size)
	return r
}

// Shrink decreases the region by size on all sides.
func (r Region) Shrink(size image.Point) Region {
	r.Min = r.Min.Add(size)
	r.Max = r.Max.Sub(size)
	return r
}

// SubImage returns the subimage bounded by the region from an image.
func (r Region) SubImage(i *ebiten.Image) *ebiten.Image {
	return i.SubImage(r.Rectangle).(*ebiten.Image)
}
