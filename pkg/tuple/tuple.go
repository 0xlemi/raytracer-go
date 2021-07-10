package tuple

type Tuple struct {
	X, Y, Z, W float64
}

func NewTuple(x, y, z, w float64) Tuple {
	return Tuple{x, y, z, w}
}
