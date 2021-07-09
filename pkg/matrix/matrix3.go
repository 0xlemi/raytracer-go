package matrix

import (
	"errors"
)

type Matrix3 struct {
	Elements [3][3]float64
}

func NewMatrix3(elemParams ...[3][3]float64) Matrix3 {
	// Optional pramenter for instancing
	elements := [3][3]float64{}
	if len(elemParams) > 0 {
		elements = elemParams[0]
	}
	return Matrix3{Elements: elements}
}

func (m Matrix3) WriteElem(row, col uint8, value float64) Matrix3 {
	m.Elements[row][col] = value
	return m
}

func (m Matrix3) ReadElem(row, col uint8) float64 {
	return m.Elements[row][col]
}

func (m Matrix3) Submatrix(delRow, delCol uint8) (Matrix2, error) {
	// Check that delRow and delCol are not greater than 2 becouse
	// Matrix3 only is 3x3
	if (delRow > 2) || (delCol > 2) {
		return NewMatrix2(), errors.New("Matrix3 Submatrix Method does not accept rows or cols greater than 2")
	}

	subM := [2][2]float64{}
	subRow := uint8(0)
	subCol := uint8(0)
	for row := uint8(0); row < 3; row++ {
		if row == delRow {
			continue
		}
		for col := uint8(0); col < 3; col++ {
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

	return NewMatrix2(subM), nil
}
