package matrix

import (
	"errors"
	"fmt"
	"math"

	t "github.com/lemidev/raytracer/pkg/tuple"
)

var IDENTITY_MATRIX4 Matrix4 = NewMatrix4([4][4]float64{
	{1, 0, 0, 0},
	{0, 1, 0, 0},
	{0, 0, 1, 0},
	{0, 0, 0, 1},
})

type Matrix4 struct {
	Elements [4][4]float64
}

func NewMatrix4(elemParams ...[4][4]float64) Matrix4 {
	// Optional pramenter for instancing
	elements := [4][4]float64{}
	if len(elemParams) > 0 {
		elements = elemParams[0]
	}
	return Matrix4{Elements: elements}
}

func (m Matrix4) WriteElem(row, col uint8, value float64) Matrix4 {
	m.Elements[row][col] = value
	return m
}

func (m Matrix4) ReadElem(row, col uint8) float64 {
	return m.Elements[row][col]
}

func (m Matrix4) GetColumn(num uint8) [4]float64 {
	res := [4]float64{}
	for i := 0; i < 4; i++ {
		res[i] = m.Elements[i][num]
	}
	return res
}

func (m1 Matrix4) Equals(m2 Matrix4) bool {
	result := true
	for row := 0; row < 4; row++ {
		for column := 0; column < 4; column++ {
			m1Value := m1.ReadElem(uint8(row), uint8(column))
			m2Value := m2.ReadElem(uint8(row), uint8(column))
			difference := m1Value - m2Value
			result = result && (math.Abs(difference) < EPSILON)
		}
	}
	return result
}

func (m Matrix4) LinearlyScale(num float64) Matrix4 {
	new := NewMatrix4()
	for row := uint8(0); row < 4; row++ {
		for col := uint8(0); col < 4; col++ {
			curVal := m.ReadElem(row, col)
			new = new.WriteElem(row, col, curVal*num)
		}
	}
	return new
}

// Multily by another matrix
func (m1 Matrix4) Multi(m2 Matrix4) Matrix4 {
	matrix := NewMatrix4()
	for row := uint8(0); row < 4; row++ {
		for column := uint8(0); column < 4; column++ {
			result :=
				m1.ReadElem(row, 0)*m2.ReadElem(0, column) +
					m1.ReadElem(row, 1)*m2.ReadElem(1, column) +
					m1.ReadElem(row, 2)*m2.ReadElem(2, column) +
					m1.ReadElem(row, 3)*m2.ReadElem(3, column)
			matrix = matrix.WriteElem(row, column, result)
		}
	}
	return matrix
}

// Multilpy By Tuple function
func (m Matrix4) MultiTuple(tuple t.Tuple) t.Tuple {
	oldTuple := tuple.ToArray()
	new := t.NewTuple(0, 0, 0, 0)
	for row := uint8(0); row < 4; row++ {
		result :=
			m.ReadElem(row, 0)*oldTuple[0] +
				m.ReadElem(row, 1)*oldTuple[1] +
				m.ReadElem(row, 2)*oldTuple[2] +
				m.ReadElem(row, 3)*oldTuple[3]
		new = new.WriteElem(row, result) 
	}
	return new
}

func (m Matrix4) MultiTL(tl t.TupleLike) t.TupleLike {
	t := m.MultiTuple(tl.ToTuple())
	return t.ToTupleLike()
}

func (m Matrix4) Transpose() Matrix4 {
	matrix := NewMatrix4()
	matrix.Elements[0] = m.GetColumn(0)
	matrix.Elements[1] = m.GetColumn(1)
	matrix.Elements[2] = m.GetColumn(2)
	matrix.Elements[3] = m.GetColumn(3)
	return matrix
}

func (m Matrix4) Submatrix(delRow, delCol uint8) (Matrix3, error) {
	// Check that delRow and delCol are not greater than 2 becouse
	// Matrix3 only is 3x3
	if (delRow > 3) || (delCol > 3) {
		return NewMatrix3(), errors.New("Matrix4.Submatrix does not accept rows or cols greater than 2")
	}

	subM := [3][3]float64{}
	subRow := uint8(0)
	subCol := uint8(0)
	for row := uint8(0); row < 4; row++ {
		if row == delRow {
			continue
		}
		for col := uint8(0); col < 4; col++ {
			if col == delCol {
				continue
			}
			// Add Values to new [2][2] Array
			subM[subRow][subCol] = m.ReadElem(row, col)
			subCol++
		}
		// Reset the Col because is a 2d Array
		subCol = 0
		subRow++
	}

	return NewMatrix3(subM), nil
}

func (m Matrix4) Minor(row, col uint8) (float64, error) {

	subM, err := m.Submatrix(row, col)
	if err != nil {
		return -1, fmt.Errorf("error Matrix4.Minor: %w", err)
	}

	det, err := subM.Determinant()
	if err != nil {
		return -1, fmt.Errorf("error Matrix4.Minor: %w", err)
	}

	return det, nil
}

func (m Matrix4) Cofactor(row, col uint8) (float64, error) {
	minor, err := m.Minor(row, col)
	if err != nil {
		return -1, fmt.Errorf("error Matrix4.Cofactor: %w", err)
	}

	isEven := ((row+col)%2 == 0)

	if isEven {
		return minor, nil
	}
	return (minor * -1), nil
}

func (m Matrix4) Determinant() (float64, error) {
	row := m.Elements[0]
	res := 0.0
	for i, num := range row {
		co, err := m.Cofactor(0, uint8(i))
		if err != nil {
			return -1, fmt.Errorf("error Matrix4.Determinant: %w", err)
		}
		res += co * num
	}
	return res, nil
}

func (m Matrix4) Inverse() (Matrix4, error) {
	det, err := m.Determinant()
	if err != nil {
		return NewMatrix4(), fmt.Errorf("error Matrix4.Inverse: %w", err)
	}

	if det == 0.0 {
		return NewMatrix4(), errors.New("Matrix4.Inverse matrix not inversable because determinant = 0")
	}

	new := NewMatrix4()
	for row := uint8(0); row < 4; row++ {
		for col := uint8(0); col < 4; col++ {

			co, err := m.Cofactor(row, col)
			if err != nil {
				return NewMatrix4(), fmt.Errorf("error Matrix4.Inverse: %w", err)
			}

			// We need 3 Oparations
			// 1) Relpace values with cofactors
			// 2) Transpose the inverted array. Thats why (col, row) instead of (row, col)
			// 3) Divade all the elements with the original determinant
			new = new.WriteElem(col, row, co/det)
		}
	}
	return new, nil
}

func (m Matrix4) Translate(x, y, z float64) Matrix4 {
	return Translation(x, y, z).Multi(m)
}

func (m Matrix4) Scaling(x, y, z float64) Matrix4 {
	return Scaling(x, y, z).Multi(m)
}

func (m Matrix4) RotateX(r float64) Matrix4 {
	return RotationX(r).Multi(m)
}

func (m Matrix4) RotateY(r float64) Matrix4 {
	return RotationY(r).Multi(m)
}

func (m Matrix4) RotateZ(r float64) Matrix4 {
	return RotationZ(r).Multi(m)
}
