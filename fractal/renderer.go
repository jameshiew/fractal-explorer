package fractal

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"image"
	"image/color"
)

type fractalRenderer struct {
	raster   *canvas.Raster
	objects  []fyne.CanvasObject
	imgCache *image.RGBA

	fractalCanvas *fractalCanvas
}

func (f fractalRenderer) Layout(size fyne.Size) {
	f.raster.Resize(size)
}

func (f fractalRenderer) MinSize() fyne.Size {
	return fyne.NewSize(minWidthPixels, minHeightPixels)
}

func (f fractalRenderer) Refresh() {
	canvas.Refresh(f.raster)
}

func (f fractalRenderer) ApplyTheme() {
	// do nothing
}

func (f fractalRenderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (f fractalRenderer) Objects() []fyne.CanvasObject {
	return f.objects
}

func (f fractalRenderer) Destroy() {
	// do nothing
}

func (f *fractalRenderer) draw(w, h int) image.Image {
	img := f.imgCache
	if img == nil || img.Bounds().Size().X != w || img.Bounds().Size().Y != h {
		img = image.NewRGBA(image.Rect(0, 0, w, h))
		f.imgCache = img
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			drwr := &drawer{
				scale: struct {
					x, y float64
				}{
					x: 0.01,
					y: 0.01,
				},
				mandelbrot: mandelbrot{
					maxIterations: 50,
					bound:         2,
				},
			}
			img.Set(x, y, drwr.pixelColor(x, y, w, h))
		}
	}
	return img
}
