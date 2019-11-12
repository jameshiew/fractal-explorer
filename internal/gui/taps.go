package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"gitlab.com/jameshiew/fractal-explorer/internal/cartesian"
)

func (f *fractalWidget) Tapped(event *fyne.PointEvent) {
	deltaX, deltaY :=
		cartesian.Convert(
			event.Position.X,
			event.Position.Y,
			f.Size().Width,
			f.Size().Height,
		)
	f.viewport.Move(deltaX, deltaY)
	widget.Refresh(f)
}

func (f *fractalWidget) TappedSecondary(*fyne.PointEvent) {
	// do nothing
}
