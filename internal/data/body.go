package data

import (
	"image/color"
	"math"
)

const G = 6.67428e-11

type Body struct {
	IsSun    bool
	Name     string
	Radius   float64
	Mass     float64
	Position Vector2
	Color    string
	ColorObj color.Color

	velocity      Vector2
	distanceToSun float64
	orbit         []Vector2
}

func (b *Body) CalculateAttraction(other *Body) Vector2 {
	d := other.Position.Sub(b.Position)
	distance := d.Length()

	if other.IsSun {
		b.distanceToSun = distance
	}

	f := G * b.Mass * other.Mass / distance / distance // F = m*M / r^2
	theta := math.Atan2(d.Y, d.X)
	return Vector2{X: math.Cos(theta), Y: math.Sin(theta)}.Mul(f)
}

func (b *Body) UpdatePosition(solar *SolarSystem, timestamp float64) {
	total := Vector2{}
	for i := range solar.Bodies {
		body := solar.Bodies[i]
		if b == body {
			continue
		}
		total = total.Add(b.CalculateAttraction(body))
	}
	b.velocity = b.velocity.Add(total.Div(b.Mass).Mul(timestamp)) // F = m*a  &  a = v*t
	b.Position = b.Position.Add(b.velocity.Mul(timestamp))
	b.orbit = append(b.orbit, b.Position)
}
