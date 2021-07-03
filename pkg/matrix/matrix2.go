package matrix

type Matrix2 struct {
	Elements [2][2]float64
}

func NewMatrix2() Matrix2 {
	elements := [2][2]float64{}
	return Matrix2{Elements: elements}
}

func (m Matrix2) WriteElem(row, col uint8, value float64) Matrix2 {
	m.Elements[row][col] = value
	return m
}

func (m Matrix2) ReadElem(row, col uint8) float64 {
	return m.Elements[row][col]
}
