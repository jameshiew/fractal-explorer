package gui

import (
	"fractal-explorer/internal/gui/viewport"
	"fractal-explorer/internal/mandelbrot"
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
	"image/color"
)

type fractalWidget struct {
	viewport viewport.Viewport

	labels struct {
		info *widget.Label
	}
	hidden   bool
	size     fyne.Size
	position fyne.Position
}

func newFractalWidget() fractalWidget {
	wdgt := fractalWidget{
		viewport: viewport.New(),
		labels:   struct{ info *widget.Label }{info: widget.NewLabel("")},
	}
	wdgt.Refresh()
	return wdgt
}

func (f *fractalWidget) Size() fyne.Size {
	return f.size
}

func (f *fractalWidget) Resize(size fyne.Size) {
	f.size = size
	widget.Renderer(f).Layout(size)
}

func (f *fractalWidget) Position() fyne.Position {
	return f.position
}

func (f *fractalWidget) Move(position fyne.Position) {
	f.position = position
	widget.Renderer(f).Layout(f.size)
}

func (f *fractalWidget) MinSize() fyne.Size {
	return widget.Renderer(f).MinSize()
}

func (f *fractalWidget) Visible() bool {
	return !f.hidden
}

func (f *fractalWidget) Show() {
	f.hidden = false
}

func (f *fractalWidget) Hide() {
	f.hidden = true
}

func (f *fractalWidget) Refresh() {
	f.labels.info.SetText(f.viewport.String())
}

func (f *fractalWidget) CreateRenderer() fyne.WidgetRenderer {
	renderer := &renderer{
		refresher: f,
		pixelColorer: func(pixelX, pixelY, width, height int) color.Color {
			colorer := forMandelbrot(mandelbrot.New(500, 2))
			return colorer(f.viewport.PixelToComplex(pixelX, pixelY, width, height))
		},
	}

	raster := canvas.NewRaster(renderer.draw)
	renderer.raster = raster
	renderer.objects = []fyne.CanvasObject{raster}
	renderer.ApplyTheme()

	return renderer
}
