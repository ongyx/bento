package bento

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

// Scroll allows several pieces of text to be scrolled across an image.
// point is the bottom-left point of the scroll.
type Scroll struct {
	Font  *Font
	Point image.Point
	Color color.Color

	text  string
	tpos  int
	tend  int
	tsize image.Point

	timer *Timer
}

// NewScroll creates a new scroll with the initial text.
func NewScroll(font *Font, tx string) *Scroll {
	s := &Scroll{Font: font, Color: color.White}
	s.SetSpeed(0.03)
	s.SetText(tx)

	return s
}

// SetSpeed changes the speed of scrolling text,
// where n is the number of seconds to wait between scrolling each character.
func (s *Scroll) SetSpeed(n float64) {
	s.timer = NewTimer(n, false)
}

// Text returns the current text in the scroll.
func (s *Scroll) Text() string {
	return s.text
}

// SetText changes the text currently scrolling.
func (s *Scroll) SetText(tx string) {
	size := text.BoundString(s.Font.Face, tx).Size()

	s.text = tx
	s.tpos = 0
	s.tend = len(tx)
	s.tsize = size
}

// Skip causes the next render to render the whole text instead of waiting for scrolling.
func (s *Scroll) Skip() {
	s.tpos = s.tend
}

// Done checks if the scrolling has finished.
func (s *Scroll) Done() bool {
	return s.tpos == s.tend
}

// Size returns the total size of the scroll.
func (s *Scroll) Size() image.Point {
	return s.tsize
}

// Update updates the state of the scroll.
func (s *Scroll) Update() error {
	if s.timer.Done() && s.tpos < s.tend {
		s.tpos++
	}

	return nil
}

// Draw renders the scroll on a new image.
func (s *Scroll) Draw(img *ebiten.Image) {
	t := s.text

	if s.tpos < s.tend {
		t = t[:s.tpos]
	}

	s.Font.Draw(t, s.Color, img, s.Point)
}
