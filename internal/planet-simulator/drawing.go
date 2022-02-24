package planet_simulator

import (
	"image/color"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
)

// onDraw : The onDraw signal handler
func (p *PlanetSimulator) onDraw(da *gtk.DrawingArea, ctx *cairo.Context) {
	p.drawBackground(da, ctx)
}

//
// HELPER FUNCTIONS
//

// drawBackground : Draws the background
func (p *PlanetSimulator) drawBackground(da *gtk.DrawingArea, ctx *cairo.Context) {
	width := float64(da.GetAllocatedWidth())
	height := float64(da.GetAllocatedHeight())
	p.setColor(ctx, color.White)
	ctx.Rectangle(0, 0, width, height)
	ctx.Fill()
}

func (p *PlanetSimulator) setColor(ctx *cairo.Context, c color.Color) {
	r, g, b, a := c.RGBA()
	ctx.SetSourceRGBA(col(r), col(g), col(b), col(a))
}
