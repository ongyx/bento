package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ongyx/bento/ecs"
)

// Components are plain data which systems act on.
type Text string

// Systems iterate through components to use their data.
type PrintSystem struct {
	// Views allow efficient traversal of entities with specific components.
	view *ecs.View
}

// Init is called only once, before the first call to the Update method.
// Initialization of views and any other state is done here.
func (p *PrintSystem) Init(w *ecs.World) {
	// We create a view that searches for entities with the Text component.
	p.view = ecs.NewView(w, ecs.Type[Text]())
}

// Update is called once per logical tick, approximately 60 times per second.
func (p *PrintSystem) Update(w *ecs.World) error {
	// Query returns a component from a world to get data from.
	text := ecs.Query[Text](w)

	// Iterate through all entities matched against by the view.
	p.view.Each(func(e ecs.Entity) {
		// Get the text data and print it.
		t := text.Get(e)
		fmt.Println(*t)
	})

	return nil
}

// Render is called once per frame, approximately 60 frames per second.
// It is optional to implement this.
func (p *PrintSystem) Render(w *ecs.World, img *ebiten.Image) {}

// Scenes are plain functions that create a world and performs setup on it.
func scene() *ecs.World {
	// Create a world with capacity for 8 entities.
	// You can change the size (8) to preallocate beforehand, so adding entities is faster.
	w := ecs.NewWorld(8)

	// Register types, which returns a component.
	// Types must be registered to add them to the world.
	text := ecs.Register[Text](w, 8)

	// Spawn entities in the world and add components to them.
	e := w.Spawn()
	text.Insert(e, "Hello World!")

	// Register systems in the world.
	w.Register(&PrintSystem{})

	return w
}

func main() {
	// Worlds can be created in isolation (i.e no global variables).
	// It also implements ebiten.Game, so it can be used as-is.
	w := scene()
	for i := 0; i < 5; i++ {
		w.Update()
	}
}
