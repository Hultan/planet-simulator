package data

// SolarSystem : A system containing a star and one or more planets
type SolarSystem struct {
	CenterX, CenterY float64
	Bodies           []*Body
}
