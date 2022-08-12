package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ongyx/bento/ecs"
)

const (
	spriteNum  = 256
	spriteSize = 4

	maxDistance = 4

	radian = math.Pi / 180

	screenSize = 256

	// frame interval between logging tps and fps.
	logInterval = (5 * 60)

	profile = true
)

func init() {
	rand.Seed(time.Now().Unix())
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

type Position image.Point

type Velocity struct {
	dx, dy int
}

type Transform struct {
	rotation int
}

type Sprite *ebiten.Image

type MovementSystem struct {
	view *ecs.View
}

func (m *MovementSystem) Init(w *ecs.World) {
	m.view = ecs.NewView(w, ecs.Query[Position](w), ecs.Query[Velocity](w), ecs.Query[Transform](w))
}

func (m *MovementSystem) Update(w *ecs.World) error {
	pos := ecs.Query[Position](w)
	vlc := ecs.Query[Velocity](w)
	tf := ecs.Query[Transform](w)

	m.view.Each(func(e ecs.Entity) {
		p := pos.Get(e)
		v := vlc.Get(e)
		t := tf.Get(e)

		p.X += v.dx
		p.Y += v.dy

		if p.X >= screenSize {
			p.X %= screenSize
		}

		if p.Y >= screenSize {
			p.Y %= screenSize
		}

		t.rotation++
	})

	return nil
}

type RenderSystem struct {
	view  *ecs.View
	frame int
}

func (r *RenderSystem) Init(w *ecs.World) {
	r.view = ecs.NewView(w, ecs.Query[Position](w), ecs.Query[Transform](w), ecs.Query[Sprite](w))
}

func (r *RenderSystem) Update(w *ecs.World) error {
	r.frame++
	if (r.frame % logInterval) == 0 {
		fmt.Printf("tps: %f, fps: %f\n", ebiten.CurrentTPS(), ebiten.CurrentFPS())
	}

	return nil
}

func (r *RenderSystem) Render(w *ecs.World, img *ebiten.Image) {
	pos := ecs.Query[Position](w)
	tf := ecs.Query[Transform](w)
	sprite := ecs.Query[Sprite](w)

	img.Fill(color.Black)

	r.view.Each(func(e ecs.Entity) {
		p := pos.Get(e)
		t := tf.Get(e)
		s := sprite.Get(e)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Rotate(float64(t.rotation) * radian)
		op.GeoM.Translate(float64(p.X), float64(p.Y))

		img.DrawImage(*s, op)
	})
}

func RenderScene() *ecs.World {
	w := ecs.NewWorld(spriteNum)

	pos := ecs.Register[Position](w, spriteNum)
	vlc := ecs.Register[Velocity](w, spriteNum)
	tf := ecs.Register[Transform](w, spriteNum)
	sprite := ecs.Register[Sprite](w, spriteNum)

	img := ebiten.NewImage(spriteSize, spriteSize)
	img.Fill(color.White)

	for i := 0; i < spriteNum; i++ {
		e := w.Spawn()

		pos.Set(e, Position{rand.Intn(screenSize), rand.Intn(screenSize)})

		vlc.Set(e, Velocity{rand.Intn(maxDistance) + 1, rand.Intn(maxDistance) + 1})

		tf.Set(e, Transform{rand.Intn(360)})

		sprite.Set(e, img)
	}

	w.Register(&MovementSystem{}, &RenderSystem{})

	return w
}
