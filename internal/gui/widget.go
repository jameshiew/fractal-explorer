package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
	"gitlab.com/jameshiew/fractal-explorer/internal/draw"
	"gitlab.com/jameshiew/fractal-explorer/internal/gui/viewport"
	"gitlab.com/jameshiew/fractal-explorer/internal/mandelbrot"
	"image/color"
)

type fractalWidget struct {
	hideable
	size     fyne.Size
	position fyne.Position

	viewport viewport.Viewport
	info     *widget.Label
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
	raster := canvas.NewRaster(draw.New(func(pixelX, pixelY, width, height int) color.Color {
		x, y := f.viewport.PixelToCartesian(pixelX, pixelY, width, height)
		z := complex(x, y)
		return forMandelbrot(green, mandelbrot.New(70, 2))(z)
	}))
	renderer := &widgetRenderer{
		onRefresh: func() {
			f.refresh()
		},
		raster:  raster,
		objects: []fyne.CanvasObject{raster},
	}
	renderer.ApplyTheme()

	return renderer
}
