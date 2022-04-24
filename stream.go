package bento

// Stream is a wrapper around a two-way unbuffered channel.
// For example, it can be used to asynchronusly send events:
//
//	type event string
//
// 	s := NewStream[event]()
//
//	// this closure acts as a generator
//	// s.Write behaves like 'yield': it blocks until another goroutine reads from the stream.
// 	go func() {
// 		s.Write("some event")
// 		s.Close()
// 	}()
//
// 	if e := s.Read(); e != nil {
// 		fmt.Printf("event: %s", *e)
// 	}
//
// Streams should never be instantiated as a struct literal: it will have a nil channel, which deadlocks any read/writes.
// Instead use the NewStream function.
type Stream[T any] struct {
	ch chan T
}

// NewStream creates a stream with values of type T.
func NewStream[T any]() Stream[T] {
	return Stream[T]{ch: make(chan T)}
}

// Read returns the next value from the stream.
// If the stream is closed, nil is returned.
func (s *Stream[T]) Read() *T {
	if v, ok := <-s.ch; ok {
		return &v
	} else {
		return nil
	}
}

// Poll returns the next value from the stream without blocking.
// If the stream is closed or there are no values in the stream, nil is returned.
func (s *Stream[T]) Poll() *T {
	var value *T

	select {
	case v, ok := <-s.ch:
		if ok {
			value = &v
		}
	// don't block if there are no values
	default:
	}

	return value
}

// Write puts the value into the stream.
func (s *Stream[T]) Write(value T) {
	s.ch <- value
}

// Close closes the stream's channel.
// NOTE: You probably should only close the channel from the writer side.
func (s *Stream[T]) Close() {
	close(s.ch)
}
