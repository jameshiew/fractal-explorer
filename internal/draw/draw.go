// Package draw is for drawing images of fractals
package draw

import (
	"image"
	"image/color"
	"runtime"
	"sync"
	"time"
)

type logger interface {
	Info(msg string, args ...any)
}

type drawerInstrumenter struct {
	log      logger
	rendered uint
}

func (i *drawerInstrumenter) instrument(nPixels int) (finish func()) {
	start := time.Now()
	return func() {
		duration := time.Since(start)
		pixelsPerSecond := float64(nPixels) / duration.Seconds()
		i.log.Info("Render completed", "pixels", nPixels, "duration", duration, "pixels_per_second", pixelsPerSecond, "render_count", i.rendered)
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

func (d *drawer) draw(width, height int) image.Image {
	nPixels := width * height
	defer d.instrument(nPixels)()
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	nWorkers := runtime.NumCPU()
	if nWorkers < 1 {
		nWorkers = 1
	}

	rows := make(chan int, height)
	var wg sync.WaitGroup

	for i := 0; i < nWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for y := range rows {
				for x := 0; x < width; x++ {
					img.Set(x, y, d.pixelColorer(x, y, width, height))
				}
			}
		}()
	}

	for y := 0; y < height; y++ {
		rows <- y
	}
	close(rows)
	wg.Wait()

	return img
}
