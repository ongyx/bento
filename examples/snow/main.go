package main

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ongyx/bento/ecs"
)

const (
	screenSize = 256

	// frame interval between logging tps and fps.
	logInterval = (5 * 60)
)

func main() {
	g := &game{RenderScene()}
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

type game struct {
	*ecs.World
}

func (g *game) Layout(w, h int) (sw, sh int) {
	return screenSize, screenSize
}
