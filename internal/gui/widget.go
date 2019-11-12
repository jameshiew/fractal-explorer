package gui

import (
	"fmt"
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"

	"gitlab.com/jameshiew/fractal-explorer/internal/cartesian"
)

type fractalWidget struct {
	hideable
	size     fyne.Size
	position fyne.Position

	viewport cartesian.Viewport
	info     *widget.Label
}

func newFractalWidget() fractalWidget {
	wdgt := fractalWidget{
		viewport: cartesian.NewViewport(),
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
	log.Printf("Resized to %v\n", size)
	widget.Renderer(f).Layout(size)
	widget.Renderer(f).Refresh()
}

func (f *fractalWidget) Position() fyne.Position {
	return f.position
}

func (f *fractalWidget) Move(position fyne.Position) {
	f.position = position
	widget.Renderer(f).Layout(f.size)
	widget.Renderer(f).Refresh()
}

func (f *fractalWidget) MinSize() fyne.Size {
	return widget.Renderer(f).MinSize()
}

func (f *fractalWidget) String() string {
	return fmt.Sprintf("%v - ", f.Size()) + f.viewport.String()
}

// InfoLabel returns a label which is updated with the information for this fractal widget
func (f *fractalWidget) InfoLabel() fyne.CanvasObject {
	return f.info
}

func (f *fractalWidget) refresh() {
	f.info.SetText(f.String())
}
