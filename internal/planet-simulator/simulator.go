package planet_simulator

import (
	"time"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/planet-simulator/internal/data"
	"github.com/hultan/planet-simulator/internal/loader"
)

// PlanetSimulator : The planet simulator struct
type PlanetSimulator struct {
	window      *gtk.ApplicationWindow
	drawingArea *gtk.DrawingArea

	tickerQuit chan struct{}
	ticker     *time.Ticker
	speed      time.Duration
	isActive   bool

	data *data.SolarSystem // Solar system data
}

// NewPlanetSimulator : Constructor for a PlanetSimulator
func NewPlanetSimulator(w *gtk.ApplicationWindow, da *gtk.DrawingArea) *PlanetSimulator {
	p := &PlanetSimulator{window: w, drawingArea: da}
	p.window.Connect("key-press-event", p.onKeyPressed)

	// Load json file with a solar system
	l := loader.NewLoader()
	d, err := l.Load()
	if err != nil {
		panic(err)
	}
	p.data = d

	return p
}

// StartSimulator : Starts the simulation
func (p *PlanetSimulator) StartSimulator() {
	p.drawingArea.Connect("draw", p.onDraw)
	p.speed = 5
	p.ticker = time.NewTicker(p.speed * time.Millisecond)
	p.tickerQuit = make(chan struct{})

	go p.mainLoop()
}

// mainLoop : Start the main loop
func (p *PlanetSimulator) mainLoop() {
	for {
		select {
		case <-p.ticker.C:
			p.calculateMovements()
			p.drawingArea.QueueDraw()
		case <-p.tickerQuit:
			p.isActive = false
			p.ticker.Stop()
			return
		}
	}
}

// onKeyPressed : The onKeyPressed signal handler
func (p *PlanetSimulator) onKeyPressed(_ *gtk.ApplicationWindow, e *gdk.Event) {
	key := gdk.EventKeyNewFromEvent(e)

	switch key.KeyVal() {
	case 113: // Button "Q" => Quit game
		p.quit()
		p.window.Close() // Close window
	}
	p.drawingArea.QueueDraw()
}

// quit : Stops the simulation and quits the application
func (p *PlanetSimulator) quit() {
	if p.isActive {
		p.isActive = false
		close(p.tickerQuit) // Stop ticker
	}
}

// calculateMovements : Calculates the planets movements each cycle
func (p *PlanetSimulator) calculateMovements() {
	for i := range p.data.Bodies {
		body := p.data.Bodies[i]
		body.UpdatePosition(p.data, TIMESTEP)
	}
}
