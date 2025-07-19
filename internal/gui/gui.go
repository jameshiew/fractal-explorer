package gui

import "fyne.io/fyne/v2/app"

type logger interface {
	Info(msg string, args ...any)
}

// Run launches the GUI and blocks until exit
func Run(log logger, title string) {
	window := app.New().NewWindow(title)
	wdgt := newFractalWidget(log)
	layoutWindow(window, wdgt)
	window.RequestFocus()
	window.ShowAndRun()
}
