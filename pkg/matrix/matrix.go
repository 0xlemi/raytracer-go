package matrix

type Matrix interface {
	WriteElem()
	ReadElem()
}

const EPSILON float64 = 0.00001

// // Generalized function
// func (m1 Matrix) Equals(m2 Matrix) bool {
// 	result := true
// 	for row, rows := range m1.Elements {
// 		for column := range rows {
// 			m1Value := m1.ReadElem(uint8(row), uint8(column))
// 			m2Value := m2.ReadElem(uint8(row), uint8(column))
// 			difference := m1Value - m2Value
// 			result = result && (math.Abs(difference) < EPSILON)
// 		}
// 	}
// 	return result
// }