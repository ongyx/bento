package bento

// event: provides utility functions for using channels as events/notifiers.

import "sync"

// Poll attempts to receive a value from the event without blocking the current goroutine.
// If the event is closed or empty, nil is returned.
func Poll[T any](ch <-chan T) *T {
	var cv *T

	select {
	case v, ok := <-ch:
		if ok {
			cv = &v
		}
	// don't block if there are no values
	default:
	}

	return cv
}

// Send asynchronously sends a value to the channel,
// notifying the waitgroup once the value has been sent.
func Send[T any](ch chan<- T, v T, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		ch <- v
		wg.Done()
	}()
}
