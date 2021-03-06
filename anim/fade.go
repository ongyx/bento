package anim

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ongyx/bento"
)

const (
	alphaMax = 255
)

// Fade is a transition that fades from/into a color.
type Fade struct {
	in    bool
	color color.NRGBA

	delta   *bento.Delta
	overlay *ebiten.Image
}

// NewFade creates a new fade transition with a duration.
// If in is true, the transition fades from solid to transparent,
// otherwise the transition fades from transparent to solid.
func NewFade(in bool, clr color.Color, duration float64) *Fade {
	d := bento.NewDelta(bento.Linear, image.Pt(alphaMax, 0), duration)

	return &Fade{
		in:    in,
		color: color.NRGBAModel.Convert(clr).(color.NRGBA),
		delta: d,
	}
}

func (f *Fade) Update() {
	f.delta.Update()

	a, _ := f.delta.Delta()

	if f.in {
		a = alphaMax - a
	}

	f.color.A = uint8(a)
}

func (f *Fade) Draw(img *ebiten.Image) {
	if f.overlay == nil {
		f.overlay = ebiten.NewImage(img.Size())
	}

	f.overlay.Fill(f.color)

	img.DrawImage(f.overlay, nil)
}

func (f *Fade) Done() bool {
	return f.delta.Done()
}
