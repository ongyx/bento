package bento

// RenderState represents the rendering state/lifecycle of an entity on the stage.
//go:generate stringer -type=RenderState
type RenderState int

const (
	// Hidden indicates the entity is not drawing to the screen.
	Hidden RenderState = iota
	// Entering indicates a enter transition is rendering on the entity.
	Entering
	// Visible indicates the entity is drawing normally to the screen.
	Visible
	// Exiting indicates a exit transition is rendering on the entity.
	Exiting
)
