package fractal

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"image"
	"image/color"
)

const title = "Fractal Explorer"

const (
	minWidth  = 320
	minHeight = 240
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
	return fyne.NewSize(minWidth, minHeight)
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

type fractalCanvas struct {
	hidden   bool
	size     fyne.Size
	position fyne.Position
}

func (f *fractalCanvas) CreateRenderer() fyne.WidgetRenderer {
	renderer := &fractalRenderer{fractalCanvas: f}

	raster := canvas.NewRaster(renderer.draw)
	renderer.raster = raster
	renderer.objects = []fyne.CanvasObject{raster}
	renderer.ApplyTheme()

	return renderer
}

func (f *fractalCanvas) Size() fyne.Size {
	return f.size
}

func (f *fractalCanvas) Resize(size fyne.Size) {
	f.size = size
	widget.Renderer(f).Layout(size)
}

func (f *fractalCanvas) Position() fyne.Position {
	return f.position
}

func (f *fractalCanvas) Move(position fyne.Position) {
	f.position = position
	widget.Renderer(f).Layout(f.size)
}

func (f *fractalCanvas) MinSize() fyne.Size {
	return widget.Renderer(f).MinSize()
}

func (f *fractalCanvas) Visible() bool {
	return !f.hidden
}

func (f *fractalCanvas) Show() {
	f.hidden = false
}

func (f *fractalCanvas) Hide() {
	f.hidden = true
}

// Run runs the application
func Run() {
	app := app.New()

	w := app.NewWindow(title)
	w.SetContent(&fractalCanvas{})

	w.ShowAndRun()
}
