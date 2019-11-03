package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

// Window gets a window containing the GUI
func Window(title string) fyne.Window {
	window := app.New().NewWindow(title)
	setUpWindow(window)
	return window
}

type controllable interface {
	Up()
	Left()
	Right()
	Down()
	ZoomIn()
	ZoomOut()
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
			controllable.ZoomIn()
		case fyne.KeyS:
			controllable.ZoomOut()
		}
	}
}

func setUpWindow(window fyne.Window) {
	cnvs := newFractalCanvas()
	window.SetContent(
		widget.NewVBox(
			&cnvs,
			cnvs.labels.info,
		),
	)
	ctrlr := controllerFor(&cnvs.viewport)
	window.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		ctrlr(event)
		widget.Refresh(&cnvs)
	})
}
