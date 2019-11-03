package fractal

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func (f *fractalCanvas) Tapped(event *fyne.PointEvent) {
	deltaX, deltaY :=
		toCartesian(
			event.Position.X,
			event.Position.Y,
			f.Size().Width,
			f.Size().Height,
		)
	f.viewport.center.x += deltaX
	f.viewport.center.y += deltaY
	widget.Refresh(f)
}

func (f *fractalCanvas) TappedSecondary(*fyne.PointEvent) {
	// do nothing
}
