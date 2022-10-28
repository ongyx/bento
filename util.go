package bento

import (
	"image"
	"math"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	tps = ebiten.DefaultTPS
)

var (
	dpi     float64
	dpiSync sync.Once
)

// Radian converts an angle in degrees to radians.
func Radian(degree float64) float64 {
	return degree * (math.Pi / 180)
}

// SecondToTick converts seconds to ticks.
func SecondToTick(seconds float64) int {
	return int(seconds * tps)
}

// TickToSecond converts ticks to seconds.
func TickToSecond(ticks int) float64 {
	return float64(ticks) / tps
}

// Poll attempts to read a value from the channel without blocking.
// If the channel is empty, ok is false.
func Poll[T any](ch <-chan T) (value T, ok bool) {
	select {
	case value = <-ch:
		ok = true
	default:
	}

	return
}

// DPI returns the device scale factor.
func DPI() float64 {
	dpiSync.Do(func() {
		dpi = ebiten.DeviceScaleFactor()
	})

	return dpi
}

// ScaleDPI scales the point with the device scale factor.
func ScaleDPI(p image.Point) image.Point {
	f := DPI()
	x := float64(p.X) * f
	y := float64(p.Y) * f

	return image.Point{int(x), int(y)}
}
