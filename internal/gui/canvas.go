package gui

import (
	"fractal-explorer/internal/gui/viewport"
	mandelbrot2 "fractal-explorer/internal/mandelbrot"
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
	"image/color"
)

func forMandelbrot(fractal mandelbrot2.Mandelbrot) func(complex128) color.Color {
	return func(c complex128) color.Color {
		iter := fractal.IterateWhileNotReachingBound(c)
		if iter == fractal.MaxIterations() {
			return color.Black
		}
		scale := float64(iter) / float64(fractal.MaxIterations())
		return color.RGBA{
			R: uint8(scale * 255),
			G: uint8(scale * 100),
			B: uint8(scale * 100),
			A: 255,
		}
	}
}

type fractalCanvas struct {
	viewport viewport.Viewport

	labels struct {
		info *widget.Label
	}
	hidden   bool
	size     fyne.Size
	position fyne.Position
}

func newFractalCanvas() fractalCanvas {
	return fractalCanvas{
		viewport: viewport.New(forMandelbrot(mandelbrot2.New(50, 2))),
		labels:   struct{ info *widget.Label }{info: widget.NewLabel("")},
	}
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

func (f *fractalCanvas) updateLabels() {
	f.labels.info.SetText(f.viewport.String())
}

func (f *fractalCanvas) CreateRenderer() fyne.WidgetRenderer {
	renderer := &renderer{canvas: f}

	raster := canvas.NewRaster(renderer.draw)
	renderer.raster = raster
	renderer.objects = []fyne.CanvasObject{raster}
	renderer.ApplyTheme()

	return renderer
}
