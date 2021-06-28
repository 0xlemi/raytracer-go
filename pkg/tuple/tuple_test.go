package tuple

import (
	// "fmt"
	"math"
	"testing"
)

func TestParameterValues(t *testing.T) {
	tuple := Tuple{X: 1.0, Y: 2.5, Z: 3.0, W: 0.0}
	expected := [4]float64{1.0, 2.5, 3.0, 0.0}
	actual := [4]float64{tuple.X, tuple.Y, tuple.Z, tuple.W}

	if actual != expected {
		t.Errorf("Tuple Parameter Values test error. Expected %v, got %v", expected, actual)
	}
}

func TestConstructor(t *testing.T) {
	expected := Tuple{X: 1.1, Y: 2.5, Z: 3.0, W: 0.0}
	actual := NewTuple(1.1, 2.5, 3.0, 0.0)

	if actual != expected {
		t.Errorf("Tuple constructor NewTuple() test error. Expected %v, got %v", expected, actual)
	}
}

func TestIsPoint(t *testing.T) {
	// Is a point
	expected := true
	actual := NewTuple(1.0, 2.5, 3.0, 1.0).IsPoint()

	if actual != expected {
		t.Errorf("Tuple isPoint test error. Expected %v, got %v", expected, actual)
	}

	// Not a point
	expected = false
	actual = NewTuple(1.0, 2.5, 3.0, 0.0).IsPoint()

	if actual != expected {
		t.Errorf("Tuple isPoint test error. Expected %v, got %v", expected, actual)
	}
}

func TestIsVector(t *testing.T) {
	// Is a vector
	expected := true
	actual := NewTuple(1.0, -2.5, 3.0, 0.0).IsVector()

	if actual != expected {
		t.Errorf("Tuple isVector test error. Expected %v, got %v", expected, actual)
	}

	// Not a vector
	expected = false
	actual = NewTuple(1.0, -2.5, 3.0, 1.0).IsVector()

	if actual != expected {
		t.Errorf("Tuple isVector test error. Expected %v, got %v", expected, actual)
	}
}

// Create point Test
func TestPoint(t *testing.T) {
	expected := NewTuple(3, 8, 9, 1)
	actual := Point(3, 8, 9)
	if actual != expected {
		t.Errorf("Point() func test error. Expected %v, got %v", expected, actual)
	}
}

// Create point Test
func TestVector(t *testing.T) {
	expected := NewTuple(3, 8, 9, 0)
	actual := Vector(3, 8, 9)
	if actual != expected {
		t.Errorf("Vector() func test error. Expected %v, got %v", expected, actual)
	}
}

// Check equality of Tuples.
// Because sometimes floats can generate round-off errors.
// That negate valid equalities. Where precision is [0.00001]
func TestEquals(t *testing.T) {
	tupleA := NewTuple(1.0, -2.500001, 3.00000009, 0.0)
	tupleB := NewTuple(1.0000002, -2.5, 3.0, 0.0)

	expected := true
	actual := Equals(tupleA, tupleB)

	if actual != expected {
		t.Errorf("Equals() func test error. Expected %v, got %v", expected, actual)
	}
}

func TestAdd(t *testing.T) {
	tupleA := NewTuple(3, -2, 5, 1)
	tupleB := NewTuple(-2, 3, 1, 0)

	expected := NewTuple(1, 1, 6, 1)
	actual := tupleA.Add(tupleB)

	if actual != expected {
		t.Errorf("TestAdd test error. Expected %v, got %v", expected, actual)
	}
}

// You should not be able to add 2 tuples that represent points.
func TestPanicAddTwoPoints(t *testing.T) {
	defer func() { recover() }()

	pointA := Point(3, -2, 5)
	pointB := Point(-2, 3, 1)

	pointA.Add(pointB)

	t.Errorf("TestPanicAddTwoPoints test error. Should have panicked")
}

func TestSubtractTwoPoints(t *testing.T) {
	expected := Vector(-2, -4, -6)
	actual := Point(3, 2, 1).Sub(Point(5, 6, 7))

	if actual != expected {
		t.Errorf("TestSubtractTwoPoints test error. Expected %v, got %v", expected, actual)
	}
}

func TestSubtractPointAndVector(t *testing.T) {
	expected := Point(-2, -4, -6)
	actual := Point(3, 2, 1).Sub(Vector(5, 6, 7))

	if actual != expected {
		t.Errorf("TestSubtractPointAndVector test error. Expected %v, got %v", expected, actual)
	}
}

func TestPanicSubtractVectorAndPoint(t *testing.T) {
	defer func() { recover() }()

	vector := Vector(3, 2, 1)
	point := Point(5, 6, 7)

	vector.Sub(point)

	t.Errorf("TestPanicSubtractVectorAndPoint test error. Should have panicked")
}

func TestSubtractTwoVectors(t *testing.T) {
	expected := Vector(-2, -4, -6)
	actual := Vector(3, 2, 1).Sub(Vector(5, 6, 7))

	if actual != expected {
		t.Errorf("TestSubtractTwoVectors test error. Expected %v, got %v", expected, actual)
	}
}

func TestSubtractingVectorFromZeroVector(t *testing.T) {
	expected := Vector(-2, -4, -6)
	actual := Vector(0, 0, 0).Sub(Vector(2, 4, 6))

	if actual != expected {
		t.Errorf("TestSubtractTwoVectors test error. Expected %v, got %v", expected, actual)
	}
}

