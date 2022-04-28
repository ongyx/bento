package bento

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	tps = ebiten.DefaultTPS
)

// Timer is a tick-based scheduler for operations.
// One tick is equivalent to a single frame, where 1 second is 60 frames.
type Timer struct {
	delta int
	once  bool

	base  int
	count int
}

// NewTimer creates a new timer that triggers every n seconds.
// If once is true, the timer will only trigger once.
func NewTimer(n float64, once bool) *Timer {
	if n <= 0 {
		panic(fmt.Sprintf("timer: negative n (%f)", n))
	}

	if d := SecondToTick(n); d == 0 {
		panic(fmt.Sprintf("timer: duration of n too small (%f)", n))
	} else {
		return &Timer{delta: d, once: once, base: Clock.Now()}
	}
}

// Delta returns the number of ticks between each trigger.
func (t *Timer) Delta() int {
	return t.delta
}

// Done checks if the timer has triggered.
func (t *Timer) Done() bool {
	ts := Clock.Now() - t.base

	if ts != 0 && (ts%t.delta) == 0 {
		if t.once {
			return t.count == 0
		}

		t.count++

		return true
	}

	return false
}
