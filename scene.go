package bento

// Scene is a collection of objects scripted with some actions.
type Scene interface {
	// Objects returns the objects in the scene.
	Objects() []Object

	// Script issues commands to the scene's compositors.
	// For example, if you have a Scroll:
	//
	// 	func (m *myScene) Script(s *Stage) {
	// 		m.scroll.SetText("Hello World!")
	// 		// TODO: add channel to wait for events (i.e scroll finish?)
	// 		s.Change(anotherScene)
	// 	}
	//
	Script(s *Stage)

	// Enter returns the enter animation of the scene.
	// If it is nil, the scene is immediately visible.
	Enter() Animation

	// Exit returns the exit animation of the scene.
	// If it is nil, the scene is immediately hidden (when changing scenes).
	Exit() Animation
}