func TestNegatingTuple(t *testing.T) {
	expected := NewTuple(1, -2, 3, -4).Negate()
	actual := NewTuple(-1, 2, -3, 4)

	if actual != expected {
		t.Errorf("TestNegatingTuple test error. Expected %v, got %v", expected, actual)
	}
}

func TestScalarMultiplication(t *testing.T) {
	expected := NewTuple(-2, -4, 6, 4)
	actual := NewTuple(-1, -2, 3, 2).Multi(2)

	if actual != expected {
		t.Errorf("TestScalarMultiplication test error. Expected %v, got %v", expected, actual)
	}
}

func TestScalarDivision(t *testing.T) {
	expected := NewTuple(3, 2.5, -14, 6)
	actual := NewTuple(6, 5, -28, 12).Div(2)

	if actual != expected {
		t.Errorf("TestScalarDivision test error. Expected %v, got %v", expected, actual)
	}
}

func TestMagnitudeVector100(t *testing.T) {
	expected := 1.0
	actual := Vector(1, 0, 0).Magnitude()

	if actual != expected {
		t.Errorf("TestMagnitudeVector100 test error. Expected %v, got %v", expected, actual)
	}
}

func TestMagnitudeVector010(t *testing.T) {
	expected := 1.0
	actual := Vector(0, 1, 0).Magnitude()

	if actual != expected {
		t.Errorf("TestMagnitude010 test error. Expected %v, got %v", expected, actual)
	}
}

func TestMagnitudeVector001(t *testing.T) {
	expected := 1.0
	actual := Vector(0, 0, 1).Magnitude()

	if actual != expected {
		t.Errorf("TestMagnitude001 test error. Expected %v, got %v", expected, actual)
	}
}

func TestMagnitudeVector123(t *testing.T) {
	expected := math.Sqrt(14)
	actual := Vector(1, 2, 3).Magnitude()

	if actual != expected {
		t.Errorf("TestMagnitude123 test error. Expected %v, got %v", expected, actual)
	}
}

func TestMagnitudeVectorNeg123(t *testing.T) {
	expected := math.Sqrt(14)
	actual := Vector(-1, -2, -3).Magnitude()

	if actual != expected {
		t.Errorf("TestMagnitudeNeg123 test error. Expected %v, got %v", expected, actual)
	}
}

func TestPanicMagnitudeOfRandTuple(t *testing.T) {
	defer func() { recover() }()

	NewTuple(1, 2, 3, 4).Magnitude()

	t.Errorf("TestPanicMagnitudeOfRandTuple test error. Should have panicked")
}

func TestNormalizeVector400(t *testing.T) {
	expected := Vector(1, 0, 0)
	actual := Vector(4, 0, 0).Normalize()

	if actual != expected {
		t.Errorf("TestNormalizeVector400 test error. Expected %v, got %v", expected, actual)
	}
}

func TestNormalizeVector123(t *testing.T) {
	expected := Vector(0.26726, 0.53452, 0.80178)
	actual := Vector(1, 2, 3).Normalize()

	if !Equals(actual, expected) {
		t.Errorf("TestNormalizeVector123 test error. Expected %v, got %v", expected, actual)
	}
}

func TestMagnitudeOfNormalizedVector(t *testing.T) {
	expected := 1.0
	actual := Vector(1, 2, 3).Normalize().Magnitude()

	if actual != expected {
		t.Errorf("TestMagnitudeOfNormalizedVector test error. Expected %v, got %v", expected, actual)
	}
}

func TestPanicNormalizeRandTuple(t *testing.T) {
	defer func() { recover() }()

	NewTuple(1, 2, 3, 5).Normalize()

	t.Errorf("TestPanicNormalizeRandTuple test error. Should have panicked")
}

func TestDotProduct123and234(t *testing.T) {
	expected := 20.0
	actual := Vector(1, 2, 3).Dot(Vector(2, 3, 4))

	if actual != expected {
		t.Errorf("TestDotProduct123and234 test error. Expected %v, got %v", expected, actual)
	}
}

func TestPanicDotForNotVectors(t *testing.T) {
	defer func() { recover() }()

	Vector(1, 2, 3).Dot(Point(2, 3, 4))

	t.Errorf("TestPanicDotForNotVectors test error. Should have panicked")
}

func TestCrossProduct1(t *testing.T) {
	expected := Vector(-1, 2, -1)
	actual := Vector(1, 2, 3).Cross(Vector(2, 3, 4))

	if actual != expected {
		t.Errorf("TestCrossProduct1 test error. Expected %v, got %v", expected, actual)
	}
}

func TestCrossProduct2(t *testing.T) {
	expected := Vector(1, -2, 1)
	actual := Vector(2, 3, 4).Cross(Vector(1, 2, 3))

	if actual != expected {
		t.Errorf("TestCrossProduct2 test error. Expected %v, got %v", expected, actual)
	}
}

func TestPanicCrossProduct(t *testing.T) {
	defer func() { recover() }()

	Vector(1, 2, 3).Cross(Point(2, 3, 4))

	t.Errorf("TestPanicCrossProduct test error. Should have panicked")
}