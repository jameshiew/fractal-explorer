package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"

	"github.com/jameshiew/fractal-explorer/internal/cartesian"
)

func (f *fractalWidget) Tapped(event *fyne.PointEvent) {
	f.log.Infof("Tapped at (%v, %v)", event.Position.X, event.Position.Y)
	defer widget.Refresh(f)

	deltaX, deltaY :=
		cartesian.Convert(
			event.Position.X,
			event.Position.Y,
			f.Size().Width,
			f.Size().Height,
		)
	f.viewport.Move(deltaX, deltaY)
}

func (f *fractalWidget) TappedSecondary(*fyne.PointEvent) {
	// do nothing
}
