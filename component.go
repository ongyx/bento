package bento

// Component is a state that updates every tick.
// Update should be called _after_ using a sub-component's state, i.e
//
//	type myComponent struct {
//		sub Component
//	}
//
//	func (mc *myComponent) Update() error {
//
//		// use mc.sub here
//
//		if err := mc.sub.Update(); err != nil {
//			return err
//		}
//
//	}
//
type Component interface {
	Update() error
}
