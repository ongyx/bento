package bento

// Component is a state that updates every tick.
//
// If a component uses another component, it should be updated before using its state:
//
//	type myComponent struct {
// 		child Component
// 	}
//
// 	func (m *myComponent) Update() error {
//		if err := m.child.Update(); err != nil {
//			return err
//		}
//
// 		// use child component's state...
//
//		return nil
// 	}
//
type Component interface {
	// Update updates the component's state.
	Update() error
}
