package gui

import (
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"

	"gitlab.com/jameshiew/fractal-explorer/internal/cartesian"
)

func (f *fractalWidget) Tapped(event *fyne.PointEvent) {
	log.Printf("Tapped at (%v, %v)", event.Position.X, event.Position.Y)
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
