package bento

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Transition is an animation drawn over an image.
// The default render state is always Visible.
type Transition struct {
	rs   RenderState
	anim Animation
}

// NewTransition creates a new transition that wraps an existing sprite.
func NewTransition() *Transition {
	return &Transition{rs: Hidden}
}

// Show sets the render state to Visible after rendering an enter animation.
// If enter is nil, the state is immediately changed.
func (t *Transition) Show(enter Animation) {
	log.Printf("(%p) entering (anim: %p)\n", t.anim, enter)

	if enter != nil {
		t.rs = Entering
		t.anim = enter
	} else {
		// NOTE: this clobbers the existing render state!
		t.rs = Visible
		t.anim = nil
	}
}

// Hide sets the render state to Hidden after rendering an exit transition.
// If exit is nil, the state is immediately changed.
func (t *Transition) Hide(exit Animation) {
	log.Printf("(%p) exiting (anim: %p)\n", t.anim, exit)

	if exit != nil {
		t.rs = Exiting
		t.anim = exit
	} else {
		t.rs = Hidden
		t.anim = nil
	}
}

// RenderState returns the rendering state of the transition.
func (t *Transition) RenderState() RenderState {
	return t.rs
}

// Update updates the transition's state.
func (t *Transition) Update() error {
	if a := t.transition(); a != nil {
		if a.Done() {
			// transition finished, change rendering state
			switch t.rs {
			case Entering:
				log.Printf("(%p) entered\n", t.anim)
				t.rs = Visible
			case Exiting:
				log.Printf("(%p) exited\n", t.anim)
				t.rs = Hidden
			default:
				// this really shouldn't happen.
				panic("transition: inconsistent state")
			}

			t.anim = nil
		}

		if err := a.Update(); err != nil {
			return err
		}
	}

	return nil
}

// Draw draws the transition over an image.
func (t *Transition) Draw(img *ebiten.Image) {
	if t.rs != Hidden {
		if a := t.transition(); a != nil {
			a.Draw(img)
		}
	}
}

func (t *Transition) transition() Animation {
	switch t.rs {
	case Entering, Exiting:
		// sanity check
		if t.anim == nil {
			panic("transition: anim is nil")
		}

		return t.anim

	default:
		return nil
	}
}
