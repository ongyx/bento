package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ongyx/bento/ecs"
)

type MovementSystem struct {
	view *ecs.View
}

func (m *MovementSystem) Init(w *ecs.World) {
	m.view = ecs.NewView(w, ecs.Type[Position](), ecs.Type[Velocity](), ecs.Type[Transform]())
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
	r.view = ecs.NewView(w, ecs.Type[Position](), ecs.Type[Transform](), ecs.Type[Sprite]())
}

func (r *RenderSystem) Update(w *ecs.World) error {
	r.frame++
	if (r.frame % logInterval) == 0 {
		fmt.Printf("tps: %f, fps: %f\n", ebiten.ActualTPS(), ebiten.ActualFPS())
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
