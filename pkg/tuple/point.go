package tuple

import (
	"math"
)

type Point struct {
	X, Y, Z float64
}

func NewPoint(x float64, y float64, z float64) Point {
	return Point{x, y, z}
}

func (p1 Point) Equals(p2 Point) bool {
	equalX := math.Abs(p1.X-p2.X) < EPSILON
	equalY := math.Abs(p1.Y-p2.Y) < EPSILON
	equalZ := math.Abs(p1.Z-p2.Z) < EPSILON
	return equalX && equalY && equalZ
}

// Get point on the tip of the vector origining from p
func (p Point) AddVector(v Vector) Point {
	p.X += v.X
	p.Y += v.Y
	p.Z += v.Z
	return p
}

// Get point on the tip of the vector origining from p
func (p Point) SubVector(v Vector) Point {
	p.X -= v.X
	p.Y -= v.Y
	p.Z -= v.Z
	return p
}

// Vector that goes from p1 to p2
func (p1 Point) VectorBetweenPoints(p2 Point) Vector {
	x := p2.X - p1.X
	y := p2.Y - p1.Y
	z := p2.Z - p1.Z
	return NewVector(x, y, z)
}

// Multilpy tuple for scalar
func (p Point) Multi(s float64) Point {
	p.X *= s
	p.Y *= s
	p.Z *= s
	return p
}

// Divide tuple for scalar
func (p Point) Div(s float64) Point {
	p.X /= s
	p.Y /= s
	p.Z /= s
	return p
}
