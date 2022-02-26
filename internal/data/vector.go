package data

import (
	"math"
)

// Vector2 : A 2-dimensional vector
type Vector2 struct {
	X, Y float64
}

// Add : Adds two vectors
func (v Vector2) Add(o Vector2) Vector2 {
	return Vector2{v.X + o.X, v.Y + o.Y}
}

// Sub : Subtracts to vectors
func (v Vector2) Sub(o Vector2) Vector2 {
	return Vector2{v.X - o.X, v.Y - o.Y}
}

// Mul : Multiplies a vector with a factor
func (v Vector2) Mul(m float64) Vector2 {
	return Vector2{v.X * m, v.Y * m}
}

// Div : Divides a vector with a divisor
func (v Vector2) Div(d float64) Vector2 {
	return Vector2{v.X / d, v.Y / d}
}

// Length : Returns the length of a vector
func (v Vector2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
