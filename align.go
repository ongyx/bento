package bento

import (
	"image"
)

const (
	// Right moves an image to the right of a point.
	Right Align = 1 << iota
	// HCenter moves an image to the horizontal center of a point.
	HCenter
	// Left moves an image to the left of a point.
	Left
	// Bottom moves an image below a point.
	Bottom
	// VCenter moves an image to the vertical center of a point.
	VCenter
	// Top moves an image above a point.
	Top

	// Default is the default alignment of an image, which is below-right of a point.
	Default = Bottom | Right

	// Center is a shorthand for both horizontal and vertical center alignment.
	Center = VCenter | HCenter
)

// Align specifies the alignment of an image relative to an origin point.
//
//	         Top
//
//	     #---------#
//	     | VCenter |
//	Left |    /    | Right
//	     | HCenter |
//	     #---------#
//
//	       Bottom
//
// Align must have at most one horizontal and/or vertical flag.
type Align int

// Align returns the top-left point of the image,
// given the origin and the image size.
//
// The aligned point can then be directly used to draw the image (i.e with GeoM.Translate).
func (a Align) Align(origin, size image.Point) image.Point {
	if a != Default {
		w := size.X
		h := size.Y

		// horizontal
		if a.has(HCenter) {
			origin.X -= w / 2
		} else if a.has(Left) {
			origin.X -= w
		}

		// vertical
		// NOTE: The top left of the screen is (0, 0)!
		if a.has(VCenter) {
			origin.Y -= h / 2
		} else if a.has(Top) {
			origin.Y -= h
		}
	}

	return origin
}

// Point calculates a point in an image, using the alignment as a location.
func (a Align) Point(img image.Image) image.Point {
	b := img.Bounds()

	// top-left point
	p := b.Min

	w := b.Dx()
	h := b.Dy()

	if a.has(Right) {
		p.X += w
	} else if a.has(HCenter) {
		p.X += w / 2
	}

	if a.has(VCenter) {
		p.Y += h / 2
	} else if a.has(Bottom) {
		p.Y += h
	}

	return p
}

func (a Align) has(flag Align) bool {
	return (a & flag) != 0
}
