package bento

// RenderState represents the rendering state/lifecycle of a sprite wrapped in a transition.
//go:generate stringer -type=RenderState
type RenderState int

const (
	// Hidden indicates the sprite is not drawing to the screen.
	Hidden RenderState = iota
	// Entering indicates a enter transition is rendering on the sprite.
	Entering
	// Visible indicates the sprite is drawing normally to the screen.
	Visible
	// Exiting indicates a exit transition is rendering on the sprite.
	Exiting
)
