package tuple

import (
	"testing"
)

func TestTupleConstructor(t *testing.T) {
	expected := Tuple{X: 5.0, Y: 3.2, Z: 5.3, W: 0.0}
	actual := NewTuple(5.0, 3.2, 5.3, 0.0)

	if actual != expected {
		t.Errorf("TestTupleConstructor test error. Expected %v, got %v", expected, actual)
	}
}

func TestTulpeEquals(t *testing.T) {
	pointA := NewTuple(1.0, -2.500001, 3.00000009, 14.005)
	pointB := NewTuple(1.0000002, -2.5, 3.0, 14.00500000001)

	expected := true
	actual := pointA.Equals(pointB)

	if actual != expected {
		t.Errorf("TestTupleEquals func test error. Expected %v, got %v", expected, actual)
	}
}

func TestTupleConstructorTupleLikeNil(t *testing.T) {
	expected := Tuple{X: 0.0, Y: 0.0, Z: 0.0, W: 0.0}
	actual := NewTupleFromTL()

	if actual != expected {
		t.Errorf("TestTupleConstructorTupleLike test error. Expected %v, got %v", expected, actual)
	}
}

func TestTupleConstructorTupleLikeVector(t *testing.T) {
	expected := Tuple{X: 5.0, Y: 3.2, Z: 5.3, W: 0.0}
	actual := NewTupleFromTL(NewVector(5, 3.2, 5.3))

	if actual != expected {
		t.Errorf("TestTupleConstructorTupleLikeVector test error. Expected %v, got %v", expected, actual)
	}
}

func TestTupleConstructorTupleLikePoint(t *testing.T) {
	expected := Tuple{X: 5.0, Y: 3.2, Z: 5.3, W: 1.0}
	actual := NewTupleFromTL(NewPoint(5, 3.2, 5.3))

	if actual != expected {
		t.Errorf("TestTupleConstructorTupleLikePoint test error. Expected %v, got %v", expected, actual)
	}
}

func TestTupleToArray(t *testing.T) {
	expected := [4]float64{5.0, 3.5, 5.3, 1.0}
	actual := NewTupleFromTL(NewPoint(5, 3.5, 5.3)).ToArray()

	if actual != expected {
		t.Errorf("TestTupleToArray test error. Expected %v, got %v", expected, actual)
	}
}

func TestTupleToTupleLikeVector(t *testing.T) {
	expected := NewVector(4.23, 3, 13.2)
	actual := Tuple{X: 4.23, Y: 3, Z: 13.2, W: 0.0}.ToTupleLike()

	if actual != expected {
		t.Errorf("TestTupleToTupleLike test error. Expected %v, got %v", expected, actual)
	}
}

func TestTupleToTupleLikePoint(t *testing.T) {
	expected := NewPoint(4.23, 3, 13.2)
	actual := Tuple{X: 4.23, Y: 3, Z: 13.2, W: 1.0}.ToTupleLike()

	if actual != expected {
		t.Errorf("TestTupleToTupleLikePoint test error. Expected %v, got %v", expected, actual)
	}
}

func TestTupleToTupleLikeError(t *testing.T) {
	expected := NewVector(0, 0, 0)
	actual := Tuple{X: 4.23, Y: 3, Z: 13.2, W: 2.0}.ToTupleLike()

	if actual != expected {
		t.Errorf("TestTupleToTupleLikeError test error. Expected %v, got %v", expected, actual)
	}
}
