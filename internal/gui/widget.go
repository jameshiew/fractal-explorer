package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
	"gitlab.com/jameshiew/fractal-explorer/internal/gui/viewport"
	"gitlab.com/jameshiew/fractal-explorer/internal/mandelbrot"
	"image/color"
)

type hideable struct {
	hidden bool
}

func (h *hideable) Visible() bool {
	return !h.hidden
}

func (h *hideable) Show() {
	h.hidden = false
}

func (h *hideable) Hide() {
	h.hidden = true
}

type fractalWidget struct {
	hideable
	viewport viewport.Viewport

	info     *widget.Label
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
	wdgt.refresh()
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

func (f *fractalWidget) refresh() {
	f.info.SetText(f.viewport.String())
}

func (f *fractalWidget) CreateRenderer() fyne.WidgetRenderer {
	renderer := &widgetRenderer{
		onRefresh: func() {
			f.refresh()
		},
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
