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
// If debug is not nil, debug mode is enabled.
type Stage struct {
	Op StageOptions

	scene      Scene
	transition *Transition

	// snapshot holds the last rendered frame of the current scene.
	// This is used mainly for transitions.
	snapshot *ebiten.Image
}

// NewStage creates a stage with an inital scene.
// NOTE: The initial scene's enter animation is rendered!
func NewStage(initial Scene) *Stage {
	s := &Stage{transition: NewTransition()}
	s.Change(initial)

	return s
}

// Change changes the scene to render in the next frame.
func (s *Stage) Change(newScene Scene) {
	oldScene := s.scene

	log.Printf("(%p) changing scene to %p\n", oldScene, newScene)

	if oldScene != nil {
		s.transition.Hide(oldScene.Exit())
	}

	s.scene = newScene
}

// Update updates the current scene's state.
func (s *Stage) Update() error {
	if s.transition.RenderState() != Exiting {
		if err := s.scene.Update(s); err != nil {
			return err
		}
	}

	if s.transition.RenderState() == Hidden {
		// finished old scene's exit transition
		// render the enter transition of the new scene
		s.transition.Show(s.scene.Enter())
	}

	if err := s.transition.Update(); err != nil {
		return err
	}

	return nil
}

// Draw renders the current scene to the screen.
func (s *Stage) Draw(screen *ebiten.Image) {
	if s.snapshot == nil {
		s.snapshot = ebiten.NewImage(screen.Size())
	}

	// render the scene only if we aren't exiting
	if s.transition.RenderState() != Exiting {
		//log.Printf("(%p) drawing to snapshot with %v state\n", s.scene, s.transition.RenderState())
		s.snapshot.Clear()
		s.scene.Draw(s.snapshot)
	}

	screen.DrawImage(s.snapshot, nil)

	s.transition.Draw(screen)

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

	// update tick
	Clock.increment()
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
