package bento

// Animation is a effect rendered on another entity.
type Animation interface {
	Entity

	Done() bool
}
