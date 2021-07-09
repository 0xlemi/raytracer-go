package matrix

import (
	"math"
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

// func (m Matrix4) GetRow(num uint8) [4]float64 {
// 	return m.Elements[num]
// }

func (m Matrix4) GetColumn(num uint8) [4]float64 {
	res := [4]float64{}
	for i := 0; i < 4; i++ {
		res[i] = m.Elements[i][num]
	}
	return res
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
// func ()  {

// }

func (m Matrix4) Transpose() Matrix4 {
	matrix := NewMatrix4()
	matrix.Elements[0] = m.GetColumn(0)
	matrix.Elements[1] = m.GetColumn(1)
	matrix.Elements[2] = m.GetColumn(2)
	matrix.Elements[3] = m.GetColumn(3)
	return matrix
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
