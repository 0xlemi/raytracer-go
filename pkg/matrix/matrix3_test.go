package matrix

import (
	"reflect"
	"fmt"
	"testing"
)

func TestMatrix3Construct(t *testing.T) {
	expected := "[[0 0 0] [0 0 0] [0 0 0]]"
	actual := fmt.Sprint(NewMatrix3().Elements)

	if actual != expected {
		t.Errorf("TestMatrix3Construct test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrix3ConstructWithValues(t *testing.T) {
	expected := "[[1 2 3] [5.5 6.5 7.5] [8 9 10]]"
	actual := fmt.Sprint(NewMatrix3([3][3]float64{
		{1, 2, 3},
		{5.5, 6.5, 7.5},
		{8, 9, 10},
	}).Elements)

	if actual != expected {
		t.Errorf("TestMatrix3ConstructWithValues test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrix3WriteElem(t *testing.T) {
	expected := "[[1 2 0] [5.5 6.5 0] [8 0 10]]"
	m := NewMatrix3()
	m = m.WriteElem(0, 0, 1)
	m = m.WriteElem(0, 1, 2)
	m = m.WriteElem(1, 0, 5.5)
	m = m.WriteElem(1, 1, 6.5)
	m = m.WriteElem(2, 0, 8)
	m = m.WriteElem(2, 2, 10)
	actual := fmt.Sprint(m.Elements)

	if actual != expected {
		t.Errorf("TestMatrix3WriteElem test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrix3ReadElem(t *testing.T) {
	expected := [4]float64{1, 3, 5.5, 10}
	m := NewMatrix3([3][3]float64{
		{1, 2, 3},
		{5.5, 6.5, 7.5},
		{8, 9, 10},
	})
	actual := [4]float64{}
	actual[0] = m.ReadElem(0, 0)
	actual[1] = m.ReadElem(0, 2)
	actual[2] = m.ReadElem(1, 0)
	actual[3] = m.ReadElem(2, 2)

	if actual != expected {
		t.Errorf("TestMatrix3ReadElem test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrix3Submatrix(t *testing.T) {
	expected := NewMatrix2([2][2]float64{
		{1, 2},
		{5.5, 6.5},
	})
	m := NewMatrix3([3][3]float64{
		{1, 2, 3},
		{5.5, 6.5, 7.5},
		{8, 9, 10},
	})
	actual, _ := m.Submatrix(2, 2)

	if actual != expected {
		t.Errorf("TestMatrix3Submatrix test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrix3Submatrix2(t *testing.T) {
	expected := NewMatrix2([2][2]float64{
		{5.5, 7.5},
		{8, 10},
	})
	m := NewMatrix3([3][3]float64{
		{1, 2, 3},
		{5.5, 6.5, 7.5},
		{8, 9, 10},
	})
	actual, _ := m.Submatrix(0, 1)

	if actual != expected {
		t.Errorf("TestMatrix3Submatrix2 test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrix3SubmatrixError(t *testing.T) {
	expected := "*errors.errorString"
	m := NewMatrix3([3][3]float64{
		{1, 2, 3},
		{5.5, 6.5, 7.5},
		{8, 9, 10},
	})
	_, err := m.Submatrix(3, 1)
	actual := reflect.TypeOf(err).String() 

	if actual != expected {
		t.Errorf("TestMatrix3Submatrix2 test error. Expected %v, got %v", expected, actual)
	}
}
