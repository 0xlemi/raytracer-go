package set

import (
	"errors"
	"math"
)

const EPSILON = 0.00001

type Tuple struct {
	X, Y, Z, W float64
}

// Tuple constructor
func NewTuple(x float64, y float64, z float64, w float64) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: w}
}

func (t Tuple) IsPoint() bool {
	return t.W == 1.0
}

func (t Tuple) IsVector() bool {
	return t.W == 0.0
}

func Point(x float64, y float64, z float64) Tuple {
	return NewTuple(x, y, z, 1.0)
}

func Vector(x float64, y float64, z float64) Tuple {
	return NewTuple(x, y, z, 0.0)
}

func Equals(a, b Tuple) bool {
	equalX := math.Abs(a.X-b.X) < EPSILON
	equalY := math.Abs(a.Y-b.Y) < EPSILON
	equalZ := math.Abs(a.Z-b.Z) < EPSILON
	equalW := math.Abs(a.W-b.W) < EPSILON
	return equalX && equalY && equalZ && equalW
}

// Add Tuples. Making sure W does the apropriat behaviour
// Sum of 2 vectors give a vertor. vertor point give point.
// Sum of 2 points makes no sense.
func (t Tuple) Add(b Tuple) Tuple {
	if t.IsPoint() && b.IsPoint() {
		err := errors.New("YOU CANNOT ADD 2 POINTS")
		panic(err)
	}
	t.X += b.X
	t.Y += b.Y
	t.Z += b.Z
	t.W += b.W
	return t
}

// Subtracting Tuples.
func (t Tuple) Sub(b Tuple) Tuple {
	if t.IsVector() && b.IsPoint() {
		err := errors.New("YOU CANNOT SUBTRACT A VECTOR - POINT")
		panic(err)
	}
	t.X -= b.X
	t.Y -= b.Y
	t.Z -= b.Z
	t.W -= b.W
	return t
}

// Negating Tuple
// In vector case: vector pointing in the exact oposite direction
func (t Tuple) Negate() Tuple {
	return NewTuple(0, 0, 0, 0).Sub(t)
}

// Multilpy tuple for scalar
func (t Tuple) Multi(s float64) Tuple {
	t.X *= s
	t.Y *= s
	t.Z *= s
	t.W *= s
	return t
}

// Divide tuple for scalar
func (t Tuple) Div(s float64) Tuple {
	t.X /= s
	t.Y /= s
	t.Z /= s
	t.W /= s
	return t
}

func (t Tuple) Magnitude() float64 {
	if !t.IsVector() {
		err := errors.New("YOU CAN ONLY GET MAGNITUD OF A VECTOR")
		panic(err)
	}

	total := 0.0
	nums := []float64{t.X, t.Y, t.Z}

	for _, num := range nums {
		total += num * num
	}
	return math.Sqrt(total)
}

func (t Tuple) Normalize() Tuple {
	if !t.IsVector() {
		err := errors.New("YOU CAN ONLY NORMALIZE OF A VECTOR")
		panic(err)
	}

	magnitud := t.Magnitude()
	t.X /= magnitud
	t.Y /= magnitud
	t.Z /= magnitud
	t.W /= magnitud

	return t
}

func (t Tuple) Dot(v Tuple) float64 {
	if !(t.IsVector() && v.IsVector()) {
		err := errors.New("YOU CAN ONLY GET DOT PRODUCTS OF VECTORS")
		panic(err)
	}

	x := (t.X * v.X)
	y := (t.Y * v.Y)
	z := (t.Z * v.Z)
	w := (t.W * v.W)

	return x + y + z + w
}

func (t Tuple) Cross(v Tuple) Tuple {
  if !(t.IsVector() && v.IsVector()) {
		err := errors.New("YOU CAN ONLY GET DOT PRODUCTS OF VECTORS")
		panic(err)
	}

	x := (t.Y * v.Z) - ( t.Z * v.Y)
	y := (t.Z * v.X) - ( t.X * v.Z)
	z := (t.X * v.Y) - ( t.Y * v.X)

	return Vector(x, y, z)
}

