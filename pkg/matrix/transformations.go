package matrix

import (
	m "math"
)

func Translation(x, y, z float64) Matrix4 {
	return NewMatrix4([4][4]float64{
		{1, 0, 0, x},
		{0, 1, 0, y},
		{0, 0, 1, z},
		{0, 0, 0, 1},
	})
}

func Scaling(x, y, z float64) Matrix4 {
	return NewMatrix4([4][4]float64{
		{x, 0, 0, 0},
		{0, y, 0, 0},
		{0, 0, z, 0},
		{0, 0, 0, 1},
	})
}

func RotationX(r float64) Matrix4 {
	return NewMatrix4([4][4]float64{
		{1, 0, 0, 0},
		{0, m.Cos(r), -m.Sin(r), 0},
		{0, m.Sin(r), m.Cos(r), 0},
		{0, 0, 0, 1},
	})
}

func RotationY(r float64) Matrix4 {
	return NewMatrix4([4][4]float64{
		{m.Cos(r), 0, m.Sin(r), 0},
		{0, 1, 0, 0},
		{-m.Sin(r), 0, m.Cos(r), 0},
		{0, 0, 0, 1},
	})
}

func RotationZ(r float64) Matrix4 {
	return NewMatrix4([4][4]float64{
		{m.Cos(r), -m.Sin(r), 0, 0},
		{m.Sin(r), m.Cos(r), 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	})
}
