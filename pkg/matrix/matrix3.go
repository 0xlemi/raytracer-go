package matrix

type Matrix3 struct {
	Elements [3][3]float64
}

func NewMatrix3() Matrix3 {
	elements := [3][3]float64{}
	return Matrix3{Elements: elements}
}

func (m Matrix3) WriteElem(row, col uint8, value float64) Matrix3 {
	m.Elements[row][col] = value
	return m
}

func (m Matrix3) ReadElem(row, col uint8) float64 {
	return m.Elements[row][col]
}
