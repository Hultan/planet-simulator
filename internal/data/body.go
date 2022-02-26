package data

import (
	"image/color"
	"math"
)

// g : Gravity constant
const g = 6.67428e-11

// Body : A celestial body, star or planet
type Body struct {
	IsStar   bool        // IsStar : Is this the star of the system
	Name     string      // Name : The name of the star/planet
	Radius   float64     // Radius : The radius of the star/planet
	Mass     float64     // Mass : Tha mass of the star/planet
	Position Vector2     // Position : The position of the star/planet
	Velocity Vector2     // Velocity : The velocity of the star/planet
	Color    string      // Color : The color of the star/planet in HEX form
	ColorObj color.Color // ColorObj : The color of the star/planet in RGB form

	DistanceToStar float64   // DistanceToStar : The distance to the star of the system
	Orbit          []Vector2 // Orbit : The orbit of the planet
}

// Attraction : Calculates the attraction between two bodies
func (b *Body) Attraction(other *Body) Vector2 {
	// Calculate the vector to the other body
	d := other.Position.Sub(b.Position)
	// Calculate the distance to the other body
	distance := d.Length()

	if other.IsStar {
		b.DistanceToStar = distance
	}

	// Calculate the force on the body
	f := g * b.Mass * other.Mass / distance / distance // F = m*M / r^2
	// Calculate the angle of the force
	theta := math.Atan2(d.Y, d.X)
	// Return the force vector
	return Vector2{X: math.Cos(theta), Y: math.Sin(theta)}.Mul(f)
}

// UpdatePosition : Update the position each cycle
func (b *Body) UpdatePosition(solar *SolarSystem, timestamp float64) {
	total := Vector2{}
	// For each body, calculate the force to every other body
	for i := range solar.Bodies {
		body := solar.Bodies[i]
		if b == body {
			continue
		}
		total = total.Add(b.Attraction(body))
	}

	// Calculate the velocity change
	b.Velocity = b.Velocity.Add(total.Mul(timestamp / b.Mass)) // F = m*a  &  a = v*t
	// Calculate the position change
	b.Position = b.Position.Add(b.Velocity.Mul(timestamp))
	// Store the new position in the orbit
	b.Orbit = append(b.Orbit, b.Position)
}
