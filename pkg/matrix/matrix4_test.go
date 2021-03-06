package matrix

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/lemidev/raytracer/pkg/tuple"
)

func TestMatrix4Construct(t *testing.T) {
	expected := "[[0 0 0 0] [0 0 0 0] [0 0 0 0] [0 0 0 0]]"
	actual := fmt.Sprint(NewMatrix4().Elements)

	if actual != expected {
		t.Errorf("TestMatrix4Construct test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrix4ConstructWithValues(t *testing.T) {
	expected := "[[1 2 3 4] [5.5 6.5 7.5 8.5] [9 10 11 12] [13.5 14.5 15.5 16.5]]"
	actual := fmt.Sprint(NewMatrix4([4][4]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	}).Elements)

	if actual != expected {
		t.Errorf("TestMatrix4ConstructWithValues test error. Expected %v, got %v", expected, actual)
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

func TestMatrix4GetColumn(t *testing.T) {
	expected := [4]float64{2, 6.5, 10, 14.5}
	m := NewMatrix4([4][4]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	})
	actual := m.GetColumn(1)

	if actual != expected {
		t.Errorf("TestMatrix4GetColumn test error. Expected %v, got %v", expected, actual)
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

func TestMatrix4LinearlyScale(t *testing.T) {
	expected := NewMatrix4([4][4]float64{
		{2, 4, 6, 8},
		{11, 13, 15, 17},
		{18, 20, 22, 24},
		{27, 29, 31, 33},
	})
	m := NewMatrix4([4][4]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	})
	actual := m.LinearlyScale(2)

	if actual != expected {
		t.Errorf("TestMatrix4Scale test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrixMulti(t *testing.T) {
	a := NewMatrix4([4][4]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	})
	b := NewMatrix4([4][4]float64{
		{-2, 1, 2, 3},
		{3, 2, 1, -1},
		{4, 3, 6, 5},
		{1, 2, 7, 8},
	})
	res := NewMatrix4([4][4]float64{
		{20, 22, 50, 48},
		{44, 54, 114, 108},
		{40, 58, 110, 102},
		{16, 26, 46, 42},
	})
	expected := res
	actual := a.Multi(b)

	if actual != expected {
		t.Errorf("TestMatrixMulti test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrix4MultiTuple(t *testing.T) {
	expected := tuple.NewTuple(18, 24, 33, 1)
	m := NewMatrix4([4][4]float64{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1},
	})
	actual := m.MultiTuple(tuple.NewTuple(1, 2, 3, 1))

	if actual != expected {
		t.Errorf("TestMatrix4MultiTuple test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrix4MultiTupleIdentity(t *testing.T) {
	expected := tuple.NewTuple(18, 24, 33, 1)
	actual := IDENTITY_MATRIX4.MultiTuple(tuple.NewTuple(18, 24, 33, 1))

	if actual != expected {
		t.Errorf("TestMatrix4MultiTupleIdentity test error. Expected %v, got %v", expected, actual)
	}
}

func TestMultiIdentityMatrix(t *testing.T) {
	a := NewMatrix4([4][4]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	})
	expected := a
	actual := a.Multi(IDENTITY_MATRIX4)

	if actual != expected {
		t.Errorf("TestMultiIdentityMatrix test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrix4Transposing(t *testing.T) {
	a := NewMatrix4([4][4]float64{
		{0, 9, 3, 0},
		{9, 8, 0, 8},
		{1, 8, 5, 3},
		{0, 0, 5, 8},
	})
	b := NewMatrix4([4][4]float64{
		{0, 9, 1, 0},
		{9, 8, 8, 0},
		{3, 0, 5, 5},
		{0, 8, 3, 8},
	})
	expected := b
	actual := a.Transpose()

	if actual != expected {
		t.Errorf("TestMatrix4Transposing test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrix4TransposingIdentityMatrix(t *testing.T) {
	expected := IDENTITY_MATRIX4
	actual := IDENTITY_MATRIX4.Transpose()

	if actual != expected {
		t.Errorf("TestMatrix4TransposingIdentityMatrix test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrix4Submatrix(t *testing.T) {
	expected := NewMatrix3([3][3]float64{
		{-6, 1, 6},
		{-8, 8, 6},
		{-7, -1, 1},
	})
	m := NewMatrix4([4][4]float64{
		{-6, 1, 1, 6},
		{-8, 5, 8, 6},
		{-1, 0, 8, 2},
		{-7, 1, -1, 1},
	})
	actual, _ := m.Submatrix(2, 1)

	if actual != expected {
		t.Errorf("TestMatrix4Submatrix test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrix4Submatrix2(t *testing.T) {
	expected := NewMatrix3([3][3]float64{
		{-6, 1, 6},
		{-8, 5, 6},
		{-1, 0, 2},
	})
	m := NewMatrix4([4][4]float64{
		{-6, 1, 1, 6},
		{-8, 5, 8, 6},
		{-1, 0, 8, 2},
		{-7, 1, -1, 1},
	})
	actual, _ := m.Submatrix(3, 2)

	if actual != expected {
		t.Errorf("TestMatrix4Submatrix2 test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrix4SubmatrixError(t *testing.T) {
	expected := "*errors.errorString"
	m := NewMatrix4([4][4]float64{
		{-6, 1, 1, 6},
		{-8, 5, 8, 6},
		{-1, 0, 8, 2},
		{-7, 1, -1, 1},
	})
	_, err := m.Submatrix(4, 2)
	actual := reflect.TypeOf(err).String()

	if actual != expected {
		t.Errorf("TestMatrixpSubmatrixError test error. Expected %v, got %v", expected, actual)
	}
}

// No need for minor test. This also test minor
func TestMatrix4Cofactor(t *testing.T) {
	expected := [4]float64{690, 447, 210, 51}
	m := NewMatrix4([4][4]float64{
		{-2, -8, 3, 5},
		{-3, 1, 7, 3},
		{1, 2, -9, 6},
		{-6, 7, 7, -9},
	})
	co1, _ := m.Cofactor(0, 0)
	co2, _ := m.Cofactor(0, 1)
	co3, _ := m.Cofactor(0, 2)
	co4, _ := m.Cofactor(0, 3)
	actual := [4]float64{
		co1, co2, co3, co4,
	}

	if actual != expected {
		t.Errorf("TestMatrix4Cofactor test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrix4Determinant(t *testing.T) {
	expected := -4071.0
	m := NewMatrix4([4][4]float64{
		{-2, -8, 3, 5},
		{-3, 1, 7, 3},
		{1, 2, -9, 6},
		{-6, 7, 7, -9},
	})
	actual, _ := m.Determinant()

	if actual != expected {
		t.Errorf("TestMatrix4Determinant test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrix4IsNotInvertable(t *testing.T) {
	expected := 0.0
	m := NewMatrix4([4][4]float64{
		{-4, 2, -2, -3},
		{9, 6, 2, 6},
		{0, -5, 1, -5},
		{0, 0, 0, 0},
	})
	actual, _ := m.Determinant()

	if actual != expected {
		t.Errorf("TestMatrix4IsInvertable test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrix4Inverse(t *testing.T) {
	expected := NewMatrix4([4][4]float64{
		{0.21805, 0.45113, 0.24060, -0.04511},
		{-0.80827, -1.45677, -0.44361, 0.52068},
		{-0.07895, -0.22368, -0.05263, 0.19737},
		{-0.52256, -0.81391, -0.30075, 0.30639},
	})
	m := NewMatrix4([4][4]float64{
		{-5, 2, 6, -8},
		{1, -5, 1, 8},
		{7, 7, -6, -7},
		{1, -3, 7, 4},
	})
	actual, _ := m.Inverse()

	if !actual.Equals(expected) {
		t.Errorf("TestMatrix4Inverse test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrix4InverseOfNotInvertible(t *testing.T) {
	expected := "*errors.errorString"
	m := NewMatrix4([4][4]float64{
		{-4, 2, -2, -3},
		{9, 6, 2, 6},
		{0, -5, 1, -5},
		{0, 0, 0, 0},
	})
	_, err := m.Inverse()
	actual := reflect.TypeOf(err).String()

	if actual != expected {
		t.Errorf("TestMatrix4InverseOfNotInvertible test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrix4MultiplyingByInverse(t *testing.T) {
	a := NewMatrix4([4][4]float64{
		{3, -9, 7, 3},
		{3, -8, 2, -9},
		{-4, 4, 4, 1},
		{-6, 5, -1, 1},
	})
	b := NewMatrix4([4][4]float64{
		{8, 2, 2, 2},
		{3, -1, 7, 0},
		{7, 0, 5, 4},
		{6, -2, 0, 5},
	})
	expected := a

	invB, _ := b.Inverse()
	actual := a.Multi(b).Multi(invB)

	if !actual.Equals(expected) {
		t.Errorf("TestMatrix4MultiplyingByInverse test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrix4InverseIdentityMatrix(t *testing.T) {
	expected := IDENTITY_MATRIX4
	actual, _ := IDENTITY_MATRIX4.Inverse()

	if actual != expected {
		t.Errorf("TestMatrix4InverseIdentityMatrix test error. Expected %v, got %v", expected, actual)
	}
}

func TestMatrix4InverseTransposeEqualsTransposeInverse(t *testing.T) {
	a := NewMatrix4([4][4]float64{
		{3, -9, 7, 3},
		{3, -8, 2, -9},
		{-4, 4, 4, 1},
		{-6, 5, -1, 1},
	})
	aInv, _ := a.Inverse()

	expected := aInv.Transpose()
	actual, _ := a.Transpose().Inverse()

	if actual != expected {
		t.Errorf("TestMatrix4InverseTransposeEqualsTransposeInverse test error. Expected %v, got %v", expected, actual)
	}
}
