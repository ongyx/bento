package bento

// Signal is a event broadcaster that notifies several slots (callbacks) when a value is emitted.
//
//	s := NewSignal(0)
//	s.Connect(func(s string) {
//		fmt.Println(s)
//	})
//
//	e.Emit("Hello World!")
//
// Emitting a signal is concurrent-safe.
type Signal[T any] struct {
	event *Event[T]
	slots []func(T)
}

// NewSignal creates a new signal with a buffer of (size) values.
func NewSignal[T any](size int) *Signal[T] {
	s := &Signal[T]{event: NewEvent[T](size)}
	s.event.Notify(s.onNotify)

	return s
}

// Connect connects the slot to the signal.
func (s *Signal[T]) Connect(slot func(T)) {
	s.slots = append(s.slots, slot)
}

// Emit emits the value through the signal.
func (s *Signal[T]) Emit(value T) {
	s.event.Emit(value)
}

// Close shut downs the signal.
func (s *Signal[T]) Close() {
	s.event.Close()
}

func (s *Signal[T]) onNotify(value T) {
	for _, slot := range s.slots {
		slot(value)
	}
}
