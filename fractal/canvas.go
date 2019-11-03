package fractal

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
)

type fractalCanvas struct {
	hidden   bool
	size     fyne.Size
	position fyne.Position
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

func (f *fractalCanvas) CreateRenderer() fyne.WidgetRenderer {
	renderer := &fractalRenderer{fractalCanvas: f}

	raster := canvas.NewRaster(renderer.draw)
	renderer.raster = raster
	renderer.objects = []fyne.CanvasObject{raster}
	renderer.ApplyTheme()

	return renderer
}
