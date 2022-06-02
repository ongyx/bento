package bento

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// StageOptions defines various options for the stage.
// If Debug is not nil, a TPS counter is drawn on the top-left of the screen with the font.
// If Size is not nil, it will be used as the screen (window) size.
// If HiDPI is true, the screen size is scaled to the device scale factor.
type StageOptions struct {
	Debug *Font
	Size  *image.Point
	HiDPI bool
}

// Stage is a scene manager which implements the ebiten.Game interface.
// The current scene must never be nil.
type Stage struct {
	sc Scene
	op *StageOptions

	ts *Transition
	es []Entity
}

// NewStage creates a stage with an inital scene.
func NewStage(initial Scene, op *StageOptions) *Stage {
	if op == nil {
		op = &StageOptions{}
	}

	s := &Stage{op: op, ts: NewTransition()}
	s.Change(initial)

	return s
}

// Change changes the scene to render in the next frame.
func (s *Stage) Change(scene Scene) {
	log.Printf("(%p) changing scene to %p\n", s.sc, scene)

	if s.sc != nil {
		s.ts.Hide(s.sc.Exit())
	}

	s.sc = scene
	s.es = scene.Entities()

	go scene.Script(s)
}

// Update updates the current scene's state.
func (s *Stage) Update() error {
	Clock.increment()

	for _, e := range s.es {
		e.Update()
	}

	s.ts.Update()

	// Render the scene's enter transition if it hasn't entered yet
	// (or the previous scene exited).
	if s.ts.State() == Hidden {
		s.ts.Show(s.sc.Enter())
	}

	return nil
}

// Draw renders the current scene to the screen.
func (s *Stage) Draw(screen *ebiten.Image) {
	for _, e := range s.es {
		e.Draw(screen)
	}

	s.ts.Draw(screen)

	if s.op.Debug != nil {
		// draw tps/fps at the top left of the screen
		s.op.Debug.Write(
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

	if s.op.Size != nil {
		w = s.op.Size.X
		h = s.op.Size.Y
	} else {
		w = outsideWidth
		h = outsideHeight
	}

	if s.op.HiDPI {
		w = int(DPIScale(w))
		h = int(DPIScale(h))
	}

	return w, h
}
