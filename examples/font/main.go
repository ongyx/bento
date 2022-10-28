package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ongyx/bento"
	"github.com/ongyx/bento/examples/font/assets"
)

const (
	message = "Hello World!"
)

var (
	screen = image.Point{256, 256}
)

type Game struct {
	font *bento.Font
	box  *ebiten.Image

	init bool
}

func NewGame(font *bento.Font) *Game {
	return &Game{font: font}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(img *ebiten.Image) {
	msgsize := g.font.Size(message)

	if !g.init {
		g.box = ebiten.NewImage(msgsize.X, msgsize.Y)
		g.box.Fill(color.Gray{127})
	}

	m := &bento.Matrix{}
	c := bento.Center.Point(img)

	p := bento.Center.Align(c, image.Pt(g.box.Size()))
	m.TranslateP(p)

	img.DrawImage(g.box, m.Op())

	m.Reset()

	p = bento.Center.Align(c, msgsize)
	p.Y += msgsize.Y
	m.TranslateP(p)

	g.font.Draw(message, img, m.Op())
}

func (g *Game) Layout(w, h int) (lw, lh int) {
	size := bento.ScaleDPI(screen)
	return size.X, size.Y
}

func main() {
	if err := ebiten.RunGame(NewGame(assets.Inter)); err != nil {
		panic(err)
	}
}
