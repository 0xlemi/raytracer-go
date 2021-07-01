package tuple

import (
	"math"
)

const EPSILON float64 = 0.00001

type Vector struct {
	X, Y, Z float64
}

func NewVector(x float64, y float64, z float64) Vector {
	return Vector{x, y, z}
}

func (v1 Vector) Equals(v2 Vector) bool {
	equalX := math.Abs(v1.X-v2.X) < EPSILON
	equalY := math.Abs(v1.Y-v2.Y) < EPSILON
	equalZ := math.Abs(v1.Z-v2.Z) < EPSILON
	return equalX && equalY && equalZ
}

// Add 2 Vectors
func (v1 Vector) Add(v2 Vector) Vector {
	v1.X += v2.X
	v1.Y += v2.Y
	v1.Z += v2.Z
	return v1
}

// Subtracting 2 Vectors.
func (v1 Vector) Sub(v2 Vector) Vector {
	v1.X -= v2.X
	v1.Y -= v2.Y
	v1.Z -= v2.Z
	return v1
}

// Multilpy tuple for scalar
func (v Vector) Multi(s float64) Vector {
	v.X *= s
	v.Y *= s
	v.Z *= s
	return v
}

// Divide tuple for scalar
func (v Vector) Div(s float64) Vector {
	v.X /= s
	v.Y /= s
	v.Z /= s
	return v
}

func (v Vector) Negate() Vector {
	return v.Multi(-1)
}

func (v Vector) Magnitude() float64 {
	total := 0.0
	nums := []float64{v.X, v.Y, v.Z}

	for _, num := range nums {
		total += num * num
	}
	return math.Sqrt(total)
}

func (v Vector) Normalize() Vector {
	magnitud := v.Magnitude()
	v.X /= magnitud
	v.Y /= magnitud
	v.Z /= magnitud

	return v
}

func (v1 Vector) Dot(v2 Vector) float64 {
	x := (v1.X * v2.X)
	y := (v1.Y * v2.Y)
	z := (v1.Z * v2.Z)

	return x + y + z
}

func (v1 Vector) Cross(v2 Vector) Vector {
	x := (v1.Y * v2.Z) - (v1.Z * v2.Y)
	y := (v1.Z * v2.X) - (v1.X * v2.Z)
	z := (v1.X * v2.Y) - (v1.Y * v2.X)

	return NewVector(x, y, z)
}
