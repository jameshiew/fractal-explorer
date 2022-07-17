package gui

import (
	"fyne.io/fyne/v2"

	"github.com/jameshiew/fractal-explorer/internal/cartesian"
)

func (f *fractalWidget) Tapped(event *fyne.PointEvent) {
	f.log.Infof("Tapped at (%v, %v)", event.Position.X, event.Position.Y)
	defer func() {
		if f.renderer == nil {
			return
		}
		f.renderer.Refresh()
	}()

	deltaX, deltaY :=
		cartesian.Convert(
			int(event.Position.X),
			int(event.Position.Y),
			int(f.Size().Width),
			int(f.Size().Height),
		)
	f.viewport.Move(deltaX, deltaY)
}

func (f *fractalWidget) TappedSecondary(*fyne.PointEvent) {
	// do nothing
}
