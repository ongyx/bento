package bento

// Animation is a effect rendered on a scene/sprite.
type Animation interface {
	Compositor

	// Done checks if the animation has finished.
	Done() bool
}
