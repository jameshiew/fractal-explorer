package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
	"gitlab.com/jameshiew/fractal-explorer/internal/gui/viewport"
	"gitlab.com/jameshiew/fractal-explorer/internal/mandelbrot"
	"image/color"
	"math"
)

type fractalWidget struct {
	viewport viewport.Viewport

	info     *widget.Label
	hidden   bool
	size     fyne.Size
	position fyne.Position
}

// InfoLabel returns a label which is updated with the information for this fractal widget
func (f *fractalWidget) InfoLabel() fyne.CanvasObject {
	return f.info
}

func newFractalWidget() fractalWidget {
	wdgt := fractalWidget{
		viewport: viewport.New(),
		info:     widget.NewLabel(""),
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
	f.info.SetText(f.viewport.String())
}

// darkBlend is quite dark
func darkBlend(z complex128) color.Color {
	return blend(
		forMandelbrot(green, mandelbrot.New(125, 2))(z),
		forMandelbrot(blue, mandelbrot.New(250, 2))(z),
		forMandelbrot(red, mandelbrot.New(500, 2))(z),
	)
}

func otherBlend(z complex128) color.Color {
	return blend(
		forMandelbrot(green, mandelbrot.New(120, math.Phi))(z),
		forMandelbrot(color.RGBA64{
			R: 20000,
			G: 50000,
			B: 20000,
			A: 65535,
		}, mandelbrot.New(100, math.E))(z),
		forMandelbrot(color.RGBA64{
			R: 16000,
			G: 65335,
			B: 16000,
			A: 65535,
		}, mandelbrot.New(75, math.Pi))(z),
	)
}

func (f *fractalWidget) CreateRenderer() fyne.WidgetRenderer {
	renderer := &renderer{
		refresher: f,
		pixelColorer: func(pixelX, pixelY, width, height int) color.Color {
			x, y := f.viewport.PixelToCartesian(pixelX, pixelY, width, height)
			z := complex(x, y)
			return forMandelbrot(green, mandelbrot.New(70, 2))(z)
		},
	}

	raster := canvas.NewRaster(renderer.draw)
	renderer.raster = raster
	renderer.objects = []fyne.CanvasObject{raster}
	renderer.ApplyTheme()

	return renderer
}
