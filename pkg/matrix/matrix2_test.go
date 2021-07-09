package matrix

import (
	"fmt"
	"testing"
)

func TestMatrix2Construct(t *testing.T) {
	expected := "[[0 0] [0 0]]"
	actual := fmt.Sprint(NewMatrix2().Elements)

	if actual != expected {
		t.Errorf("TestMatrix2Construct test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrix2ConstructWithValues(t *testing.T) {
	expected := "[[1 2] [5.5 6.5]]"
	actual := fmt.Sprint(NewMatrix2([2][2]float64{
		{1, 2},
		{5.5, 6.5},
	}).Elements)

	if actual != expected {
		t.Errorf("TestMatrix2ConstructWithValues test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrix2WriteElem(t *testing.T) {
	expected := "[[1 0] [5.5 7.5]]"
	m := NewMatrix2()
	m = m.WriteElem(0, 0, 1)
	m = m.WriteElem(1, 0, 5.5)
	m = m.WriteElem(1, 1, 7.5)
	actual := fmt.Sprint(m.Elements)

	if actual != expected {
		t.Errorf("TestMatrix2WriteElem test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrix2ReadElem(t *testing.T) {
	expected := [4]float64{1, 4, 5.5, 7.5}
	m := NewMatrix2([2][2]float64{
		{1, 4},
		{5.5, 7.5},
	})
	actual := [4]float64{}
	actual[0] = m.ReadElem(0, 0)
	actual[1] = m.ReadElem(0, 1)
	actual[2] = m.ReadElem(1, 0)
	actual[3] = m.ReadElem(1, 1)

	if actual != expected {
		t.Errorf("TestMatrix2ReadElem test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrix2Determinant(t *testing.T) {
	m := NewMatrix2([2][2]float64{
		{1, 5},
		{-3, 2},
	})
	expected := 17.0
	actual := m.Determinant()

	if actual != expected {
		t.Errorf("TestMatrix2Determinant test error. Expected %v, got %v", expected, actual)
	}
}
