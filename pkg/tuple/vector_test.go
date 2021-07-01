package tuple

import (
	"math"
	"testing"
)

func TestVectorParameterValues(t *testing.T) {
	tuple := Vector{X: 1.0, Y: 2.5, Z: 3.0}
	expected := [4]float64{1.0, 2.5, 3.0}
	actual := [4]float64{tuple.X, tuple.Y, tuple.Z}

	if actual != expected {
		t.Errorf("Vector Parameter Values test error. Expected %v, got %v", expected, actual)
	}
}

func TestVectorConstructor(t *testing.T) {
	expected := Vector{X: 1.1, Y: 2.5, Z: 3.0}
	actual := NewVector(1.1, 2.5, 3.0)

	if actual != expected {
		t.Errorf("Vector constructor NewVector() test error. Expected %v, got %v", expected, actual)
	}
}

// Check equality of Tuples.
// Because sometimes floats can generate round-off errors.
// That negate valid equalities. Where precision is [0.00001]
func TestVectorEquals(t *testing.T) {
	vectorA := NewVector(1.0, -2.500001, 3.00000009)
	vectorB := NewVector(1.0000002, -2.5, 3.0)

	expected := true
	actual := vectorA.Equals(vectorB)

	if actual != expected {
		t.Errorf("Vector Equals() func test error. Expected %v, got %v", expected, actual)
	}
}

func TestVectorsAdd(t *testing.T) {
	vectorA := NewVector(3, -2, 5)
	vectorB := NewVector(-2, 3, 1)

	expected := NewVector(1, 1, 6)
	actual := vectorA.Add(vectorB)

	if actual != expected {
		t.Errorf("TestAdd Vector test error. Expected %v, got %v", expected, actual)
	}
}

func TestSubtractTwoVectors(t *testing.T) {
	expected := NewVector(-2, -4, -6)
	actual := NewVector(3, 2, 1).Sub(NewVector(5, 6, 7))

	if actual != expected {
		t.Errorf("TestSubtractTwoVectors test error. Expected %v, got %v", expected, actual)
	}
}

func TestSubtractingVectorFromZeroVector(t *testing.T) {
	expected := NewVector(-2, -4, -6)
	actual := NewVector(0, 0, 0).Sub(NewVector(2, 4, 6))

	if actual != expected {
		t.Errorf("TestSubtractTwoVectors test error. Expected %v, got %v", expected, actual)
	}
}

func TestNegatingVector(t *testing.T) {
	expected := NewVector(1, -2, 3).Negate()
	actual := NewVector(-1, 2, -3)

	if actual != expected {
		t.Errorf("TestNegatingVector test error. Expected %v, got %v", expected, actual)
	}
}

func TestVectorScalarMultiplication(t *testing.T) {
	expected := NewVector(-2, -4, 6)
	actual := NewVector(-1, -2, 3).Multi(2)

	if actual != expected {
		t.Errorf("TestVectorScalarMultiplication test error. Expected %v, got %v", expected, actual)
	}
}

func TestVectorScalarDivision(t *testing.T) {
	expected := NewVector(3, 2.5, -14)
	actual := NewVector(6, 5, -28).Div(2)

	if actual != expected {
		t.Errorf("TestVectorDivision test error. Expected %v, got %v", expected, actual)
	}
}

func TestMagnitudeVector100(t *testing.T) {
	expected := 1.0
	actual := NewVector(1, 0, 0).Magnitude()

	if actual != expected {
		t.Errorf("TestMagnitudeVector100 test error. Expected %v, got %v", expected, actual)
	}
}

func TestMagnitudeVector010(t *testing.T) {
	expected := 1.0
	actual := NewVector(0, 1, 0).Magnitude()

	if actual != expected {
		t.Errorf("TestMagnitude010 test error. Expected %v, got %v", expected, actual)
	}
}

func TestMagnitudeVector001(t *testing.T) {
	expected := 1.0
	actual := NewVector(0, 0, 1).Magnitude()

	if actual != expected {
		t.Errorf("TestMagnitude001 test error. Expected %v, got %v", expected, actual)
	}
}

func TestMagnitudeVector123(t *testing.T) {
	expected := math.Sqrt(14)
	actual := NewVector(1, 2, 3).Magnitude()

	if actual != expected {
		t.Errorf("TestMagnitude123 test error. Expected %v, got %v", expected, actual)
	}
}

func TestMagnitudeVectorNeg123(t *testing.T) {
	expected := math.Sqrt(14)
	actual := NewVector(-1, -2, -3).Magnitude()

	if actual != expected {
		t.Errorf("TestMagnitudeNeg123 test error. Expected %v, got %v", expected, actual)
	}
}

func TestNormalizeVector400(t *testing.T) {
	expected := NewVector(1, 0, 0)
	actual := NewVector(4, 0, 0).Normalize()

	if actual != expected {
		t.Errorf("TestNormalizeVector400 test error. Expected %v, got %v", expected, actual)
	}
}

func TestNormalizeVector123(t *testing.T) {
	expected := NewVector(0.26726, 0.53452, 0.80178)
	actual := NewVector(1, 2, 3).Normalize()

	if !actual.Equals(expected) {
		t.Errorf("TestNormalizeVector123 test error. Expected %v, got %v", expected, actual)
	}
}

func TestMagnitudeOfNormalizedVector(t *testing.T) {
	expected := 1.0
	actual := NewVector(1, 2, 3).Normalize().Magnitude()

	if actual != expected {
		t.Errorf("TestMagnitudeOfNormalizedVector test error. Expected %v, got %v", expected, actual)
	}
}

func TestDotProduct123and234(t *testing.T) {
	expected := 20.0
	actual := NewVector(1, 2, 3).Dot(NewVector(2, 3, 4))

	if actual != expected {
		t.Errorf("TestDotProduct123and234 test error. Expected %v, got %v", expected, actual)
	}
}

func TestCrossProduct1(t *testing.T) {
	expected := NewVector(-1, 2, -1)
	actual := NewVector(1, 2, 3).Cross(NewVector(2, 3, 4))

	if actual != expected {
		t.Errorf("TestCrossProduct1 test error. Expected %v, got %v", expected, actual)
	}
}

func TestCrossProduct2(t *testing.T) {
	expected := NewVector(1, -2, 1)
	actual := NewVector(2, 3, 4).Cross(NewVector(1, 2, 3))

	if actual != expected {
		t.Errorf("TestCrossProduct2 test error. Expected %v, got %v", expected, actual)
	}
}
