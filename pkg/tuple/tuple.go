package tuple

import (
	"math"
	"reflect"
)

const EPSILON float64 = 0.00001

type TupleLike interface {
	ToTuple() Tuple
}

type Tuple struct {
	X, Y, Z, W float64
}

func NewTuple(x, y, z, w float64) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: w}
}

func (t1 Tuple) Equals(t2 Tuple) bool {
	equalX := math.Abs(t1.X-t2.X) < EPSILON
	equalY := math.Abs(t1.Y-t2.Y) < EPSILON
	equalZ := math.Abs(t1.Z-t2.Z) < EPSILON
	equalW := math.Abs(t1.W-t2.W) < EPSILON
	return equalX && equalY && equalZ && equalW
}

// Create New Tuple From TupleLike
func NewTupleFromTL(elemParams ...TupleLike) Tuple {
	// If no paramenters recived
	if len(elemParams) == 0 {
		return Tuple{}
	}
	tupleLike := elemParams[0]
	if reflect.TypeOf(tupleLike).String() == "tuple.Vector" {
		return tupleLike.ToTuple()
	} else if reflect.TypeOf(tupleLike).String() == "tuple.Point" {
		return tupleLike.ToTuple()
	}
	return Tuple{}
}

func (t Tuple) WriteElem(row uint8, num float64) Tuple {
	tArray := t.ToArray()
	tArray[row] = num
	return NewTuple(tArray[0], tArray[1], tArray[2], tArray[3])
}

func (t Tuple) ToTupleLike() TupleLike {
	if t.W == 0.0 {
		return NewVector(t.X, t.Y, t.Z)
	} else if t.W == 1.0 {
		return NewPoint(t.X, t.Y, t.Z)
	}
	return NewVector(0, 0, 0)
}

func (t Tuple) ToArray() [4]float64 {
	return [4]float64{t.X, t.Y, t.Z, t.W}
}
