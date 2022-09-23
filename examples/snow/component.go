package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Position image.Point

type Velocity struct {
	dx, dy int
}

type Transform struct {
	rotation int
}

type Sprite *ebiten.Image
