// Package draw is for drawing images of fractals
package draw

import (
	"image"
	"image/color"
	"sync"
	"time"
)

type logger interface {
	Infof(string, ...interface{})
}

type drawerInstrumenter struct {
	log      logger
	rendered uint
}

func (i *drawerInstrumenter) instrument(nPixels int) (finish func()) {
	start := time.Now()
	return func() {
		duration := time.Since(start)
		i.log.Infof("%vpx in %v [%v px/s] (%v renders)", nPixels, duration, float64(nPixels)/float64(duration)*1_000_000_000, i.rendered)
		i.rendered++
	}
}

type drawer struct {
	drawerInstrumenter
	pixelColorer func(pixelX, pixelY, width, height int) color.Color
}

// New returns a new drawing function that uses the passed pixel colorer func
func New(log logger, pixelColorer func(pixelX, pixelY, width, height int) color.Color) func(width, height int) image.Image {
	drwr := &drawer{
		drawerInstrumenter: drawerInstrumenter{log: log},
		pixelColorer:       pixelColorer,
	}
	return drwr.draw
}

// drawSingleThreaded is faster for larger canvases for whatever reason
func (d *drawer) drawSingleThreaded(width, height int) image.Image {
	nPixels := width * height
	defer d.instrument(nPixels)()
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, d.pixelColorer(x, y, width, height))
		}
	}
	return img
}

func (d *drawer) draw(width, height int) image.Image {
	nPixels := width * height
	defer d.instrument(nPixels)()
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	const nWorkers = 1024
	jobs := make(chan struct {
		x, y int
	}, nPixels)
	pixels := make(chan struct {
		x, y  int
		color color.Color
	}, nPixels)
	var wg sync.WaitGroup
	wg.Add(nPixels)
	go func() { // img.Set should only be called by one goroutine at a time, handle all calls via this goroutine
		for {
			select {
			case pxl := <-pixels:
				img.Set(pxl.x, pxl.y, pxl.color)
				wg.Done()
			}
		}
	}()
	for i := 0; i < nWorkers; i++ {
		go func() {
			for j := range jobs {
				pixels <- struct {
					x, y  int
					color color.Color
				}{
					x:     j.x,
					y:     j.y,
					color: d.pixelColorer(j.x, j.y, width, height),
				}
			}
		}()
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			jobs <- struct{ x, y int }{x: x, y: y}
		}
	}
	close(jobs)
	wg.Wait()
	return img
}
