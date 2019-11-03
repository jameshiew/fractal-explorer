package fractal

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func (f *fractalCanvas) Tapped(event *fyne.PointEvent) {
	f.viewport.center.x, f.viewport.center.y =
		f.viewport.toCartesian(
			event.Position.X,
			event.Position.Y,
			f.size.Width,
			f.size.Height,
		)
	widget.Refresh(f)
}

func (f *fractalCanvas) TappedSecondary(*fyne.PointEvent) {
	// do nothing
}
