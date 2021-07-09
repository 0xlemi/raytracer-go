package matrix

type Matrix2 struct {
	Elements [2][2]float64
}

func NewMatrix2(elemParams ...[2][2]float64) Matrix2 {
	// Optional pramenter for instancing
	elements := [2][2]float64{}
	if len(elemParams) > 0 {
		elements = elemParams[0]
	}
	return Matrix2{Elements: elements}
}

func (m Matrix2) WriteElem(row, col uint8, value float64) Matrix2 {
	m.Elements[row][col] = value
	return m
}

func (m Matrix2) ReadElem(row, col uint8) float64 {
	return m.Elements[row][col]
}

// Determinant comes from solving system of equations using matrixes
// if det() is 0 then system of equations have no solution
func (m Matrix2) Determinant() float64 {
	a := m.ReadElem(0, 0)
	b := m.ReadElem(0, 1)
	c := m.ReadElem(1, 0)
	d := m.ReadElem(1, 1)
	return (a*d) - (b*c)
}