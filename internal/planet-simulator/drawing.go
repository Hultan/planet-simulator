package planet_simulator

import (
	"image/color"
	"math"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/planet-simulator/internal/data"
)

// onDraw : The onDraw signal handler
func (p *PlanetSimulator) onDraw(da *gtk.DrawingArea, ctx *cairo.Context) {
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
func (p *PlanetSimulator) drawBackground(da *gtk.DrawingArea, ctx *cairo.Context) {
	width := float64(da.GetAllocatedWidth())
	height := float64(da.GetAllocatedHeight())
	p.setColor(ctx, color.Black)
	ctx.Rectangle(0, 0, width, height)
	ctx.Fill()
}

// drawPlanet : Draws a planet
func (p *PlanetSimulator) drawPlanet(da *gtk.DrawingArea, ctx *cairo.Context, planet *data.Body) {
	width := float64(da.GetAllocatedWidth())
	height := float64(da.GetAllocatedHeight())
	x := planet.Position.X*AU*SCALE + width/2 // TODO : Fix center coords, use translation?
	y := planet.Position.Y*AU*SCALE + height/2
	p.setColor(ctx, planet.ColorObj)
	ctx.Arc(x, y, planet.Radius, 0, 2*math.Pi)
	ctx.Fill()
}

func (p *PlanetSimulator) setColor(ctx *cairo.Context, c color.Color) {
	r, g, b, a := c.RGBA()
	ctx.SetSourceRGBA(col(r), col(g), col(b), col(a))
}
