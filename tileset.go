package bento

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type tileCache map[int]*ebiten.Image

// Tileset is a image with tiles that can be tiled into a single sprite.
type Tileset struct {
	image *ebiten.Image
	tsize int

	cache tileCache
	size  image.Point
}

// NewTileset creates a new tileset with a image and tile size.
// If cache is true, calls to the Tile method will lazily cache tile images for faster loading.
func NewTileset(img *ebiten.Image, tile int, cache bool) *Tileset {
	w, h := img.Size()

	if tile <= 0 || tile > w || tile > h {
		panic("tileset: tilesize must be within bounds of tileset")
	}

	var tc tileCache
	if cache {
		tc = make(tileCache)
	}

	return &Tileset{img, tile, tc, image.Pt(w/tile, h/tile)}
}

// Size returns the tileset size as (columns, rows).
func (t *Tileset) Size() image.Point {
	return t.size
}

// Tilesize returns the size of a tile in the tileset.
func (t *Tileset) Tilesize() int {
	return t.tsize
}

// Tile returns the tile at index as a subimage of the tileset,
// where index is the tile column multiplied by the tile row.
// This panics if the index is out of bounds.
func (t *Tileset) Tile(index int) *ebiten.Image {
	if index >= (t.size.X * t.size.Y) {
		panic(fmt.Sprintf("tileset: index %d out of bounds", index))
	}

	var tile *ebiten.Image

	if t.cache != nil {
		if ct, ok := t.cache[index]; ok {
			tile = ct
		} else {
			tile = t.tile(index)
			t.cache[index] = tile
		}
	} else {
		tile = t.tile(index)
	}

	return tile
}

// Render renders a tilemap using a tileset to an image.
// The tileset should be well-formed: all rows must be the same length and it should have at least one row.
func (t *Tileset) Render(tileset [][]int) *ebiten.Image {
	w := len(tileset[0]) * t.tsize
	h := len(tileset) * t.tsize
	img := ebiten.NewImage(w, h)

	for y, row := range tileset {
		for x, tile := range row {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*t.tsize), float64(y*t.tsize))

			img.DrawImage(t.Tile(tile), op)
		}
	}

	return img
}

func (t *Tileset) tile(index int) *ebiten.Image {
	x := (index % t.size.X) * t.tsize
	y := (index / t.size.X) * t.tsize

	return t.image.SubImage(image.Rect(x, y, x+t.tsize, y+t.tsize)).(*ebiten.Image)
}
