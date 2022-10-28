package bento

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Scroll allows several pieces of text to be scrolled across an image.
// point is the bottom-left point of the scroll.
type Scroll struct {
	Font        *Font
	FontOptions *FontOptions

	text  string
	tpos  int
	tend  int
	tsize image.Point

	timer *Timer
}

// NewScroll creates a new scroll with the initial text.
func NewScroll(font *Font, txt string) *Scroll {
	s := &Scroll{Font: font}
	s.SetSpeed(0.03)
	s.SetText(txt)

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
func (s *Scroll) SetText(txt string) {
	size := s.Font.Size(txt)

	s.text = txt
	s.tpos = 0
	s.tend = len(txt)
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
func (s *Scroll) Update() {
	s.timer.Tick()
	if s.timer.Done() && s.tpos < s.tend {
		s.tpos++
	}
}

// Draw renders the scroll on a new image.
func (s *Scroll) Draw(img *ebiten.Image, op *ebiten.DrawImageOptions) {
	s.Font.Draw(s.text[:s.tpos], img, op)
}
