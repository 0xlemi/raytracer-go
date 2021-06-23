package set

import (
	"testing"
)

func TestCreateColor(t *testing.T) {
	expected := Color{Red: 1, Green: .78, Blue: .42}
	actual := NewColor(1, .78, .42)

	if actual != expected {
		t.Errorf("TestCreateColor test error. Expected %v, got %v", expected, actual)
	}
}

func TestColorAdd(t *testing.T) {
	expected := NewColor(1.6, 0.7, 1.0)
	actual := NewColor(0.9, 0.6, 0.75).Add(NewColor(0.7, 0.1, 0.25))

	if actual != expected {
		t.Errorf("TestColorAdd test error. Expected %v, got %v", expected, actual)
	}
}

func TestColorSubtracting(t *testing.T) {
	expected := NewColor(0.2, 0.5, 0.5)
	actual := NewColor(0.9, 0.6, 0.75).Sub(NewColor(0.7, 0.1, 0.25))

	if !ColorEquals(actual, expected){
		t.Errorf("TestColorSubtracting test error. Expected %v, got %v", expected, actual)
	}
}

func TestColorMultiply(t *testing.T) {
	expected := NewColor(0.2, 0.3, 0.4).Multi(2)
	actual := NewColor(0.4, 0.6, 0.8)

	if !ColorEquals(actual, expected){
		t.Errorf("TestColorMultiply test error. Expected %v, got %v", expected, actual)
	}
}

func TestColorHadamard(t *testing.T) {
	expected := NewColor(0.9, 0.2, 0.04)
	actual := NewColor(1, 0.2, 0.4).Hadamard(NewColor(0.9, 1, 0.1))

	if !ColorEquals(actual, expected) {
		t.Errorf("TestColorHadamard test error. Expected %v, got %v", expected, actual)
	}
}

