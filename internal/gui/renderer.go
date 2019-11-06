package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"image"
	"image/color"
	"log"
	"sync"
	"time"
)

const (
	minWidthPixels  = 320
	minHeightPixels = 240
)

type instrumenter struct {
	rendered uint
}

func (i *instrumenter) instrument(nPixels int) (finish func()) {
	start := time.Now()
	return func() {
		duration := time.Since(start)
		log.Printf("%vpx in %v [%v px/s] (%v renders)", nPixels, duration, float64(nPixels)/float64(duration)*1_000_000_000, i.rendered)
		i.rendered++
	}
}

type widgetRenderer struct {
	instrumenter
	raster  *canvas.Raster
	objects []fyne.CanvasObject

	pixelColorer func(pixelX, pixelY, width, height int) color.Color
	onRefresh    func()
}

func (w widgetRenderer) Layout(size fyne.Size) {
	w.raster.Resize(size)
}

func (w widgetRenderer) MinSize() fyne.Size {
	return fyne.NewSize(minWidthPixels, minHeightPixels)
}

func (w widgetRenderer) Refresh() {
	w.onRefresh()
	canvas.Refresh(w.raster)
}

func (w widgetRenderer) ApplyTheme() {
	// do nothing
}

func (w widgetRenderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (w widgetRenderer) Objects() []fyne.CanvasObject {
	return w.objects
}

func (w widgetRenderer) Destroy() {
	// do nothing
}

// drawSingleThreaded is faster for larger canvases for whatever reason
func (w *widgetRenderer) drawSingleThreaded(width, height int) image.Image {
	nPixels := width * height
	defer w.instrument(nPixels)()
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, w.pixelColorer(x, y, width, height))
		}
	}
	return img
}

func (w *widgetRenderer) draw(width, height int) image.Image {
	nPixels := width * height
	defer w.instrument(nPixels)()
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
					color: w.pixelColorer(j.x, j.y, width, height),
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
