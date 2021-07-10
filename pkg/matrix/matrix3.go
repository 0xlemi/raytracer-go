package matrix

import (
	"errors"
	"fmt"
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

func (m Matrix3) Minor(row, col uint8) (float64, error) {
	subM, err := m.Submatrix(row, col)
	if err != nil {
		return -1, fmt.Errorf("error Matrix3.Minor: %w", err)
	}
	return subM.Determinant(), nil
}

func (m Matrix3) Cofactor(row, col uint8) (float64, error) {
	minor, err := m.Minor(row, col)
	if err != nil {
		return -1, fmt.Errorf("error Matrix3.Cofactor: %w", err)
	}
	// check if row + col is even
	isEven := ((row+col)%2 == 0)

	if isEven {
		return minor, nil
	}
	return (minor * -1), nil
}

func (m Matrix3) Determinant() (float64, error) {
	row := m.Elements[0]
	res := 0.0
	for i, num := range row {
		co, err := m.Cofactor(0, uint8(i))
		if err != nil {
			return -1, fmt.Errorf("error Matrix3.Determinant: %w", err)
		}
		res += co * num
	}
	return res, nil
}
