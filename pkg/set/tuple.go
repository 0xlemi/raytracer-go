package set

type Tuple struct {
	X, Y, Z, W float64
}

func (t Tuple) IsPoint() bool {
	return t.W == 1.0
}

func (t Tuple) IsVector() bool {
	return t.W == 0.0
}

func Point(x float64, y float64, z float64) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: 1.0}
}

func Vector(x float64, y float64, z float64) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: 0.0}
}
