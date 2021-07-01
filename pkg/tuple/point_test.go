package tuple

import (
	"testing"
)

func TestPointParameterValues(t *testing.T) {
	tuple := Point{X: 1.0, Y: 2.5, Z: 3.0}
	expected := [4]float64{1.0, 2.5, 3.0}
	actual := [4]float64{tuple.X, tuple.Y, tuple.Z}

	if actual != expected {
		t.Errorf("Point Parameter Values test error. Expected %v, got %v", expected, actual)
	}
}

func TestPointConstructor(t *testing.T) {
	expected := Point{X: 1.1, Y: 2.5, Z: 3.0}
	actual := NewPoint(1.1, 2.5, 3.0)

	if actual != expected {
		t.Errorf("Point constructor NewTuple() test error. Expected %v, got %v", expected, actual)
	}
}

// Check equality of Tuples.
// Because sometimes floats can generate round-off errors.
// That negate valid equalities. Where precision is [0.00001]
func TestPointEquals(t *testing.T) {
	pointA := NewPoint(1.0, -2.500001, 3.00000009)
	pointB := NewPoint(1.0000002, -2.5, 3.0)

	expected := true
	actual := pointA.Equals(pointB)

	if actual != expected {
		t.Errorf("Point Equals() func test error. Expected %v, got %v", expected, actual)
	}
}

func TestAddVectorToPoint(t *testing.T) {
	vector := NewVector(3, -2, 5)
	point := NewPoint(-2, 3, 1)

	expected := NewPoint(1, 1, 6)
	actual := point.AddVector(vector)

	if actual != expected {
		t.Errorf("TestAddVectorToPoint test error. Expected %v, got %v", expected, actual)
	}
}

func TestSubtractVectorToPoint(t *testing.T) {
	expected := NewPoint(-2, -4, -6)
	actual := NewPoint(3, 2, 1).SubVector(NewVector(5, 6, 7))

	if actual != expected {
		t.Errorf("TestSubtractVectorToPoint test error. Expected %v, got %v", expected, actual)
	}
}

func TestVectorBetweenPoints(t *testing.T) {
	expected := NewVector(-2, -4, -6)
	actual := NewPoint(5, 6, 7).VectorBetweenPoints(NewPoint(3, 2, 1))

	if actual != expected {
		t.Errorf("TestVectorBetweenPoints test error. Expected %v, got %v", expected, actual)
	}
}

func TestPointScalarMultiplication(t *testing.T) {
	expected := NewPoint(-2, -4, 6)
	actual := NewPoint(-1, -2, 3).Multi(2)

	if actual != expected {
		t.Errorf("TestPointScalarMultiplication test error. Expected %v, got %v", expected, actual)
	}
}

func TestPointScalarDivision(t *testing.T) {
	expected := NewPoint(3, 2.5, -14)
	actual := NewPoint(6, 5, -28).Div(2)

	if actual != expected {
		t.Errorf("TestPointScalarDivision test error. Expected %v, got %v", expected, actual)
	}
}
