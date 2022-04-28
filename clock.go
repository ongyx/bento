package bento

var (
	// Clock is a monotonically increasing tick counter that can be used to schedule operarations (see the Timer type).
	// There is only one instance per game, and tick updates are handled by the stage.
	// The tick counter always starts from zero.
	Clock = clock{}
)

type clock struct {
	tick int
}

// Now returns the current tick.
func (c *clock) Now() int {
	return c.tick
}

func (c *clock) increment() {
	c.tick++
}
