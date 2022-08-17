package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ongyx/bento/ecs"
)

const (
	screenSize = 256

	// frame interval between logging tps and fps.
	logInterval = (5 * 60)

	profile = true
)

func init() {
	if profile {
		go func() {
			fmt.Println(http.ListenAndServe("localhost:6060", nil))
		}()
	}
}

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
