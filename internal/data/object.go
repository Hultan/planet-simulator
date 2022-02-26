package data

import (
	"image/color"
)

type Object struct {
	Name     string
	Mass     float64
	X, Y     float64
	Color    string
	ColorObj color.Color
}
