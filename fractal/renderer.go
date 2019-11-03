package fractal

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"image"
	"image/color"
)

type renderer struct {
	raster   *canvas.Raster
	objects  []fyne.CanvasObject
	imgCache *image.RGBA

	fractalCanvas *fractalCanvas
}

func (f renderer) Layout(size fyne.Size) {
	f.raster.Resize(size)
}

func (f renderer) MinSize() fyne.Size {
	return fyne.NewSize(minWidthPixels, minHeightPixels)
}

func (f renderer) Refresh() {
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
	img := f.imgCache
	if img == nil || img.Bounds().Size().X != w || img.Bounds().Size().Y != h {
		img = image.NewRGBA(image.Rect(0, 0, w, h))
		f.imgCache = img
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.Set(x, y, f.fractalCanvas.drawer.pixelColor(x, y, w, h))
		}
	}
	return img
}
