package bento

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	// Hidden indicates the transition has exited or not yet entered.
	Hidden TransitionState = iota
	// Visible indicates the transition has entered.
	Visible
	// Entering indicates the transition is entering.
	Entering
	// Exiting indicates the transition is exiting.
	Exiting
)

// TransitionState represents the rendering state of a transition.
type TransitionState int

// Transition is an animation drawn over an image.
// The default render state is always Visible.
type Transition struct {
	state TransitionState
	anim  Animation
}

// NewTransition creates a new transition.
func NewTransition() *Transition {
	return &Transition{}
}

// Show sets the render state to Visible after rendering an enter animation.
// If enter is nil, the state is immediately changed.
func (t *Transition) Show(enter Animation) {
	log.Printf("(%p) entering (anim: %p)\n", t.anim, enter)

	if enter != nil {
		t.state = Entering
		t.anim = enter
	} else {
		// NOTE: this clobbers the existing render state!
		t.state = Visible
		t.anim = nil
	}
}

// Hide sets the render state to Hidden after rendering an exit transition.
// If exit is nil, the state is immediately changed.
func (t *Transition) Hide(exit Animation) {
	log.Printf("(%p) exiting (anim: %p)\n", t.anim, exit)

	if exit != nil {
		t.state = Exiting
		t.anim = exit
	} else {
		t.state = Hidden
		t.anim = nil
	}
}

// State returns the transition's current state.
func (t *Transition) State() TransitionState {
	return t.state
}

// Update updates the transition's logical state.
func (t *Transition) Update() {
	if a := t.transition(); a != nil {
		a.Update()

		if a.Done() {
			// transition finished, change rendering state
			switch t.state {
			case Entering:
				log.Printf("(%p) entered\n", t.anim)
				t.state = Visible
			case Exiting:
				log.Printf("(%p) exited\n", t.anim)
				t.state = Hidden
			default:
				// this really shouldn't happen.
				panic("transition: inconsistent state")
			}

			t.anim = nil
		}
	}
}

// Draw draws the transition to the image.
func (t *Transition) Draw(img *ebiten.Image) {
	if a := t.transition(); a != nil {
		a.Draw(img)
	}
}

func (t *Transition) transition() Animation {
	switch t.state {
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
