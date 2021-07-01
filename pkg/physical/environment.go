package physical

import (
	t "github.com/lemidev/raytracer/pkg/tuple"
)

type Environment struct {
	Gravity t.Vector
	Wind    t.Vector
}

func NewEnvironment(g t.Vector, w t.Vector) Environment {
	return Environment{Gravity: g, Wind: w}
}
