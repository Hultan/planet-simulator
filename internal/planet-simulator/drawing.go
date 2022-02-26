package planet_simulator

import (
	"image/color"
	"math"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/planet-simulator/internal/data"
)

var center = data.Vector2{}
var width float64
var height float64

// onDraw : The onDraw signal handler
func (p *PlanetSimulator) onDraw(da *gtk.DrawingArea, ctx *cairo.Context) {
	width = float64(da.GetAllocatedWidth())
	height = float64(da.GetAllocatedHeight())
	center = data.Vector2{X: width / 2, Y: height / 2}

	p.drawBackground(da, ctx)

	for i := range p.data.Bodies {
		body := p.data.Bodies[i]
		p.drawPlanet(da, ctx, body)
	}
}

//
// HELPER FUNCTIONS
//

// drawBackground : Draws the background
func (p *PlanetSimulator) drawBackground(_ *gtk.DrawingArea, ctx *cairo.Context) {
	p.setColor(ctx, color.Black)
	ctx.Rectangle(0, 0, width, height)
	ctx.Fill()
}

// drawPlanet : Draws a planet
func (p *PlanetSimulator) drawPlanet(da *gtk.DrawingArea, ctx *cairo.Context, planet *data.Body) {
	// Scale position and center it
	c := planet.Position.Mul(SCALE).Add(center)

	p.setColor(ctx, planet.ColorObj)
	p.drawPlanetOrbit(da, ctx, planet)
	ctx.Arc(c.X, c.Y, planet.Radius, 0, 2*math.Pi)
	ctx.Fill()
}

func (p *PlanetSimulator) drawPlanetOrbit(_ *gtk.DrawingArea, ctx *cairo.Context, planet *data.Body) {
	ctx.SetLineWidth(1)
	for i, point := range planet.Orbit {
		// Scale position and center it
		pos := point.Mul(SCALE).Add(center)

		if i == 0 {
			ctx.MoveTo(pos.X, pos.Y)
		} else {
			ctx.LineTo(pos.X, pos.Y)
		}
	}
	ctx.Stroke()
}

// setColor : Sets the draw color to and RGB color
func (p *PlanetSimulator) setColor(ctx *cairo.Context, c color.Color) {
	r, g, b, a := c.RGBA()
	ctx.SetSourceRGBA(col(r), col(g), col(b), col(a))
}
