package bento

// Event notifies one or more callbacks when a event value is sent over the channel.
//
//	ch := make(chan string)
//
//	e := NewEvent(ch)
//
//	e.Notify(func(s string) {
//		fmt.Println(s)
//	})
//
//	e.Listen()
//
//	ch <- "Hello World!"
type Event[T any] struct {
	ch  <-chan T
	fns []func(T)
}

// NewEvent creates a new event backed by the read-only channel ch.
func NewEvent[T any](ch <-chan T) *Event[T] {
	e := &Event[T]{ch: ch}
	go e.listen()

	return e
}

// Notify registers the callback to the event.
func (e *Event[T]) Notify(fn func(T)) {
	e.fns = append(e.fns, fn)
}

func (e *Event[T]) listen() {
	for v := range e.ch {
		for _, fn := range e.fns {
			fn(v)
		}
	}
}
