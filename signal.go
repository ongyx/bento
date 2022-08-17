package bento

// Signal is a broadcaster that notifies several slots (callbacks) when a value is emitted.
//
//	s := NewSignal[string](0)
//	s.Connect(func(s string) {
//		fmt.Println(s)
//	})
//
//	s.Emit("Hello World!")
//
// Emitting a signal is concurrent-safe.
type Signal[T any] struct {
	event *Event[T]
	slots []func(T)
}

func NewSignal[T any](size int) *Signal[T] {
	s := &Signal[T]{event: NewEvent[T](size)}
	s.event.Notify(s.onNotify)

	return s
}

// Connect connects the slot to the signal.
func (s *Signal[T]) Connect(slot func(T)) {
	s.slots = append(s.slots, slot)
}

// Emit notifies all connected slots.
func (s *Signal[T]) Emit(value T) {
	s.event.Emit(value)
}

func (s *Signal[T]) onNotify(value T) {
	for _, slot := range s.slots {
		slot(value)
	}
}
