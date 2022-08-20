# Entity Component System

## Introduction

ECS, or Entity Component System, is a programming paradigm commonly found in gamedev which emphasizes separation of behaviour from data [[1]].

```go
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
}

// Render is called once per frame, approximately 60 frames per second.
// It is optional to implement this.
func (p *PrintSystem) Render(w *ecs.World, img *ebiten.Image) {}

// Scenes are plain functions that create a world and performs setup on it.
func scene() *ecs.World {
  // Create a world with capacity for 8 entities.
  // You can change the size (8) to preallocate beforehand, so adding entities is faster.
  w := NewWorld(8)

  // Register types, which returns a component.
  // Types must be registered to add them to the world.
  text := ecs.Register[Text](w, 8)

  // Spawn entities in the world and add components to them.
  e := w.Spawn()
  text.Set(e, "Hello World!")

  // Register systems in the world.
  w.Register(&PrintSystem{})

  return w
}

func main() {
  w := scene()
  // From here, you can wrap the world in your own ebiten.Game implementation.
  // This allows flexibility in using ECS within existing code.
}
```

## Technical Details

With the introduction of generics in Go 1.18, it is now possible to create a ECS implementation without having to use reflection for the most part.

The only reflection involved is `reflect.TypeOf` to generate a unique identifer for a generic component type:

```go
func typeof[T any]() reflect.Type {
  var zeroValue T
  return reflect.TypeOf(zeroValue)
}
```

A bitset is used with the unique identifer to efficiently query entities in the world with the correct components.
Afterwards, `Query[T]()` and `Register[T]()` uses regular type assertions to get a `Component[T]`, which is much cheaper than using plain reflection.

Part of this hybrid bitset-generics approach was inspired by this blog post [[2]], which was originally written in C++.

[1]: https://github.com/SanderMertens/ecs-faq
[2]: https://austinmorlan.com/posts/entity_component_system
