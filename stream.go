package bento

// Stream is a convinence wrapper around a channel.
// For example, it can be used to asynchronusly send events:
//
//	type event string
//
//	// to use an existing channel, Stream{C: <channel>} can be used instead.
// 	s := NewStream[event](0)
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
type Stream[T any] struct {
	C chan T
}

// NewStream creates a stream of type T.
// If size is more than 1, the underlying channel will have a buffer of that length.
// Otherwise, the channel is unbuffered.
func NewStream[T any](size int) Stream[T] {
	var ch chan T

	if size <= 1 {
		ch = make(chan T)
	} else {
		ch = make(chan T, size)
	}

	return Stream[T]{C: ch}
}

// Read blocks the current goroutine until a value can be read.
// If the stream is closed, nil is returned.
func (s *Stream[T]) Read() *T {
	if v, ok := <-s.C; ok {
		return &v
	} else {
		return nil
	}
}

// Poll reads a value without blocking.
// If the stream is closed or empty, nil is returned.
func (s *Stream[T]) Poll() *T {
	var value *T

	select {
	case v, ok := <-s.C:
		if ok {
			value = &v
		}
	// don't block if there are no values
	default:
	}

	return value
}

// Write blocks the current goroutine until the value can be written.
func (s *Stream[T]) Write(value T) {
	s.C <- value
}

// Close closes the stream's channel.
// NOTE: You probably should only close the channel from the writer side.
func (s *Stream[T]) Close() {
	close(s.C)
}
