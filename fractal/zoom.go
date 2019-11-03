package fractal

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func (f *fractalCanvas) Tapped(event *fyne.PointEvent) {
	f.drawer.position.x, f.drawer.position.y =
		f.drawer.toCartesian(
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
