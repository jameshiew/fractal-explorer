package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"image/color"
)

const (
	minWidthPixels  = 320
	minHeightPixels = 240
)

type widgetRenderer struct {
	raster    *canvas.Raster
	objects   []fyne.CanvasObject
	onRefresh func()
}

func newWidgetRenderer(raster *canvas.Raster, onRefresh func()) widgetRenderer {
	return widgetRenderer{
		raster:    raster,
		objects:   []fyne.CanvasObject{raster},
		onRefresh: onRefresh,
	}
}

func (w widgetRenderer) Layout(size fyne.Size) {
	w.raster.Resize(size)
}

func (w widgetRenderer) MinSize() fyne.Size {
	return fyne.NewSize(minWidthPixels, minHeightPixels)
}

func (w widgetRenderer) Refresh() {
	w.onRefresh()
	canvas.Refresh(w.raster)
}

func (w widgetRenderer) ApplyTheme() {
	// do nothing
}

func (w widgetRenderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (w widgetRenderer) Objects() []fyne.CanvasObject {
	return w.objects
}

func (w widgetRenderer) Destroy() {
	// do nothing
}
