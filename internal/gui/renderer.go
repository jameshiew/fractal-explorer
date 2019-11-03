package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"image"
	"image/color"
	"log"
	"time"
)

const (
	minWidthPixels  = 320
	minHeightPixels = 240
)

// refresher implementations provide a callback for when they should be refreshed
type refresher interface {
	Refresh()
}

type instrumenter struct {
	rendered uint
}

func (i *instrumenter) instrument() (finish func()) {
	start := time.Now()
	return func() {
		duration := time.Since(start)
		log.Printf("Took %v to render image", duration)
		i.rendered++
	}
}

type renderer struct {
	instrumenter
	raster  *canvas.Raster
	objects []fyne.CanvasObject

	pixelColorer func(pixelX, pixelY, width, height int) color.Color
	refresher    refresher
}

func (f renderer) Layout(size fyne.Size) {
	f.raster.Resize(size)
}

func (f renderer) MinSize() fyne.Size {
	return fyne.NewSize(minWidthPixels, minHeightPixels)
}

func (f renderer) Refresh() {
	f.refresher.Refresh()
	canvas.Refresh(f.raster)
}

func (f renderer) ApplyTheme() {
	// do nothing
}

func (f renderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (f renderer) Objects() []fyne.CanvasObject {
	return f.objects
}

func (f renderer) Destroy() {
	// do nothing
}

func (f *renderer) draw(w, h int) image.Image {
	defer f.instrument()()
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.Set(x, y, f.pixelColorer(x, y, w, h))
		}
	}
	return img
}
