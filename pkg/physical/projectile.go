package physical

import (
	t "github.com/lemidev/raytracer/pkg/tuple"
)

type Projectile struct {
	Position t.Point
	Velocity t.Vector
}

func NewProjectile(p t.Point, v t.Vector) *Projectile {
	return &Projectile{Position: p, Velocity: v}
}

func (p *Projectile) Tick(env Environment) {
	p.Position = p.Position.AddVector(p.Velocity)
	p.Velocity = p.Velocity.Add(env.Gravity).Add(env.Wind)
}
