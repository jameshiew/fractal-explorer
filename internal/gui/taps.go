package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func (f *fractalCanvas) Tapped(event *fyne.PointEvent) {
	deltaX, deltaY :=
		ToCartesian(
			event.Position.X,
			event.Position.Y,
			f.Size().Width,
			f.Size().Height,
		)
	f.viewport.Move(deltaX, deltaY)
	widget.Refresh(f)
}

func (f *fractalCanvas) TappedSecondary(*fyne.PointEvent) {
	// do nothing
}
