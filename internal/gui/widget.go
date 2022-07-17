package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	"github.com/jameshiew/fractal-explorer/internal/cartesian"
)

type fractalWidget struct {
	hideable
	log      logger
	renderer fyne.WidgetRenderer
	size     fyne.Size
	position fyne.Position

	viewport cartesian.Viewport
	info     *widget.Label
}

func newFractalWidget(log logger) fractalWidget {
	wdgt := fractalWidget{
		log:      log,
		viewport: cartesian.NewViewport(),
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
	f.log.Infof("Resized to %v", size)
	if f.renderer == nil {
		return
	}
	f.renderer.Layout(size)
	f.renderer.Refresh()
}

func (f *fractalWidget) Position() fyne.Position {
	return f.position
}

func (f *fractalWidget) Move(position fyne.Position) {
	f.position = position
	if f.renderer == nil {
		return
	}
	f.renderer.Layout(f.size)
	f.renderer.Refresh()
}

func (f *fractalWidget) MinSize() fyne.Size {
	if f.renderer == nil {
		return fyne.Size{}
	}
	return f.renderer.MinSize()
}

func (f *fractalWidget) String() string {
	return fmt.Sprintf("%v - ", f.Size()) + f.viewport.String()
}

// InfoLabel returns a label which is updated with the information for this fractal widget
func (f *fractalWidget) InfoLabel() fyne.CanvasObject {
	return f.info
}

func (f *fractalWidget) Refresh() {
	f.info.SetText(f.String())
}
