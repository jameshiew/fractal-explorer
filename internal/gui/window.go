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
