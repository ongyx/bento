package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ongyx/bento"
)

// In bento, games are represented as one or more Scenes, or a game 'level'.
// The demo struct implements the bento.Scene interface.
type demo struct{}

// Entities are how bento implements game objects.
// They update every game tick and draw to the screen.
// Entities can be composited to form more advanced entities.
func (d *demo) Entities() []bento.Entity {
	return nil
}

// Scripts can alter entities by sending events to them,
// and run in their own goroutine.
func (d *demo) Script(s *bento.Stage) {}

// Enter and exit transitions can be shown in-between changing scenes.
func (d *demo) Enter() bento.Animation { return nil }
func (d *demo) Exit() bento.Animation  { return nil }

func main() {
	stage := bento.NewStage(&demo{}, nil)
	ebiten.SetWindowTitle("Hello World!")

	// You can directly start the game by passing the stage to RunGame, as bento.Stage implements the ebiten.Game interface.
	if err := ebiten.RunGame(stage); err != nil {
		log.Fatal(err)
	}
}
