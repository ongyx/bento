package bento

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// StageOptions defines various options for the stage.
// If Font is not nil, a TPS counter is drawn on the top-left of the screen.
// If Size is not nil, it will be used as the screen (window) size.
// If HiDPI is true, the screen size is scaled to the device scale factor.
type StageOptions struct {
	Font  *Font
	Size  *image.Point
	HiDPI bool
}

// Stage is a scene manager which implements the ebiten.Game interface.
// The current scene must never be nil.
type Stage struct {
	Op StageOptions

	ts   *Transition
	sc   Scene
	objs []Object

	fn Stream[func()]
}

// NewStage creates a stage with an inital scene.
func NewStage(inital Scene) *Stage {
	s := &Stage{ts: NewTransition(), fn: NewStream[func()](0)}
	s.Change(inital)

	return s
}

func (s *Stage) OnNextFrame(fn func()) {
	s.fn.Write(fn)
}

// Change changes the scene to render in the next frame.
func (s *Stage) Change(scene Scene) {
	log.Printf("(%p) changing scene to %p\n", s.sc, scene)

	if s.sc != nil {
		s.ts.Hide(s.sc.Exit())
	}

	s.sc = scene
	s.objs = scene.Objects()

	go scene.Script(s)
}

// Update updates the current scene's state.
func (s *Stage) Update() error {
	Clock.increment()

	for _, o := range s.objs {
		if err := o.Update(); err != nil {
			return err
		}
	}

	if err := s.ts.Update(); err != nil {
		return err
	}

	// Render the scene's enter transition if it hasn't entered yet
	// (or the previous scene exited).
	if s.ts.RenderState() == Hidden {
		s.ts.Show(s.sc.Enter())
	}

	if fn := s.fn.Poll(); fn != nil {
		(*fn)()
	}

	return nil
}

// Draw renders the current scene to the screen.
func (s *Stage) Draw(screen *ebiten.Image) {
	for _, o := range s.objs {
		o.Draw(screen)
	}

	s.ts.Draw(screen)

	if s.Op.Font != nil {
		// draw tps/fps at the top left of the screen
		s.Op.Font.Write(
			fmt.Sprintf("tps: %0.2f", ebiten.CurrentTPS()),
			color.White,
			screen,
			image.Pt(0, 0),
			Default,
		)
	}
}

// Layout returns the screen's size.
func (s *Stage) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	var w, h int

	if s.Op.Size != nil {
		w = s.Op.Size.X
		h = s.Op.Size.Y
	} else {
		w = outsideWidth
		h = outsideHeight
	}

	if s.Op.HiDPI {
		w = int(DPIScale(w))
		h = int(DPIScale(h))
	}

	return w, h
}
