// Package bento provides a higher-level API on top of ebiten to simplify 2D game development, akin to Unity.
//
// bento is meant to be used alongside ebiten; here is a simple example.
//
// 	// In bento, games are represented as one or more Scenes, or a game 'level'.
// 	// The demo struct implements the bento.Scene interface.
// 	type demo struct{}
//
// 	// bento follows ebiten's game loop: Update -> Draw -> repeat.
// 	// Here, you update logical state such as changing to another scene with stage.Change(newScene).
// 	func (d *demo) Update(stage *bento.Stage) error { return nil }
//
// 	// bento allows drawing directly to the screen too, but for more advanced mechanics/images consider using entities instead.
// 	func (d *demo) Draw(screen *ebiten.Image) {}
//
// 	// Enter and exit transitions can be shown in-between changing scenes.
// 	func (d *demo) Enter() bento.Animation { return nil }
// 	func (d *demo) Exit() bento.Animation { return nil }
//
// 	// Entities are the core of any 2D game made with bento; the scene is a special entity that renders other entities.
// 	func (d *demo) Entities() []*bento.Entity { return nil }
//
// 	func main() {
// 		stage := bento.NewStage(&demo{})
// 		ebiten.SetWindowTitle("Hello World!")
//
// 		// You can directly start the game by passing the stage to RunGame, as bento.Stage implements the ebiten.Game interface.
// 		if err := ebiten.RunGame(stage); err != nil {
// 			log.Fatal(err)
// 		}
// 	}
package bento
