package gui

import "fyne.io/fyne"

type controllable interface {
	Up()
	Left()
	Right()
	Down()
	Zoom(factor float64)
}

func controllerFor(controllable controllable) func(*fyne.KeyEvent) {
	return func(event *fyne.KeyEvent) {
		switch event.Name {
		case fyne.KeyUp:
			controllable.Up()
		case fyne.KeyLeft:
			controllable.Left()
		case fyne.KeyRight:
			controllable.Right()
		case fyne.KeyDown:
			controllable.Down()
		case fyne.KeyW:
			controllable.Zoom(0.5)
		case fyne.KeyS:
			controllable.Zoom(2)
		}
	}
}
