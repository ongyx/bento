package bento

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	tps = ebiten.DefaultTPS
)

// Bound calculates a bound, given its top-left point and its size.
func Bound(point, size image.Point) image.Rectangle {
	return image.Rectangle{Min: point, Max: point.Add(size)}
}

// Pad adds padding to the bound by a fixed amount.
func Pad(bound image.Rectangle, pad image.Point) image.Rectangle {
	return image.Rectangle{
		Min: bound.Min.Sub(pad),
		Max: bound.Max.Add(pad),
	}
}

// Unpad removes padding from a bound by a fixed amount.
func Unpad(bound image.Rectangle, pad image.Point) image.Rectangle {
	return image.Rectangle{
		Min: bound.Min.Add(pad),
		Max: bound.Max.Sub(pad),
	}
}

// Radian converts an angle in degrees to radians.
func Radian(degree float64) float64 {
	return degree * (math.Pi / 180)
}

// DPIScale scales the given resolution by the device's scale factor.
// This allows high-DPI rendering.
func DPIScale(res int) float64 {
	return float64(res) * ebiten.DeviceScaleFactor()
}

// Coordinate returns the given point as a float64 pair.
func Coordinate(point image.Point) (x, y float64) {
	return float64(point.X), float64(point.Y)
}

// Point returns the point of a geometry matrix.
func Point(m *ebiten.GeoM) image.Point {
	return image.Pt(int(m.Element(0, 2)), int(m.Element(1, 2)))
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
