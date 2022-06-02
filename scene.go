package bento

// Scene is a collection of entities scripted with some actions.
type Scene interface {
	// Entities returns the entities in the scene.
	Entities() []Entity

	// Script asynchronously sends events to entities,
	// allowing entity state to be safely mutated without data races.
	// Entities typically expose a write-only channel to send events on,
	// then receive and handle events once per game tick in their Update method.
	Script(s *Stage)

	// Enter returns the enter animation of the scene.
	// If it is nil, the scene is immediately visible.
	Enter() Animation

	// Exit returns the exit animation of the scene.
	// If it is nil, the scene is immediately hidden (when changing scenes).
	Exit() Animation
}
