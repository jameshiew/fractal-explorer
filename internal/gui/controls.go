package gui

import (
	"fyne.io/fyne/v2"
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
		if f.renderer == nil {
			return
		}
		f.renderer.Refresh()
	}
}
