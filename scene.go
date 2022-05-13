package bento

// Scene is a collection of compositors to draw on the screen.
type Scene interface {
	// Compositors returns the compositors in the scene.
	Compositors() []Compositor

	Script(s *Stage)

	Enter() Animation
	Exit() Animation
}
