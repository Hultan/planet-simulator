package planet_simulator

import (
	"time"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/planet-simulator/internal/data"
	"github.com/hultan/planet-simulator/internal/loader"
)

type PlanetSimulator struct {
	window      *gtk.ApplicationWindow
	drawingArea *gtk.DrawingArea

	tickerQuit chan struct{}
	ticker     *time.Ticker
	speed      time.Duration
	isActive   bool

	data *data.SolarSystem
}

func NewPlanetSimulator(w *gtk.ApplicationWindow, da *gtk.DrawingArea) *PlanetSimulator {
	t := &PlanetSimulator{window: w, drawingArea: da}
	t.window.Connect("key-press-event", t.onKeyPressed)

	l := loader.NewLoader()
	d, err := l.Load()
	if err != nil {
		panic(err)
	}
	t.data = d

	return t
}

func (p *PlanetSimulator) StartGame() {
	p.window.Maximize()
	p.drawingArea.Connect("draw", p.onDraw)
	p.speed = 500
	p.ticker = time.NewTicker(p.speed * time.Millisecond)
	p.tickerQuit = make(chan struct{})

	go p.mainLoop()
}

func (p *PlanetSimulator) mainLoop() {
	for {
		select {
		case <-p.ticker.C:
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

func (p *PlanetSimulator) quit() {
	if p.isActive {
		p.isActive = false
		close(p.tickerQuit) // Stop ticker
	}
}
