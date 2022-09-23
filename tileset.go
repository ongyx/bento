package bento

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Layer is a 2D slice of tile indices.
type Layer [][]int

// Tileset is a image with tiles that can be tiled into a single sprite.
type Tileset struct {
	atlas, empty *ebiten.Image

	size  int
	bound image.Point
}

// NewTileset creates a new tileset with a image and tile size.
func NewTileset(img *ebiten.Image, size int) *Tileset {
	w, h := img.Size()

	if size <= 0 || size > w || size > h {
		panic("tileset: tile size must be within bounds of tileset")
	}

	return &Tileset{
		atlas: img,
		// created here so rendering an empty tile in a layer is efficient
		empty: ebiten.NewImage(size, size),
		size:  size,
		bound: image.Pt(w/size, h/size),
	}
}

// Bound returns the index bounds of the tileset as (columns, rows).
func (t *Tileset) Bound() image.Point {
	return t.bound
}

// Size returns the length of a tile in pixels.
func (t *Tileset) Size() int {
	return t.size
}

// Tile returns the tile at index as a subimage of the tileset, where index is the tile column multiplied by the tile row.
// If index is negative, an empty image the size of a tile is returned instead.
func (t *Tileset) Tile(index int) *ebiten.Image {
	if index >= (t.bound.X * t.bound.Y) {
		panic(fmt.Sprintf("tileset: index %d out of bounds", index))
	}

	if index < 0 {
		return t.empty
	}

	x := (index % t.bound.X) * t.size
	y := (index / t.bound.X) * t.size

	b := image.Rect(x, y, x+t.size, y+t.size)

	return t.atlas.SubImage(b).(*ebiten.Image)
}

// Render renders a layer to an image.
// The layer must have rows of the same length.
func (t *Tileset) Render(layer Layer) *ebiten.Image {
	w := len(layer[0]) * t.size
	h := len(layer) * t.size
	img := ebiten.NewImage(w, h)

	for y, row := range layer {
		for x, tile := range row {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*t.size), float64(y*t.size))

			img.DrawImage(t.Tile(tile), op)
		}
	}

	return img
}

// Composite renders several layers into one image in order.
func (t *Tileset) Composite(layers []Layer) *ebiten.Image {
	var img *ebiten.Image

	for _, l := range layers {
		i := t.Render(l)
		if img != nil {
			img.DrawImage(i, nil)
		} else {
			img = i
		}
	}

	return img
}
