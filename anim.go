package bento

// Animation is a effect rendered on a scene/sprite.
type Animation interface {
	Object

	Done() bool
}
