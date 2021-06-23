package set

import (
	"math"
)

type Color struct {
	Red, Blue, Green float64
}

func NewColor(r, g, b float64) Color {
	return Color{Red: r, Green: g, Blue: b}
}

func (c Color) Add(c2 Color) Color {
	c.Red += c2.Red
	c.Blue += c2.Blue
	c.Green += c2.Green
	return c
}

func (c Color) Sub(c2 Color) Color {
	c.Red -= c2.Red
	c.Blue -= c2.Blue
	c.Green -= c2.Green
	return c
}

func (c Color) Multi(s float64) Color {
	c.Red *= s
	c.Blue *= s
	c.Green *= s
	return c
}

func (c Color) Hadamard(c2 Color) Color {
	c.Red *= c2.Red
	c.Blue *= c2.Blue
	c.Green *= c2.Green
	return c
}

func ColorEquals(a, b Color) bool {
	equalRed := math.Abs(a.Red-b.Red) < EPSILON
	equalGreen := math.Abs(a.Green-b.Green) < EPSILON
	equalBlue := math.Abs(a.Blue-b.Blue) < EPSILON
	return equalRed && equalGreen && equalBlue 
}