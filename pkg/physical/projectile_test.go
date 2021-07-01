package physical

import (
	"testing"

	"github.com/lemidev/raytracer/pkg/tuple"
)

func TestProjectileParabola(t *testing.T) {
	start := tuple.NewPoint(0, 1, 0)
	velocity := tuple.NewVector(1, 1.8, 0).Normalize()

	proj := NewProjectile(start, velocity)

	gravity := tuple.NewVector(0, -0.1, 0)
	wind := tuple.NewVector(-0.01, 0, 0)

	env := NewEnvironment(gravity, wind)

	parabola := []tuple.Point{}
	for 0 < proj.Position.Y {
		parabola = append(parabola, proj.Position)
		proj.Tick(env)
	}

	expected := tuple.NewPoint(4.4064293117863205, 5.241572761215378, 0)
	actual := parabola[10]

	if actual != expected {
		t.Errorf("TestProjectileParabola test error. Expected %v, got %v", expected, actual)
	}
}
