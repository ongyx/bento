package bento

import "sync/atomic"

// Clock is a monotonically increasing tick counter.
var Clock = clock{}

type clock struct {
	tick uint64
}

// Now returns the current tick of the clock.
func (c *clock) Now() uint64 {
	return c.tick
}

func (c *clock) increment() {
	atomic.AddUint64(&c.tick, 1)
}
