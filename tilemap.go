package bento

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Tilemap is a map of tiles that can be used to construct an image from a tileset.
type Tilemap struct {
	tiles   [][]int
	tileset *Tileset

	render *ebiten.Image
}

// NewTilemap creates a new tilemap from a slice of tiles and a tileset.
func NewTilemap(tiles [][]int, tileset *Tileset) *Tilemap {
	return &Tilemap{tiles: tiles, tileset: tileset}
}

func (m *Tilemap) Update() error {
	return nil
}

// Render renders the tilemap to an image.
func (m *Tilemap) Render(entity *Entity, size image.Point) *ebiten.Image {
	if m.render == nil {
		ts := m.tileset.tsize

		w := len(m.tiles[0]) * ts
		h := len(m.tiles) * ts
		m.render = ebiten.NewImage(w, h)

		for y, row := range m.tiles {
			for x, index := range row {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*ts), float64(y*ts))

				m.render.DrawImage(m.tileset.Tile(index), op)
			}
		}
	}

	return m.render
}
