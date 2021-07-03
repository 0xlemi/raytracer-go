package matrix

import (
	"math"
)

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
