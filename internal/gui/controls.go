package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func (f *fractalWidget) controllerFunc() func(event *fyne.KeyEvent) {
	return func(event *fyne.KeyEvent) {
		switch event.Name {
		case fyne.KeyUp:
			f.viewport.Up()
		case fyne.KeyLeft:
			f.viewport.Left()
		case fyne.KeyRight:
			f.viewport.Right()
		case fyne.KeyDown:
			f.viewport.Down()
		case fyne.KeyW:
			f.viewport.Zoom(0.5)
		case fyne.KeyS:
			f.viewport.Zoom(2)
		default:
			// unhandled key event so don't do anything
			return
		}
		widget.Refresh(f)
	}
}
