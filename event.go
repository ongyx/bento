package bento

type Handler[T any] func(T) error

type Event[T any] struct {
	fn Handler[T]

	stream Stream[T]
	status chan error
}

func NewEvent[T any](fn Handler[T]) *Event[T] {
	return &Event[T]{
		fn:     fn,
		stream: NewStream[T](0),
		status: make(chan error),
	}
}

func (e *Event[T]) Start() {
	go func() {
		for v := range e.stream.C {
			e.status <- e.fn(v)
		}
	}()
}
