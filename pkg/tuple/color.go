package tuple

import (
	"math"
)

type Color struct {
	Red, Green, Blue float64
}

func NewColor(r, g, b float64) Color {
	return Color{Red: r, Green: g, Blue: b}
}

func (c Color) Add(c2 Color) Color {
	c.Red += c2.Red
	c.Green += c2.Green
	c.Blue += c2.Blue
	return c
}

func (c Color) Sub(c2 Color) Color {
	c.Red -= c2.Red
	c.Green -= c2.Green
	c.Blue -= c2.Blue
	return c
}

func (c Color) Multi(s float64) Color {
	c.Red *= s
	c.Green *= s
	c.Blue *= s
	return c
}

func (c Color) Hadamard(c2 Color) Color {
	c.Red *= c2.Red
	c.Green *= c2.Green
	c.Blue *= c2.Blue
	return c
}

func (c Color) Cap(max uint) Color {
	c.Red = MinMax(c.Red, 0, float64(max))
	c.Green = MinMax(c.Green, 0, float64(max))
	c.Blue = MinMax(c.Blue, 0, float64(max))
	return c
}

// If number gets over a threshold set to min or max
func MinMax(n float64, min float64, max float64) float64 {
	if n < min {
		return min
	} else if n > max {
		return max
	} else {
		return n
	}
}

func ColorEquals(a, b Color) bool {
	equalRed := math.Abs(a.Red-b.Red) < EPSILON
	equalGreen := math.Abs(a.Green-b.Green) < EPSILON
	equalBlue := math.Abs(a.Blue-b.Blue) < EPSILON
	return equalRed && equalGreen && equalBlue
}
