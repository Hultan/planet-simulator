package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVector2_Add(t *testing.T) {
	v := Vector2{X: 5, Y: 7}
	v2 := Vector2{X: 3, Y: 1}
	got := v.Add(v2)
	assert.Equal(t, 8.0, got.X)
	assert.Equal(t, 8.0, got.Y)
}

func TestVector2_Div(t *testing.T) {
	v := Vector2{X: 5, Y: 7}
	got := v.Div(2.0)
	assert.Equal(t, 2.5, got.X)
	assert.Equal(t, 3.5, got.Y)
}

func TestVector2_Mul(t *testing.T) {
	v := Vector2{X: 5, Y: 7}
	got := v.Mul(3.0)
	assert.Equal(t, 15.0, got.X)
	assert.Equal(t, 21.0, got.Y)
}

func TestVector2_Sub(t *testing.T) {
	v := Vector2{X: 5, Y: 7}
	v2 := Vector2{X: 3, Y: 1}
	got := v.Sub(v2)
	assert.Equal(t, 2.0, got.X)
	assert.Equal(t, 6.0, got.Y)
}

func TestVector2_Advanced(t *testing.T) {
	v := Vector2{X: 5, Y: 7}
	v2 := Vector2{X: 3, Y: 1}
	got := v.Add(v2).Mul(2).Div(2).Sub(v2)

	assert.Equal(t, 5.0, got.X)
	assert.Equal(t, 7.0, got.Y)
}

func TestVector2_Distance(t *testing.T) {
	v := Vector2{X: 4, Y: 0}
	v2 := Vector2{X: 0, Y: 3}
	d := v.Sub(v2)
	got := d.Length()

	assert.Equal(t, 5.0, got)
}
