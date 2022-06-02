package bento

// Component encapsulates some opaque state that updates every tick.
// Implementations should call the Update method of sub-components, if any:
//
//	type myComponent struct {
// 		child Component
// 	}
//
// 	func (m *myComponent) Update() error {
//		m.child.Update()
//
// 		// use child component's state...
//
//		return nil
// 	}
//
// Trying to access sub-component state before the first call to Update is undefined behaviour.
type Component interface {
	Update()
}
