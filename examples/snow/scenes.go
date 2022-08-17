package main

import (
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ongyx/bento/ecs"
)

const (
	spriteNum  = 256
	spriteSize = 4

	maxDistance = 4

	radian = math.Pi / 180
)

func init() {
	rand.Seed(time.Now().Unix())
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
