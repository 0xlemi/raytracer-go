package set

import "testing"

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


