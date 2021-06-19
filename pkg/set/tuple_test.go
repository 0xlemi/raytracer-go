package set

import "testing"

func TestTupleParameterValues(t *testing.T) {
	tuple := Tuple{X: 1.0, Y: 2.5, Z: 3.0, W: 0.0}
	expected := [4]float64{1.0, 2.5, 3.0, 0.0}
	actual := [4]float64{tuple.X, tuple.Y, tuple.Z, tuple.W}

	if actual != expected {
		t.Errorf("Tuple Parameter Values test error. Expected %v, got %v", expected, actual)
	}
}

func TestTupleIsPoint(t *testing.T) {
	// Is a point
	expected := true
	actual := Tuple{X: 1.0, Y: 2.5, Z: 3.0, W: 1.0}.IsPoint()

	if actual != expected {
		t.Errorf("Tuple isPoint test error. Expected %v, got %v", expected, actual)
	}

	// Not a point
	expected = false
	actual = Tuple{X: 1.0, Y: 2.5, Z: 3.0, W: 0.0}.IsPoint()

	if actual != expected {
		t.Errorf("Tuple isPoint test error. Expected %v, got %v", expected, actual)
	}
}

func TestTupleIsVector(t *testing.T) {
	// Is a vector
	expected := true
	actual := Tuple{X: 1.0, Y: -2.5, Z: 3.0, W: 0.0}.IsVector()

	if actual != expected {
		t.Errorf("Tuple isVector test error. Expected %v, got %v", expected, actual)
	}

	// Not a vector
	expected = false
	actual = Tuple{X: 1.0, Y: -2.5, Z: 3.0, W: 1.0}.IsVector()

	if actual != expected {
		t.Errorf("Tuple isVector test error. Expected %v, got %v", expected, actual)
	}
}

// Create point Test
func TestPoint(t *testing.T) {
	expected := Tuple{X: 3, Y: 8, Z: 9, W: 1}
	actual := Point(3, 8, 9);
	if actual != expected {
		t.Errorf("Point() func test error. Expected %v, got %v", expected, actual)
	}
}

// Create point Test
func TestVector(t *testing.T) {
	expected := Tuple{X: 3, Y: 8, Z: 9, W: 0}
	actual := Vector(3, 8, 9);
	if actual != expected {
		t.Errorf("Vector() func test error. Expected %v, got %v", expected, actual)
	}
}


