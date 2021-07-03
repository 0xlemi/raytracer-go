package matrix

import (
	"fmt"
	"testing"
)

func TestMatrix4Construct(t *testing.T) {
	expected := "[[0 0 0 0] [0 0 0 0] [0 0 0 0] [0 0 0 0]]"
	actual := fmt.Sprint(NewMatrix4().Elements)

	if actual != expected {
		t.Errorf("TestMatrix4Construct test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrix4WriteElem(t *testing.T) {
	expected := "[[1 0 0 4] [5.5 0 7.5 0] [0 0 11 0] [13.5 0 15.5 0]]"
	m := NewMatrix4()
	m = m.WriteElem(0, 0, 1)
	m = m.WriteElem(0, 3, 4)
	m = m.WriteElem(1, 0, 5.5)
	m = m.WriteElem(1, 2, 7.5)
	m = m.WriteElem(2, 2, 11)
	m = m.WriteElem(3, 0, 13.5)
	m = m.WriteElem(3, 2, 15.5)
	actual := fmt.Sprint(m.Elements)

	if actual != expected {
		t.Errorf("TestMatrix4WriteElem test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrix4ReadElem(t *testing.T) {
	expected := [7]float64{1, 4, 5.5, 7.5, 11, 13.5, 15.5}
	elements := [4][4]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	}
	m := NewMatrix4(elements)
	actual := [7]float64{}
	actual[0] = m.ReadElem(0, 0)
	actual[1] = m.ReadElem(0, 3)
	actual[2] = m.ReadElem(1, 0)
	actual[3] = m.ReadElem(1, 2)
	actual[4] = m.ReadElem(2, 2)
	actual[5] = m.ReadElem(3, 0)
	actual[6] = m.ReadElem(3, 2)

	if actual != expected {
		t.Errorf("TestMatrix4ReadElem test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrix4Equals(t *testing.T) {

	m1 := [4][4]float64{
		{1.000002, 2, 3.000003, 4},
		{5.5, 6.5, 7.500001, 8.5},
		{9.000003, 10, 11.000004, 12},
		{13.5, 14.5000021, 15.5000045, 16.500001},
	}
	m2 := [4][4]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	}
	matrix1 := NewMatrix4(m1)
	matrix2 := NewMatrix4(m2)

	expected := true
	actual := matrix1.Equals(matrix2)

	if actual != expected {
		t.Errorf("TestMatrix4Equals test error. Expected %v, got %v", expected, actual)
	}
}
