package bento

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Atlas is a tile-based texture atlas for sprites.
type Atlas struct {
	img, empty *ebiten.Image

	tile, size image.Point
}

// NewAtlas creates a new atlas with an image and tile size.
// The tile size for each axis must satisfy (1 <= size <= image_size).
func NewAtlas(img *ebiten.Image, tile image.Point) *Atlas {
	w, h := img.Size()
	b := image.Rect(1, 1, w, h)

	if !tile.In(b) {
		panic("atlas: tile size out of bounds")
	}

	return &Atlas{
		img:   img,
		empty: ebiten.NewImage(tile.X, tile.Y),
		tile:  tile,
		size:  image.Pt(w/tile.X, w/tile.Y),
	}
}

// Size returns the size of the atlas in tiles.
// The size should be consider a closed interval, i.e (size.X * size.Y) is the maximum texture index.
func (a *Atlas) Size() image.Point {
	return a.size
}

// Tile returns the size of a tile.
func (a *Atlas) Tile() image.Point {
	return a.tile
}

// Texture returns the tile at index as a subimage of the atlas.
// Index is the tile column multiplied by the tile row.
// If index is negative, an empty image the same size as a tile is returned instead.
func (a *Atlas) Texture(index int) *ebiten.Image {
	if index >= (a.size.X * a.size.Y) {
		panic(fmt.Sprintf("atlas: index %d out of bounds", index))
	}

	if index < 0 {
		return a.empty
	}

	x := (index % a.size.X) * a.tile.X
	y := (index / a.size.Y) * a.tile.Y

	b := image.Rect(x, y, x+a.tile.X, y+a.tile.Y)

	return a.img.SubImage(b).(*ebiten.Image)
}

// Draw draws a layer to an image.
// If img is nil, a new image large enough to display all the tiles is created.
func (a *Atlas) Draw(img *ebiten.Image, layer Layer) *ebiten.Image {
	if img == nil {
		w := len(layer[0]) * a.tile.X
		h := len(layer) * a.tile.Y

		img = ebiten.NewImage(w, h)
	}

	for y, row := range layer {
		for x, tile := range row {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*a.tile.X), float64(y*a.tile.Y))

			img.DrawImage(a.Texture(tile), op)
		}
	}

	return img
}

// Composite draws several layers to an image in order.
func (a *Atlas) Composite(img *ebiten.Image, layers ...Layer) *ebiten.Image {
	for _, layer := range layers {
		img = a.Draw(img, layer)
	}

	return img
}

// Layer is a 2D slice of tile indices.
type Layer [][]int
