package bento

import (
	"fmt"
)

// Timer is a tick-based scheduler for operations.
// One tick is equivalent to a single frame, where 1 second is 60 frames.
type Timer struct {
	delta, tick, count uint64

	once bool
}

// NewTimer creates a new timer that triggers every n seconds.
// If once is true, the timer will only trigger once.
func NewTimer(n float64, once bool) *Timer {
	if n <= 0 {
		panic(fmt.Sprintf("timer: negative n %f", n))
	}

	if d := SecondToTick(n); d == 0 {
		panic(fmt.Sprintf("timer: duration of n %f too small", n))
	} else {
		return &Timer{delta: uint64(d), once: once}
	}
}

// Tick increments the timer.
func (t *Timer) Tick() {
	t.tick++
}

// Delta returns the number of ticks between each trigger.
func (t *Timer) Delta() uint64 {
	return t.delta
}

// Count returns the number of times the timer has triggered.
func (t *Timer) Count() uint64 {
	return t.count
}

// Done checks if the timer has triggered.
func (t *Timer) Done() bool {
	if t.tick != 0 && (t.tick%t.delta) == 0 {
		if t.once {
			return t.count == 0
		}

		t.count++

		return true
	}

	return false
}
