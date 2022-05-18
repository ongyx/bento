package bento

import (
	"image"

	"gonum.org/v1/gonum/floats"
)

const (
	// Linear specifies a delta is constant.
	Linear DeltaAlgorithm = iota
	// Exponential specifies a delta in exponential (e^x) space.
	Exponential
)

// DeltaAlgorithm specifies the algorithm to use when generating deltas.
type DeltaAlgorithm int

// Delta is a delta that changes over time.
type Delta struct {
	delta image.Point
	timer *Timer

	index, limit int

	dx, dy []float64
}

// NewDelta creates an delta with the total delta, and the period over which to increase the current delta.
func NewDelta(
	algo DeltaAlgorithm,
	delta image.Point,
	period float64,
) *Delta {
	t := NewTimer(period, true)

	dt := t.Delta()
	// TODO(ongyx): a more efficient way to store the delta per tick?
	dx := make([]float64, dt)
	dy := make([]float64, dt)

	// some algorithms require the start to be at least 1, so add 1 to the delta here.
	delta.X += 1
	delta.Y += 1

	x := float64(delta.X)
	y := float64(delta.Y)

	var algoFunc func(buf []float64, from, to float64) []float64

	switch algo {
	case Linear:
		algoFunc = floats.Span
	case Exponential:
		algoFunc = floats.LogSpan
	}

	algoFunc(dx, 1, x)
	algoFunc(dy, 1, y)

	return &Delta{
		delta: delta,
		timer: t,
		index: -1,
		limit: int(dt),
		dx:    dx,
		dy:    dy,
	}
}

// Update updates the delta.
func (d *Delta) Update() error {
	d.index++

	if d.index >= d.limit || d.timer.Done() {
		d.index = d.limit - 1
	}

	return nil
}

// Delta returns the current delta.
func (d *Delta) Delta() (x, y float64) {
	// special case: if x/y delta is 0, return 0 here too
	// otherwise it will return NaN
	if d.delta.X != 0 {
		x = d.dx[d.index] - 1
	}

	if d.delta.Y != 0 {
		y = d.dy[d.index] - 1
	}

	return x, y
}

func (d *Delta) DeltaPt() image.Point {
	x, y := d.Delta()
	return image.Pt(int(x), int(y))
}

// Done checks if the current delta is equal to the total delta.
func (d *Delta) Done() bool {
	return d.index == (d.limit - 1)
}
